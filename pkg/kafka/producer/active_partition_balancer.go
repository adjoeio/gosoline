package producer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/segmentio/kafka-go"
)

const (
	RetriesToOpenCircuit = 1
	RetryDelay           = time.Minute
)

type Initable interface {
	Init(logger log.Logger)
}

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
	clock                   clock.Clock
	partitionCircuitBreaker *topicPartitionCircuitBreakerStore
	publishedPartition      sync.Map
	settings                *PartitionCircuitBreakerSettings
	logger                  log.Logger
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

func (b *activePartitionBalancer) Init(logger log.Logger) {
	b.logger = logger
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

func (b *activePartitionBalancer) Balance(msg kafka.Message, partitions ...int) (partition int) {
	defer func() { b.cachePartition(msg.Topic, string(msg.Key), partition) }()
	defer func() { b.logger.Info("returned partition:%d", partition) }()

	// calculate balancer and check if the circuit breaker is open for the calculated partition, if so select a different one
	partition = b.balancer.Balance(msg, partitions...)
	cb, ok := b.partitionCircuitBreaker.Load(msg.Topic, partition)
	if !ok {
		return partition
	}

	b.logger.Info("partition circuit was found for partition:%d", partition)

	if b.circuitIsClosed(cb) {
		// circuit is closed proceed with the defaultPartition
		b.logger.Info("partition circuit was closed for partition:%d", partition)
		return partition
	}

	if b.canRetry(cb) {
		// enough time has passed so we can retry
		b.logger.Info("partition circuit can retry partition:%d", partition)
		return partition
	}

	eligiblePartitions := b.partitionCircuitBreaker.GetActivePartitions(msg.Topic, partitions)

	partition = b.Balance(msg, eligiblePartitions...)
	b.logger.Info("re calculated partition:%d from eligible partitions count:%d", partition, len(eligiblePartitions))

	return partition
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
