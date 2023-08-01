package logging_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"
	logMocks "github.com/justtrackio/gosoline/pkg/log/mocks"
)

func TestKafkaLogger(t *testing.T) {
	logger := logMocks.NewLoggerMock(logMocks.WithTestingT(t))

	logger.EXPECT().WithChannel("stream.kafka").Return(logger).Once()

	logger.EXPECT().WithFields(log.Fields{"details": "debug message"}).Return(logger).Once()
	logger.EXPECT().Debug("segmentio kafka-go debug", []interface{}(nil)).Once()

	logger.EXPECT().WithFields(log.Fields{"error": "error message"}).Return(logger).Once()
	logger.EXPECT().Error("segmentio kafka-go error").Once()

	kLogger := logging.NewKafkaLogger(logger, logging.WithDebugLogging(true))
	kLogger.DebugLogger().Printf("debug message")
	kLogger.ErrorLogger().Printf("error message")
}
