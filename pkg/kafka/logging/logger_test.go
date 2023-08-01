package logging_test

import (
	"testing"

	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"
	logMocks "github.com/justtrackio/gosoline/pkg/log/mocks"
)

func TestKafkaLogger(t *testing.T) {
	var (
		logger            = logMocks.NewLogger(t)
		loggerWithChannel = logMocks.NewLogger(t)
	)

	logger.EXPECT().WithChannel("stream.kafka").Return(loggerWithChannel).Once()

	loggerWithChannel.EXPECT().WithFields(log.Fields{"details": "debug message"}).Return(loggerWithChannel).Once()
	loggerWithChannel.EXPECT().Debug("segmentio kafka-go debug").Once()

	loggerWithChannel.EXPECT().WithFields(log.Fields{"error": "error message"}).Return(loggerWithChannel).Once()
	loggerWithChannel.EXPECT().Error("segmentio kafka-go error").Once()

	kLogger := logging.NewKafkaLogger(logger)
	kLogger.DebugLogger().Printf("debug message")
	kLogger.ErrorLogger().Printf("error message")
}
