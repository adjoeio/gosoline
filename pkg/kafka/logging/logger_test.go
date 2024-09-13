package logging_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"
	logMocks "github.com/justtrackio/gosoline/pkg/log/mocks"
)

func TestKafkaLogger(t *testing.T) {
	var (
		logger            = new(logMocks.Logger)
		loggerWithChannel = new(logMocks.Logger)
	)
	defer logger.AssertExpectations(t)
	defer loggerWithChannel.AssertExpectations(t)

	logger.On("WithChannel", "stream.kafka").Return(loggerWithChannel).Once()
	loggerWithChannel.On("WithFields", log.Fields{"details": "debug message"}).Return(loggerWithChannel).Once()
	loggerWithChannel.On("WithFields", log.Fields{"error": "error message"}).Return(loggerWithChannel).Once()

	loggerWithChannel.On("Debug", "segmentio kafka-go debug").Once()
	loggerWithChannel.On("Error", "segmentio kafka-go error").Once()

	kLogger := logging.NewKafkaLogger(logger, logging.WithDebugLogging(true))
	kLogger.DebugLogger().Printf("debug message")
	kLogger.ErrorLogger().Printf("error message")
}
