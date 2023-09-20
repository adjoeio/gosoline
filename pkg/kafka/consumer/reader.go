package consumer

import (
	"context"
	"time"

	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/segmentio/kafka-go"
)

const (
	// DefaultMaxRetryAttempts is how many times to retry a failed operation.
	DefaultMaxRetryAttempts = 15

	// DefaultMinRetryWaitTime is lower limit for backoff interval.
	DefaultMinRetryWaitTime = 250 * time.Millisecond

	// DefaultMaxRetryWaitTime is upper limit for backoff interval.
	DefaultMaxRetryWaitTime = 2 * time.Second

	// DefaultConsumerGroupRetentionTime is the retention period of current offsets.
	DefaultConsumerGroupRetentionTime = 7 * 24 * time.Hour

	// DefaultMaxWait is a reasonable minimum for MaxWait.
	DefaultMaxWait = time.Second

	// CommitOffsetsSync == 0 means that auto-commit is disabled.
	CommitOffsetsSync = time.Duration(0)
)

//go:generate mockery --name Reader
type Reader interface {
	FetchMessage(context.Context) (kafka.Message, error)
	ReadMessage(context.Context) (kafka.Message, error)
	CommitMessages(context.Context, ...kafka.Message) error
	Stats() kafka.ReaderStats
	Close() error
}

func NewReader(
	logger log.Logger,
	dialer *kafka.Dialer,
	conf *Settings,
	opts ...ReaderOption,
) (*kafka.Reader, error) {
	c := &kafka.ReaderConfig{
		Brokers: conf.Connection().Bootstrap,
		Dialer:  dialer,

		// Topic.
		Topic:                 conf.FQTopic,
		GroupID:               conf.FQGroupID,
		WatchPartitionChanges: true,

		// No batching by default.
		MinBytes: 1,
		MaxBytes: MaxBatchBytes,
		MaxWait:  DefaultMaxWait,

		// Explicit commits.
		CommitInterval: CommitOffsetsSync,

		// Safe defaults.
		RetentionTime:  DefaultConsumerGroupRetentionTime,
		MaxAttempts:    DefaultMaxRetryAttempts,
		ReadBackoffMin: DefaultMinRetryWaitTime,
		ReadBackoffMax: DefaultMaxRetryWaitTime,

		IsolationLevel: kafka.ReadCommitted,

		StartOffset: kafka.LastOffset,

		Logger: func() logging.LoggerWrapper {
			// Sets logger if debug logs are enabled.
			if conf.DebugLogs {
				return logging.NewKafkaLogger(logger).DebugLogger()
			}
			return logging.NewKafkaLogger(logger).NOOPLogger()
		}(),
		ErrorLogger: logging.NewKafkaLogger(logger).ErrorLogger(),
	}

	for _, opt := range opts {
		opt(c)
	}

	if err := c.Validate(); err != nil {
		return nil, err
	}

	return kafka.NewReader(*c), nil
}
