package logging

import (
	"fmt"

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

func (l *KafkaLogger) Debug(format string, args ...interface{}) {
	if l.debugLogs {
		l.Logger.Debug(format, args)
	}
}

func (l *KafkaLogger) DebugLogger() LoggerWrapper {
	if l.debugLogs {
		return func(template string, values ...interface{}) {
			l.WithFields(log.Fields{
				"details": fmt.Sprintf(template, values...),
			}).Debug("segmentio kafka-go debug")
		}
	}
	return func(_ string, _ ...interface{}) {
	}
}

func (l *KafkaLogger) ErrorLogger() LoggerWrapper {
	return func(template string, values ...interface{}) {
		l.WithFields(log.Fields{
			"error": fmt.Sprintf(template, values...),
		}).Error("segmentio kafka-go error")
	}
}
