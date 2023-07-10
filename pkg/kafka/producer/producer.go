package producer

import (
	"context"
	"fmt"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/coffin"
	"github.com/justtrackio/gosoline/pkg/kafka/connection"
	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	Settings       *Settings
	Writer         Writer
	Logger         log.Logger
	pool           coffin.Coffin
	activeBalancer KafkaBalancer
}

// NewProducer returns a topic producer.
func NewProducer(_ context.Context, conf cfg.Config, logger log.Logger, name string) (*Producer, error) {
	settings := ParseSettings(conf, name)

	// Connection.
	dialer, err := connection.NewDialer(settings.Connection())
	if err != nil {
		return nil, fmt.Errorf("kafka: failed to get dialer: %w", err)
	}

	// Writer.
	writer, err := NewWriter(logger, dialer, settings.Connection().Bootstrap, getOptions(settings)...)
	if err != nil {
		return nil, fmt.Errorf("kafka: failed to get writer: %w", err)
	}

	return NewProducerWithInterfaces(settings, logger, writer)
}

func NewProducerWithInterfaces(conf *Settings, logger log.Logger, writer Writer) (*Producer, error) {
	logger = logger.WithFields(
		log.Fields{
			"kafka_topic":         conf.FQTopic,
			"kafka_batch_size":    conf.BatchSize,
			"kafka_batch_timeout": conf.BatchTimeout,
		},
	)
	if initableBalancer, ok := kafkaBalancers[conf.Balancer].(Initable); ok {
		initableBalancer.Init(logger)
	}

	return &Producer{
		Settings:       conf,
		Writer:         writer,
		Logger:         logging.NewKafkaLogger(logger),
		pool:           coffin.New(),
		activeBalancer: kafkaBalancers[conf.Balancer],
	}, nil
}

// Run starts background routine for flushing messages.
func (p *Producer) Run(ctx context.Context) error {
	p.Logger.Info("starting producer")
	defer p.Logger.Info("shutdown producer")

	p.pool.GoWithContext(ctx, p.flushOnExit)

	return p.pool.Wait()
}

func (p *Producer) WriteOne(ctx context.Context, m kafka.Message) error {
	return p.write(ctx, m)
}

func (p *Producer) Write(ctx context.Context, ms ...kafka.Message) error {
	return p.write(ctx, ms...)
}

func (p *Producer) write(ctx context.Context, ms ...kafka.Message) error {
	ctx, cancel := context.WithTimeout(ctx, DefaultWriterWriteTimeout)
	defer cancel()

	p.Logger.Debug("producing messages")

	// Prepare batch.
	batch := []kafka.Message{}

	for _, m := range ms {
		batch = append(batch, kafka.Message{
			Topic:   p.Settings.FQTopic,
			Key:     m.Key,
			Value:   m.Value,
			Headers: m.Headers,
		})
	}

	return p.writeBatch(ctx, batch, 0)
}

func (p *Producer) writeBatch(ctx context.Context, batch []kafka.Message, attempt int) error {
	attempt += 1
	failedWrites := []kafka.Message{}

	// if the error return is of type write Errors
	switch err := p.Writer.WriteMessages(ctx, batch...).(type) {
	case nil:
		for i := range batch {
			p.activeBalancer.OnSuccess(batch[i])
		}
		return nil
	case kafka.WriteErrors:
		for i := range err {
			if err[i] == nil {
				p.activeBalancer.OnSuccess(batch[i])
				continue
			}
			p.Logger.WithFields(log.Fields{
				"err": err[i],
			}).Error("error while writing a message to kafka")
			p.activeBalancer.OnError(batch[i], err[i])
			failedWrites = append(failedWrites, batch[i])
		}

		if attempt > p.Settings.Retries {
			return err
		}

		return p.writeBatch(ctx, failedWrites, attempt)
	default:
		return err
	}
}

func (p *Producer) flushOnExit(ctx context.Context) error {
	<-ctx.Done()

	p.Logger.Info("flushing messages")
	defer p.Logger.Info("flushed messages")

	if err := p.Writer.Close(); err != nil {
		p.Logger.WithFields(log.Fields{"Error": err}).Error("failed to flush messages")
	}

	return nil
}
