package producer

import (
	"context"
	"time"

	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"

	"github.com/segmentio/kafka-go"
)

const (
	// RequireAllReplicas means that ALL nodes in the replica-set must to confirm the write for a write
	// to be considered durable.
	RequireAllReplicas = -1

	// DefaultWriterReadTimeout is how much to wait for reads.
	DefaultWriterReadTimeout = 30 * time.Second

	// DefaultMetadataTTL is the frequency of metadata refreshes.
	DefaultMetadataTTL = 5 * time.Second

	// DefaultIdleTimeout is the period during which an idle connection can be resued.
	DefaultIdleTimeout = 30 * time.Second
)

//go:generate mockery --name Writer --unroll-variadic=False
type Writer interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
	Stats() kafka.WriterStats
	Close() error
}

func NewWriter(
	logger log.Logger,
	dialer *kafka.Dialer,
	settings *Settings,
	opts ...WriterOption,
) (*kafka.Writer, error) {
	// Config.
	conf := &kafka.WriterConfig{
		Brokers: settings.Connection().Bootstrap,
		Dialer:  dialer,

		// Non-batched by default.
		BatchSize: 1,
		Async:     false,

		// Use a safe default for durability.
		RequiredAcks: RequireAllReplicas,
		// MaxAttempts is set to 1 because the retries are handled by the producer.
		// The amount of the retries is configurable via the producer Settings.Retries.
		// NOTE: If MaxAttempts is set to 0, it will default to 10!
		MaxAttempts: 1,

		ReadTimeout: DefaultWriterReadTimeout,

		CompressionCodec: kafka.Snappy.Codec(),

		Logger: func() logging.LoggerWrapper {
			if settings.DebugLogs {
				return logging.NewKafkaLogger(logger).DebugLogger()
			}
			return logging.NewKafkaLogger(logger).NOOPLogger()
		}(),
		ErrorLogger: logging.NewKafkaLogger(logger).ErrorLogger(),
	}
	for _, opt := range opts {
		opt(conf)
	}

	// Transport.
	transport := &kafka.Transport{
		Dial:        dialer.DialFunc,
		SASL:        dialer.SASLMechanism,
		TLS:         dialer.TLS,
		ClientID:    dialer.ClientID,
		DialTimeout: dialer.Timeout,
		IdleTimeout: DefaultIdleTimeout,

		MetadataTTL: DefaultMetadataTTL,
	}

	// Writer.
	writer := &kafka.Writer{
		Addr:      kafka.TCP(conf.Brokers...),
		Transport: transport,

		Topic: conf.Topic,

		ReadTimeout:  conf.ReadTimeout,
		WriteTimeout: conf.WriteTimeout,
		MaxAttempts:  conf.MaxAttempts,

		Async:        conf.Async,
		BatchSize:    conf.BatchSize,
		BatchBytes:   int64(conf.BatchBytes),
		BatchTimeout: conf.BatchTimeout,

		Balancer:     conf.Balancer,
		RequiredAcks: kafka.RequiredAcks(conf.RequiredAcks),

		Compression: kafka.Compression(conf.CompressionCodec.Code()),

		Logger:      conf.Logger,
		ErrorLogger: conf.ErrorLogger,
	}

	return writer, nil
}
