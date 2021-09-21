package stream

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/cloud/aws/sqs"
	"github.com/justtrackio/gosoline/pkg/coffin"
	"github.com/justtrackio/gosoline/pkg/log"
)

var _ AcknowledgeableInput = &sqsInput{}

type SqsInputSettings struct {
	cfg.AppId
	QueueId             string            `cfg:"queue_id"`
	MaxNumberOfMessages int32             `cfg:"max_number_of_messages" default:"10" validate:"min=1,max=10"`
	WaitTime            int32             `cfg:"wait_time"`
	VisibilityTimeout   int               `cfg:"visibility_timeout"`
	RunnerCount         int               `cfg:"runner_count"`
	Fifo                sqs.FifoSettings  `cfg:"fifo"`
	RedrivePolicy       sqs.RedrivePolicy `cfg:"redrive_policy"`
	ClientName          string            `cfg:"client_name"`
	Unmarshaller        string            `cfg:"unmarshaller" default:"msg"`
}

func (s SqsInputSettings) GetAppid() cfg.AppId {
	return s.AppId
}

func (s SqsInputSettings) GetQueueId() string {
	return s.QueueId
}

func (s SqsInputSettings) IsFifoEnabled() bool {
	return s.Fifo.Enabled
}

type sqsInput struct {
	logger      log.Logger
	queue       sqs.Queue
	settings    *SqsInputSettings
	unmarshaler UnmarshallerFunc

	cfn     coffin.Coffin
	channel chan *Message
	stopped bool
}

func NewSqsInput(ctx context.Context, config cfg.Config, logger log.Logger, settings *SqsInputSettings) (*sqsInput, error) {
	settings.AppId.PadFromConfig(config)

	queueName := sqs.GetQueueName(settings)
	queueSettings := &sqs.Settings{
		QueueName:         queueName,
		VisibilityTimeout: settings.VisibilityTimeout,
		Fifo:              settings.Fifo,
		RedrivePolicy:     settings.RedrivePolicy,
		ClientName:        settings.ClientName,
	}

	var ok bool
	var err error
	var queue sqs.Queue
	var unmarshaller UnmarshallerFunc

	if queue, err = sqs.NewQueue(ctx, config, logger, queueSettings); err != nil {
		return nil, fmt.Errorf("can not create queue: %w", err)
	}

	if unmarshaller, ok = unmarshallers[settings.Unmarshaller]; !ok {
		return nil, fmt.Errorf("unknown unmarshaller %s", settings.Unmarshaller)
	}

	return NewSqsInputWithInterfaces(logger, queue, unmarshaller, settings), nil
}

func NewSqsInputWithInterfaces(logger log.Logger, queue sqs.Queue, unmarshaller UnmarshallerFunc, settings *SqsInputSettings) *sqsInput {
	if settings.RunnerCount <= 0 {
		settings.RunnerCount = 1
	}

	return &sqsInput{
		logger:      logger,
		queue:       queue,
		settings:    settings,
		unmarshaler: unmarshaller,
		cfn:         coffin.New(),
		channel:     make(chan *Message),
	}
}

func (i *sqsInput) Data() chan *Message {
	return i.channel
}

func (i *sqsInput) Run(ctx context.Context) error {
	defer close(i.channel)
	defer i.logger.Info("leaving sqs input")

	i.logger.Info("starting sqs input with %d runners", i.settings.RunnerCount)

	for j := 0; j < i.settings.RunnerCount; j++ {
		i.cfn.Gof(func() error {
			return i.runLoop(ctx)
		}, "panic in sqs input runner")
	}

	<-i.cfn.Dying()
	i.Stop()

	return i.cfn.Wait()
}

func (i *sqsInput) runLoop(ctx context.Context) error {
	defer i.logger.Info("leaving sqs input runner")

	for {
		if i.stopped {
			return nil
		}

		sqsMessages, err := i.queue.Receive(ctx, i.settings.MaxNumberOfMessages, i.settings.WaitTime)
		if err != nil {
			i.logger.Error("could not get messages from sqs: %w", err)
			continue
		}

		for _, sqsMessage := range sqsMessages {
			msg, err := i.unmarshaler(sqsMessage.Body)
			if err != nil {
				i.logger.Error("could not unmarshal message: %w", err)
				continue
			}

			if msg.Attributes == nil {
				msg.Attributes = make(map[string]interface{})
			}

			msg.Attributes[AttributeSqsMessageId] = *sqsMessage.MessageId
			msg.Attributes[AttributeSqsReceiptHandle] = *sqsMessage.ReceiptHandle

			i.channel <- msg
		}
	}
}

func (i *sqsInput) Stop() {
	i.stopped = true
}

func (i *sqsInput) Ack(ctx context.Context, msg *Message) error {
	var ok bool
	var receiptHandleInterface interface{}
	var receiptHandleString string

	if receiptHandleInterface, ok = msg.Attributes[AttributeSqsReceiptHandle]; !ok {
		return fmt.Errorf("the message has no attribute %s", AttributeSqsReceiptHandle)
	}

	if receiptHandleString, ok = receiptHandleInterface.(string); !ok {
		return fmt.Errorf("the attribute %s of the message should be string but instead is %T", AttributeSqsReceiptHandle, receiptHandleInterface)
	}

	if receiptHandleString == "" {
		return fmt.Errorf("the attribute %s of the message should not be empty", AttributeSqsReceiptHandle)
	}

	return i.queue.DeleteMessage(ctx, receiptHandleString)
}

func (i *sqsInput) AckBatch(ctx context.Context, msgs []*Message) error {
	receiptHandles := make([]string, 0, len(msgs))
	multiError := new(multierror.Error)

	for _, msg := range msgs {
		receiptHandleInterface, ok := msg.Attributes[AttributeSqsReceiptHandle]

		if !ok {
			multiError = multierror.Append(multiError, fmt.Errorf("the message has no attribute %s", AttributeSqsReceiptHandle))

			continue
		}

		receiptHandleString, ok := receiptHandleInterface.(string)

		if !ok {
			multiError = multierror.Append(multiError, fmt.Errorf("the attribute %s of the message should be string but instead is %T", AttributeSqsReceiptHandle, receiptHandleInterface))

			continue
		}

		if receiptHandleString == "" {
			multiError = multierror.Append(multiError, fmt.Errorf("the attribute %s of the message must not be empty", AttributeSqsReceiptHandle))

			continue
		}

		receiptHandles = append(receiptHandles, receiptHandleString)
	}

	if len(receiptHandles) == 0 {
		return multiError.ErrorOrNil()
	}

	if err := i.queue.DeleteMessageBatch(ctx, receiptHandles); err != nil {
		multiError = multierror.Append(multiError, err)
	}

	return multiError.ErrorOrNil()
}

func (i *sqsInput) SetUnmarshaler(unmarshaler UnmarshallerFunc) {
	i.unmarshaler = unmarshaler
}

func (i *sqsInput) GetQueueUrl() string {
	return i.queue.GetUrl()
}

func (i *sqsInput) GetQueueArn() string {
	return i.queue.GetArn()
}
