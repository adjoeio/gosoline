package stream_test

import (
	"context"
	"errors"
	"testing"

	logMocks "github.com/justtrackio/gosoline/pkg/log/mocks"
	"github.com/justtrackio/gosoline/pkg/stream"
	"github.com/justtrackio/gosoline/pkg/stream/mocks"
	"github.com/stretchr/testify/suite"
)

type testContent struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ProducerTestSuite struct {
	suite.Suite

	ctx         context.Context
	encoder     stream.MessageEncoder
	output      *mocks.Output
	producer    stream.Producer
	retryDaemon *mocks.ProducerRetryDaemon
}

func (s *ProducerTestSuite) SetupTest() {
	logger := logMocks.NewLoggerMock(logMocks.WithMockAll)
	s.ctx = context.Background()
	s.encoder = stream.NewMessageEncoder(&stream.MessageEncoderSettings{
		Encoding: stream.EncodingJson,
	})
	s.output = new(mocks.Output)
	s.retryDaemon = new(mocks.ProducerRetryDaemon)
	s.producer = stream.NewProducerWithInterfaces(logger, s.encoder, s.output, s.retryDaemon, false)
}

func (s *ProducerTestSuite) TestProducer_WriteOne() {
	content := &testContent{
		Id:   3,
		Name: "foobar",
	}

	expectedMsg := &stream.Message{
		Attributes: map[string]string{
			stream.AttributeEncoding: stream.EncodingJson.String(),
		},
		Body: `{"id":3,"name":"foobar"}`,
	}

	s.output.On("WriteOne", s.ctx, expectedMsg).Return(nil)
	err := s.producer.WriteOne(s.ctx, content)

	s.NoError(err)
	s.output.AssertExpectations(s.T())
}

func (s *ProducerTestSuite) TestProducer_Write() {
	content := []*testContent{
		{
			Id:   3,
			Name: "foobar",
		},
		{
			Id:   5,
			Name: "foobaz",
		},
	}

	expectedMsg := []stream.WritableMessage{
		&stream.Message{
			Attributes: map[string]string{
				stream.AttributeEncoding: stream.EncodingJson.String(),
			},
			Body: `{"id":3,"name":"foobar"}`,
		},
		&stream.Message{
			Attributes: map[string]string{
				stream.AttributeEncoding: stream.EncodingJson.String(),
			},
			Body: `{"id":5,"name":"foobaz"}`,
		},
	}

	s.output.On("Write", s.ctx, expectedMsg).Return(nil)
	err := s.producer.Write(s.ctx, content)

	s.NoError(err)
	s.output.AssertExpectations(s.T())
}

func (s *ProducerTestSuite) TestProducer_Write_SliceError() {
	err := s.producer.Write(s.ctx, "string")

	s.EqualError(err, "can not cast models interface to slice: input is not an slice but instead of type string")
}

func (s *ProducerTestSuite) TestProducer_WriteOne_ShouldRetryAndSucceed() {
	content := &testContent{
		Id:   1,
		Name: "retry-success",
	}

	msg := &stream.Message{
		Attributes: map[string]string{
			stream.AttributeEncoding: stream.EncodingJson.String(),
		},
		Body: `{"id":1,"name":"retry-success"}`,
	}

	// set up producer with retry enabled
	s.producer = stream.NewProducerWithInterfaces(nil, s.encoder, s.output, s.retryDaemon, true)

	s.output.On("WriteOne", s.ctx, msg).Return(errors.New("primary output failed"))
	s.retryDaemon.On("RetryOne", s.ctx, msg).Return(nil)

	err := s.producer.WriteOne(s.ctx, content)

	s.NoError(err)
	s.output.AssertExpectations(s.T())
	s.retryDaemon.AssertExpectations(s.T())
}

func (s *ProducerTestSuite) TestProducer_WriteOne_ShouldRetryAndFail() {
	content := &testContent{
		Id:   2,
		Name: "retry-fail",
	}

	msg := &stream.Message{
		Attributes: map[string]string{
			stream.AttributeEncoding: stream.EncodingJson.String(),
		},
		Body: `{"id":2,"name":"retry-fail"}`,
	}

	s.producer = stream.NewProducerWithInterfaces(nil, s.encoder, s.output, s.retryDaemon, true)

	s.output.On("WriteOne", s.ctx, msg).Return(errors.New("primary output failed"))
	s.retryDaemon.On("RetryOne", s.ctx, msg).Return(errors.New("retry failed too"))

	err := s.producer.WriteOne(s.ctx, content)

	s.Error(err)
	s.Contains(err.Error(), "can not write messages to output")
	s.output.AssertExpectations(s.T())
	s.retryDaemon.AssertExpectations(s.T())
}

func (s *ProducerTestSuite) TestProducer_Write_ShouldRetryAndSucceed() {
	content := []*testContent{
		{Id: 1, Name: "a"},
		{Id: 2, Name: "b"},
	}

	msgs := []stream.WritableMessage{
		&stream.Message{
			Attributes: map[string]string{
				stream.AttributeEncoding: stream.EncodingJson.String(),
			},
			Body: `{"id":1,"name":"a"}`,
		},
		&stream.Message{
			Attributes: map[string]string{
				stream.AttributeEncoding: stream.EncodingJson.String(),
			},
			Body: `{"id":2,"name":"b"}`,
		},
	}

	s.producer = stream.NewProducerWithInterfaces(nil, s.encoder, s.output, s.retryDaemon, true)

	s.output.On("Write", s.ctx, msgs).Return(errors.New("write failure"))
	s.retryDaemon.On("RetryMany", s.ctx, msgs).Return(nil)

	err := s.producer.Write(s.ctx, content)

	s.NoError(err)
	s.output.AssertExpectations(s.T())
	s.retryDaemon.AssertExpectations(s.T())
}

func (s *ProducerTestSuite) TestProducer_Write_ShouldRetryAndFail() {
	content := []*testContent{
		{Id: 9, Name: "fail-a"},
		{Id: 10, Name: "fail-b"},
	}

	msgs := []stream.WritableMessage{
		&stream.Message{
			Attributes: map[string]string{
				stream.AttributeEncoding: stream.EncodingJson.String(),
			},
			Body: `{"id":9,"name":"fail-a"}`,
		},
		&stream.Message{
			Attributes: map[string]string{
				stream.AttributeEncoding: stream.EncodingJson.String(),
			},
			Body: `{"id":10,"name":"fail-b"}`,
		},
	}

	s.producer = stream.NewProducerWithInterfaces(nil, s.encoder, s.output, s.retryDaemon, true)

	s.output.On("Write", s.ctx, msgs).Return(errors.New("write failure"))
	s.retryDaemon.On("RetryMany", s.ctx, msgs).Return(errors.New("retry also failed"))

	err := s.producer.Write(s.ctx, content)

	s.Error(err)
	s.Contains(err.Error(), "can not write messages to output")
	s.output.AssertExpectations(s.T())
	s.retryDaemon.AssertExpectations(s.T())
}

func TestProducerTestSuite(t *testing.T) {
	suite.Run(t, new(ProducerTestSuite))
}
