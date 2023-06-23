package producer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/segmentio/kafka-go"
)

const (
	RetriesToOpenCircuit = 3
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
	clock                   clock.Clock
	partitionCircuitBreaker sync.Map
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
		partitionCircuitBreaker: sync.Map{},
		settings:                settings,
	}
}

func (b *activePartitionBalancer) OnSuccess(msg kafka.Message) {
	publishedPartition, loaded := b.loadAndDeleteCachedPartition(msg.Topic, string(msg.Key))
	if !loaded {
		return
	}

	// close the circuit if it was open
	b.partitionCircuitBreaker.Delete(circuitBreakerKey(msg.Topic, publishedPartition))
}

func (b *activePartitionBalancer) OnError(msg kafka.Message, _ error) {
	attemptedPartition, loaded := b.loadAndDeleteCachedPartition(msg.Topic, string(msg.Key))
	if !loaded {
		return
	}
	// TODO: figure out if some of the errors should not affect the count
	partitionCircuit, ok := b.partitionCircuitBreaker.Load(circuitBreakerKey(msg.Topic, attemptedPartition))
	if !ok {
		b.partitionCircuitBreaker.Store(circuitBreakerKey(msg.Topic, attemptedPartition), &PartitionCircuitBreaker{
			recentFailures: 1,
			nextRetryAt:    b.clock.Now().Add(b.settings.Delay).UnixMilli(),
		})
		return
	}

	cb := partitionCircuit.(*PartitionCircuitBreaker)
	atomic.AddInt64(&cb.recentFailures, 1)
	atomic.StoreInt64(&cb.nextRetryAt, b.clock.Now().Add(b.settings.Delay).UnixMilli())
}

func (b *activePartitionBalancer) Balance(msg kafka.Message, partitions ...int) (partition int) {
	defer func() { b.cachePartition(msg.Topic, string(msg.Key), partition) }()

	// calculate balancer and check if the circuit breaker is open for the calculated partition, if so select a different one
	partition = b.balancer.Balance(msg, partitions...)
	partitionCircuit, ok := b.partitionCircuitBreaker.Load(circuitBreakerKey(msg.Topic, partition))
	if !ok {
		return partition
	}

	cb := partitionCircuit.(*PartitionCircuitBreaker)
	if b.circuitIsClosed(cb) {
		// circuit is closed proceed with the defaultPartition
		return partition
	}

	if b.canRetry(cb) {
		// enough time has passed so we can retry
		return partition
	}

	// re-balance without the current partition
	var eligiblePartitions = make([]int, 0, len(partitions)-1)
	for _, other := range partitions {
		if other == partition {
			continue
		}
		eligiblePartitions = append(eligiblePartitions, other)
	}

	partition = b.Balance(msg, eligiblePartitions...)

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

func circuitBreakerKey(topic string, partition int) string {
	return fmt.Sprintf("%s[%d]", topic, partition)
}
