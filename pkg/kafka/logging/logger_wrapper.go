package logging

import (
	"fmt"
	"strings"

	"github.com/justtrackio/gosoline/pkg/funk"
	"github.com/justtrackio/gosoline/pkg/log"
)

var nonCriticalErrors = []string{
	// the clients metadata might be outdated if the partition leader has meanwhile changed.
	// the kafka library will log an error in this case before retrying with the proper leader but
	// this is not critical unless the max retry attempts are exceeded and in that case
	// the library will actually return the error instead of logging it via the provided logger.
	"Not Leader For Partition",
	"not the leader",
}

func isNonCriticalError(msg string) bool {
	return funk.ContainsFunc(nonCriticalErrors, func(nonCriticalError string) bool {
		return strings.Contains(msg, nonCriticalError)
	})
}

type DebugLoggerWrapper KafkaLogger

func (logger DebugLoggerWrapper) Printf(msg string, args ...any) {
	logger.WithFields(log.Fields{
		"details": fmt.Sprintf(msg, args...),
	}).Debug("segmentio kafka-go debug")
}

type ErrorLoggerWrapper KafkaLogger

func (logger ErrorLoggerWrapper) Printf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)

	if isNonCriticalError(msg) {
		logger.WithFields(log.Fields{
			"error": fmt.Sprintf(format, args...),
		}).Info("segmentio kafka-go error")

		return
	}

	logger.WithFields(log.Fields{
		"error": fmt.Sprintf(format, args...),
	}).Error("segmentio kafka-go error")
}
