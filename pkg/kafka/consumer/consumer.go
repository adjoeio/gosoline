package consumer

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/coffin"
	"github.com/justtrackio/gosoline/pkg/kafka/connection"
	"github.com/justtrackio/gosoline/pkg/kafka/logging"
	"github.com/justtrackio/gosoline/pkg/log"
)

type Offset struct {
	Partition int
	Index     int64
}

type Consumer struct {
	logger   log.Logger
	settings *Settings

	pool    coffin.Coffin
	backlog chan kafka.Message
	manager OffsetManager
}

func NewConsumer(
	_ context.Context, conf cfg.Config, logger log.Logger, key string) (*Consumer, error) {
	settings := ParseSettings(conf, key)

	// Connection.
	dialer, err := connection.NewDialer(settings.Connection())
	if err != nil {
		return nil, fmt.Errorf("kafka: failed to get dialer: %w", err)
	}

	// Reader.
	reader, err := NewReader(logger, dialer, settings, getOptions(settings)...)
	if err != nil {
		return nil, fmt.Errorf("kafka: failed to get reader: %w", err)
	}

	manager := NewOffsetManager(logger, reader, settings)

	return NewConsumerWithInterfaces(settings, logger, manager)
}

func NewConsumerWithInterfaces(settings *Settings, logger log.Logger, manager OffsetManager) (*Consumer, error) {
	logger = logger.WithFields(
		log.Fields{
			"kafka_topic":          settings.FQTopic,
			"kafka_consumer_group": settings.FQGroupID,
			"kafka_batch_size":     settings.BatchSize,
			"kafka_max_wait":       settings.BatchTimeout.Milliseconds(),
		},
	)

	return &Consumer{
		settings: settings,
		logger:   logging.NewKafkaLogger(logger, logging.WithDebugLogging(settings.DebugLogs)),
		pool:     coffin.New(),
		backlog:  make(chan kafka.Message, settings.BatchSize),
		manager:  manager,
	}, nil
}

func (c *Consumer) Run(ctx context.Context) error {
	c.logger.Info("starting consumer")
	defer c.logger.Info("shutdown consumer")

	return c.run(c.pool.Context(ctx))
}

func (c *Consumer) Data() chan kafka.Message {
	return c.backlog
}

func (c *Consumer) Commit(ctx context.Context, msgs ...kafka.Message) error {
	return c.manager.Commit(ctx, msgs...)
}

func (c *Consumer) run(ctx context.Context) error {
	c.pool.GoWithContext(ctx, c.manager.Start)

OUT:
	for ctx.Err() == nil {
		for _, msg := range c.manager.Batch(ctx) {
			select {
			case c.backlog <- msg:
			case <-ctx.Done():
				break OUT
			}
		}
	}

	c.pool.Wait()
	return c.pool.Err()
}
