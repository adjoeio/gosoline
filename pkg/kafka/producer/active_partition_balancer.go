package producer

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/segmentio/kafka-go"
)

const (
	RetriesToOpenCircuit = 1
	RetryDelay           = time.Minute
)

type PartitionCircuitBreaker struct {
	recentFailures int64
	nextRetryAt    int64
}

type PartitionCircuitBreakerSettings struct {
	Retries int64
	Delay   time.Duration
}

// activePartitionBalancer
// should store the circuit breakers per topics/partition
type activePartitionBalancer struct {
	balancer                kafka.Balancer
	logger                  log.Logger
	clock                   clock.Clock
	partitionCircuitBreaker *topicPartitionCircuitBreakerStore
	publishedPartition      sync.Map
	settings                *PartitionCircuitBreakerSettings
}

func NewActivePartitionHashBalancer() *activePartitionBalancer {
	return NewActivePartitionHashBalancerWithInterfaces(
		&kafka.Hash{},
		clock.NewRealClock(),
		&PartitionCircuitBreakerSettings{
			Retries: RetriesToOpenCircuit,
			Delay:   RetryDelay,
		},
	)
}

func NewActivePartitionHashBalancerWithInterfaces(balancer kafka.Balancer, clock clock.Clock, settings *PartitionCircuitBreakerSettings) *activePartitionBalancer {
	return &activePartitionBalancer{
		balancer:                balancer,
		clock:                   clock,
		partitionCircuitBreaker: NewTopicPartitionCircuitBreakerStore(settings.Retries),
		publishedPartition:      sync.Map{},
		settings:                settings,
	}
}

func (b *activePartitionBalancer) Init(_ context.Context, _ cfg.Config, logger log.Logger) error {
	logger = logger.WithChannel("kafka_balancer")

	b.logger = logger
	if b.partitionCircuitBreaker != nil {
		b.partitionCircuitBreaker.WithLogger(logger)
	}

	return nil
}

func (b *activePartitionBalancer) OnSuccess(msg kafka.Message) {
	publishedPartition, loaded := b.loadAndDeleteCachedPartition(msg.Topic, string(msg.Key))
	if !loaded {
		return
	}

	// close the circuit if it was open
	b.partitionCircuitBreaker.Delete(msg.Topic, publishedPartition)
}

func (b *activePartitionBalancer) OnError(msg kafka.Message, _ error) {
	attemptedPartition, loaded := b.loadAndDeleteCachedPartition(msg.Topic, string(msg.Key))
	if !loaded {
		return
	}
	// TODO: figure out if some of the errors should not affect the count
	cb, ok := b.partitionCircuitBreaker.Load(msg.Topic, attemptedPartition)
	if !ok {
		b.partitionCircuitBreaker.Store(msg.Topic, attemptedPartition, &PartitionCircuitBreaker{
			recentFailures: 1,
			nextRetryAt:    b.clock.Now().Add(b.settings.Delay).UnixMilli(),
		})
		return
	}

	atomic.AddInt64(&cb.recentFailures, 1)
	atomic.StoreInt64(&cb.nextRetryAt, b.clock.Now().Add(b.settings.Delay).UnixMilli())
}

func (b *activePartitionBalancer) Balance(msg kafka.Message, partitions ...int) int {
	partitionIndex := b.balance(msg, partitions...)
	b.cachePartition(msg.Topic, string(msg.Key), partitions[partitionIndex])

	return partitionIndex
}

func (b *activePartitionBalancer) balance(msg kafka.Message, partitions ...int) int {
	// calculate balancer and check if the circuit breaker is open for the calculated partition, if so select a different one
	partitionIndex := b.balancer.Balance(msg, partitions...)
	cb, ok := b.partitionCircuitBreaker.Load(msg.Topic, partitions[partitionIndex])
	if !ok {
		return partitionIndex
	}

	if b.circuitIsClosed(cb) {
		// circuit is closed proceed with the defaultPartition
		return partitionIndex
	}

	if b.canRetry(cb) {
		b.logInfo("retrying %s-%d", msg.Topic, partitionIndex)
		// enough time has passed so we can retry
		return partitionIndex
	} else {
		b.logInfo("no retrying %s-%d, next retry at %d", msg.Topic, partitionIndex, cb.nextRetryAt)
	}

	eligiblePartitions := b.partitionCircuitBreaker.GetActivePartitions(msg.Topic, partitions)
	if len(eligiblePartitions) == 0 {
		b.logInfo("no eligible partitions for topic %s", msg.Topic)
		return 0
	}

	eligiblePartitionIndex := b.balance(msg, eligiblePartitions...)
	for i, v := range partitions {
		if v == eligiblePartitions[eligiblePartitionIndex] {
			partitionIndex = i
			break
		}
	}

	b.logInfo("rebalanced partition %s-%d", msg.Topic, partitionIndex)

	return partitionIndex
}

func (b *activePartitionBalancer) circuitIsClosed(cb *PartitionCircuitBreaker) bool {
	return atomic.LoadInt64(&cb.recentFailures) < b.settings.Retries
}

func (b *activePartitionBalancer) canRetry(cb *PartitionCircuitBreaker) bool {
	nextRetryAt := atomic.LoadInt64(&cb.nextRetryAt)
	now := b.clock.Now().UnixMilli()
	if nextRetryAt > now {
		return false
	}

	return atomic.CompareAndSwapInt64(&cb.nextRetryAt, nextRetryAt, now+b.settings.Delay.Milliseconds())
}

func (b *activePartitionBalancer) cachePartition(topic string, key string, partition int) {
	b.publishedPartition.Store(fmt.Sprintf("%s-%s", topic, key), partition)
}

func (b *activePartitionBalancer) loadAndDeleteCachedPartition(topic string, key string) (int, bool) {
	val, loaded := b.publishedPartition.LoadAndDelete(fmt.Sprintf("%s-%s", topic, key))
	if !loaded {
		return 0, false
	}

	return val.(int), true
}

func (b *activePartitionBalancer) logInfo(fmt string, args ...interface{}) {
	if b.logger == nil {
		return
	}

	b.logger.Info(fmt, args...)
}
