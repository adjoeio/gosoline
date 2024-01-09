package logging

type KafkaLoggerOpt func(*KafkaLogger)

func WithDebugLogging(debugLogs bool) KafkaLoggerOpt {
	return func(logger *KafkaLogger) {
		logger.debugLogs = debugLogs
	}
}
