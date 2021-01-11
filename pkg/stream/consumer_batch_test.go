package stream_test

import (
	"context"
	"fmt"
	"github.com/applike/gosoline/pkg/mdl"
	monMocks "github.com/applike/gosoline/pkg/mon/mocks"
	"github.com/applike/gosoline/pkg/stream"
	"github.com/applike/gosoline/pkg/stream/mocks"
	"github.com/applike/gosoline/pkg/tracing"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type acknowledgableInput struct {
	mocks.AcknowledgeableInput
	mocks.Input
}

type BatchConsumerTestSuite struct {
	suite.Suite

	data chan *stream.Message
	once sync.Once
	stop func()

	input *acknowledgableInput

	callback      *mocks.RunnableBatchConsumerCallback
	batchConsumer *stream.BatchConsumer
}

func TestBatchConsumerTestSuite(t *testing.T) {
	suite.Run(t, new(BatchConsumerTestSuite))
}

func (s *BatchConsumerTestSuite) SetupTest() {
	s.data = make(chan *stream.Message, 10)
	s.once = sync.Once{}
	s.stop = func() {
		s.once.Do(func() {
			close(s.data)
		})
	}

	s.input = new(acknowledgableInput)
	s.callback = new(mocks.RunnableBatchConsumerCallback)

	logger := monMocks.NewLoggerMockedAll()
	tracer := tracing.NewNoopTracer()
	mw := monMocks.NewMetricWriterMockedAll()
	me := stream.NewMessageEncoder(&stream.MessageEncoderSettings{})
	settings := &stream.ConsumerSettings{
		Input:       "test",
		RunnerCount: 1,
		IdleTimeout: time.Second,
	}
	batchSettings := &stream.BatchConsumerSettings{
		IdleTimeout: time.Second,
		BatchSize:   5,
	}

	s.batchConsumer = stream.NewBatchConsumer("test", s.callback)
	s.batchConsumer.BootWithInterfaces(logger, tracer, mw, s.input, me, settings, batchSettings)
}

func (s *BatchConsumerTestSuite) TestRun_ProcessOnStop() {
	s.input.Input.
		On("Data").
		Return(s.data)

	s.input.Input.
		On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Run(func(args mock.Arguments) {
			s.data <- stream.NewJsonMessage(`"foo"`)
			s.data <- stream.NewJsonMessage(`"bar"`)
			s.data <- stream.NewJsonMessage(`"foobar"`)
			s.stop()
		}).Return(nil)

	processed := 0

	s.input.Input.
		On("Stop").
		Return()

	s.input.AcknowledgeableInput.
		On("AckBatch", mock.AnythingOfType("[]*stream.Message")).
		Run(func(args mock.Arguments) {
			msgs := args[0].([]*stream.Message)
			processed = len(msgs)
		}).
		Return(nil)

	acks := []bool{true, true, true}
	s.callback.On("Consume", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]interface {}"), mock.AnythingOfType("[]map[string]interface {}")).
		Return(acks, nil)

	s.callback.On("GetModel", mock.AnythingOfType("map[string]interface {}")).
		Return(func(_ map[string]interface{}) interface{} {
			return mdl.String("")
		})

	s.callback.On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Return(nil)

	err := s.batchConsumer.Run(context.Background())

	s.NoError(err, "there should be no error during run")
	s.Equal(3, processed)

	s.input.Input.AssertExpectations(s.T())
	s.input.AcknowledgeableInput.AssertExpectations(s.T())
	s.callback.AssertExpectations(s.T())
}

func (s *BatchConsumerTestSuite) TestRun_BatchSizeReached() {
	s.input.Input.
		On("Data").
		Return(s.data)

	s.input.Input.
		On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Run(func(args mock.Arguments) {
			s.data <- stream.NewJsonMessage(`"foo"`)
			s.data <- stream.NewJsonMessage(`"bar"`)
			s.data <- stream.NewJsonMessage(`"foobar"`)
			s.data <- stream.NewJsonMessage(`"barfoo"`)
			s.data <- stream.NewJsonMessage(`"foobarfoo"`)
		}).Return(nil)

	s.input.Input.
		On("Stop").
		Return()

	processed := 0

	s.input.AcknowledgeableInput.
		On("AckBatch", mock.AnythingOfType("[]*stream.Message")).
		Run(func(args mock.Arguments) {
			msgs := args[0].([]*stream.Message)
			processed = len(msgs)

			s.stop()
		}).
		Return(nil)

	acks := []bool{true, true, true, true, true}
	s.callback.On("Consume", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]interface {}"), mock.AnythingOfType("[]map[string]interface {}")).
		Return(acks, nil)

	s.callback.On("GetModel", mock.AnythingOfType("map[string]interface {}")).
		Return(func(_ map[string]interface{}) interface{} {
			return mdl.String("")
		})

	s.callback.On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Return(nil)

	err := s.batchConsumer.Run(context.Background())

	s.NoError(err, "there should be no error during run")
	s.Equal(5, processed)

	s.input.Input.AssertExpectations(s.T())
	s.input.AcknowledgeableInput.AssertExpectations(s.T())
	s.callback.AssertExpectations(s.T())
}

func (s *BatchConsumerTestSuite) TestRun_ContextCanceled() {
	ctx, cancel := context.WithCancel(context.Background())
	stopped := make(chan struct{})
	once := sync.Once{}

	s.input.Input.
		On("Data").
		Return(s.data)

	s.input.Input.
		On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Run(func(args mock.Arguments) {
			cancel()
			<-stopped
			s.stop()
		}).Return(nil)

	processed := 0

	s.input.Input.
		On("Stop").
		Run(func(args mock.Arguments) {
			once.Do(func() {
				close(stopped)
			})
		}).
		Return(nil)

	s.callback.On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Return(nil)

	err := s.batchConsumer.Run(ctx)

	s.NoError(err, "there should be no error during run")
	s.Equal(0, processed)

	s.input.Input.AssertExpectations(s.T())
	s.input.AcknowledgeableInput.AssertExpectations(s.T())
	s.callback.AssertExpectations(s.T())
}

func (s *BatchConsumerTestSuite) TestRun_InputRunError() {
	s.input.Input.
		On("Data").
		Return(s.data)

	s.input.Input.
		On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Return(fmt.Errorf("read error"))

	s.callback.
		On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Run(func(args mock.Arguments) {
			<-args[0].(context.Context).Done()
		}).Return(nil)

	err := s.batchConsumer.Run(context.Background())

	s.EqualError(err, "error while waiting for all routines to stop: panic during run of the consumer input: read error")

	s.input.Input.AssertExpectations(s.T())
	s.input.AcknowledgeableInput.AssertExpectations(s.T())
	s.callback.AssertExpectations(s.T())
}

func (s *BatchConsumerTestSuite) TestRun_CallbackRunError() {
	s.input.Input.On("Data").
		Return(s.data)

	s.input.Input.On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Run(func(args mock.Arguments) {
			<-args[0].(context.Context).Done()
		}).
		Return(nil)

	s.callback.On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Return(fmt.Errorf("consumerCallback run error"))

	err := s.batchConsumer.Run(context.Background())

	s.EqualError(err, "error while waiting for all routines to stop: panic during run of the consumerCallback: consumerCallback run error")

	s.input.Input.AssertExpectations(s.T())
	s.input.AcknowledgeableInput.AssertExpectations(s.T())
	s.callback.AssertExpectations(s.T())
}

func (s *BatchConsumerTestSuite) TestRun_AggregateMessage() {
	message1 := stream.NewJsonMessage(`"foo"`, map[string]interface{}{
		"attr1": "a",
	})
	message2 := stream.NewJsonMessage(`"bar"`, map[string]interface{}{
		"attr1": "b",
	})

	aggregate, err := stream.BuildAggregateMessage(stream.MarshalJsonMessage, []stream.WritableMessage{message1, message2})

	s.Require().NoError(err)

	s.input.Input.
		On("Data").
		Return(s.data)

	s.input.Input.
		On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Run(func(args mock.Arguments) {
			s.data <- aggregate.(*stream.Message)
			s.stop()
		}).Return(nil)

	s.input.Input.
		On("Stop").
		Return()

	processed := 0
	s.input.AcknowledgeableInput.
		On("AckBatch", mock.AnythingOfType("[]*stream.Message")).
		Run(func(args mock.Arguments) {
			msgs := args[0].([]*stream.Message)
			processed = len(msgs)
		}).
		Return(nil)

	s.callback.On("Run", mock.AnythingOfType("*context.cancelCtx")).
		Return(nil)

	acks := []bool{true, true}

	s.callback.On("Consume", mock.AnythingOfType("*context.cancelCtx"), mock.AnythingOfType("[]interface {}"), mock.AnythingOfType("[]map[string]interface {}")).
		Return(acks, nil)

	s.callback.
		On("GetModel", mock.AnythingOfType("map[string]interface {}")).
		Return(mdl.String("")).
		Twice()

	err = s.batchConsumer.Run(context.Background())

	s.Nil(err, "there should be no error returned on consume")
	s.Equal(processed, 2)

	s.input.Input.AssertExpectations(s.T())
	s.input.AcknowledgeableInput.AssertExpectations(s.T())
	s.callback.AssertExpectations(s.T())
}
