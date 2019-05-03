package stream

import (
	"context"
	"fmt"
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/applike/gosoline/pkg/kernel"
	"github.com/applike/gosoline/pkg/mon"
	"github.com/applike/gosoline/pkg/tracing"
	"gopkg.in/tomb.v2"
	"sync"
	"time"
)

//go:generate mockery -name=BatchConsumerCallback
type BatchConsumerCallback interface {
	Boot(config cfg.Config, logger mon.Logger)
	Consume(ctx context.Context, msg []*Message) error
}

type baseBatchConsumer struct {
	kernel.ForegroundModule

	logger mon.Logger
	tracer tracing.Tracer

	input  Input
	tmb    tomb.Tomb
	ticker *time.Ticker

	name     string
	callback BatchConsumerCallback

	m         sync.Mutex
	processed int
	batchSize int
	batch     []*Message
	force     bool
}

func (c *baseBatchConsumer) Boot(config cfg.Config, logger mon.Logger) error {
	c.logger = logger
	c.tracer = tracing.NewAwsTracer(config)

	idleTimeout := config.GetDuration("consumer_idle_timeout")
	c.ticker = time.NewTicker(idleTimeout * time.Second)

	c.batchSize = config.GetInt("consumer_batch_size")
	c.batch = make([]*Message, 0, c.batchSize)

	c.callback.Boot(config, logger)

	return nil
}

func (c *baseBatchConsumer) Run(ctx context.Context) error {
	defer c.logger.Info("leaving consumer ", c.name)

	c.tmb.Go(c.input.Run)
	c.tmb.Go(c.consume)

	for {
		select {
		case <-ctx.Done():
			c.input.Stop()
			return c.tmb.Wait()

		case <-c.tmb.Dead():
			return c.tmb.Err()

		case <-c.ticker.C:
			c.consumeBatch()

			c.logger.Info(fmt.Sprintf("processed %v messages", c.processed))
			c.processed = 0
		}
	}
}

func (c *baseBatchConsumer) consume() error {
	for {
		msg, ok := <-c.input.Data()

		if !ok {
			return nil
		}

		c.m.Lock()
		c.batch = append(c.batch, msg)
		c.m.Unlock()

		if len(c.batch) < c.batchSize {
			continue
		}

		c.consumeBatch()
	}
}

func (c *baseBatchConsumer) consumeBatch() {
	if len(c.batch) == 0 {
		return
	}

	c.m.Lock()
	defer c.m.Unlock()

	_ = c.callback.Consume(context.Background(), c.batch)

	c.processed += len(c.batch)
	c.batch = make([]*Message, 0, c.batchSize)
}
