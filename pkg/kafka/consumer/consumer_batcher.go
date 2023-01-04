package consumer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

var _ Batcher = &batcher{}

const (
	minBatchSize    = 1
	minBatchTimeout = 10 * time.Millisecond
)

//go:generate mockery --name Batcher --unroll-variadic=False
type Batcher interface {
	Get(ctx context.Context) []kafka.Message
}

type batcher struct {
	input chan kafka.Message

	BatchSize    int
	BatchTimeout time.Duration
}

func NewBatcher(input chan kafka.Message, size int, timeout time.Duration) *batcher {
	if size < minBatchSize {
		size = minBatchSize
	}
	if timeout < minBatchTimeout {
		timeout = minBatchTimeout
	}

	return &batcher{input: input, BatchSize: size, BatchTimeout: timeout}
}

func (b *batcher) Get(ctx context.Context) []kafka.Message {
	ticker := time.NewTicker(b.BatchTimeout)
	defer ticker.Stop()

	var (
		batch     = []kafka.Message{}
		processed = map[Offset]bool{}
	)

OUT:
	for len(batch) < b.BatchSize {
		select {
		case m := <-b.input:
			if _, ok := processed[Offset{Partition: m.Partition, Index: m.Offset}]; ok {
				// skip already processed message
				continue
			}
			processed[Offset{Partition: m.Partition, Index: m.Offset}] = true
			batch = append(batch, m)
		case <-ticker.C:
			// Batch timeout has elapsed.
			if len(batch) > 0 {
				// A non-empty batch has already been compiled.
				break OUT
			}
		case <-ctx.Done():
			// Shutting down, return compiled messages so far.
			break OUT
		}
	}

	return batch
}
