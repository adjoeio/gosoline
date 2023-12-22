package log

import "context"

type noopLogger struct{}

func NewNOOPLogger() *noopLogger {
	return &noopLogger{}
}

func (l *noopLogger) Debug(_ string, _ ...interface{}) {
}

func (l *noopLogger) Info(_ string, _ ...interface{}) {
}

func (l *noopLogger) Warn(_ string, _ ...interface{}) {
}

func (l *noopLogger) Error(_ string, _ ...interface{}) {
}

func (l *noopLogger) WithChannel(_ string) Logger {
	return l
}

func (l *noopLogger) WithContext(_ context.Context) Logger {
	return l
}

func (l *noopLogger) WithFields(_ Fields) Logger {
	return l
}
