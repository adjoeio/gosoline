package logging

import (
	"context"

	"github.com/justtrackio/gosoline/pkg/log"
)

const (
	KafkaLoggingChannel = "stream.kafka"
)

type KafkaLogger struct {
	log.Logger
	debugLogs bool
}

func NewKafkaLogger(logger log.Logger, opts ...KafkaLoggerOpt) *KafkaLogger {
	kafkaLogger := &KafkaLogger{
		Logger: logger.WithChannel(KafkaLoggingChannel),
	}

	for _, e := range opts {
		e(kafkaLogger)
	}

	return kafkaLogger
}

func (l *KafkaLogger) WithContext(ctx context.Context) log.Logger {
	return &KafkaLogger{
		Logger:    l.Logger.WithContext(ctx),
		debugLogs: l.debugLogs,
	}
}

func (l *KafkaLogger) WithFields(fields log.Fields) log.Logger {
	return &KafkaLogger{
		Logger:    l.Logger.WithFields(fields),
		debugLogs: l.debugLogs,
	}
}

func (logger KafkaLogger) DebugLogger() DebugLoggerWrapper {
	return DebugLoggerWrapper{logger.Logger, logger.debugLogs}
}

func (logger KafkaLogger) ErrorLogger() ErrorLoggerWrapper {
	return ErrorLoggerWrapper{logger.Logger, logger.debugLogs}
}
