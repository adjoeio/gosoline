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
}

func NewKafkaLogger(logger log.Logger) *KafkaLogger {
	return &KafkaLogger{Logger: logger.WithChannel(KafkaLoggingChannel)}
}

func (l *KafkaLogger) DebugLogger() LoggerWrapper {
	return func(template string, values ...interface{}) {
		l.WithFields(log.Fields{
			"details": fmt.Sprintf(template, values...),
		}).Debug("segmentio kafka-go debug")
	}
}

func (l *KafkaLogger) ErrorLogger() LoggerWrapper {
	return func(template string, values ...interface{}) {
		l.WithFields(log.Fields{
			"error": fmt.Sprintf(template, values...),
		}).Error("segmentio kafka-go error")
	}
}
