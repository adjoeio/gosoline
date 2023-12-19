package producer

import (
	"sync"

	"github.com/justtrackio/gosoline/pkg/log"
)

type topicPartitionCircuitBreakerStore struct {
	logger              log.Logger
	topicCircuitBreaker sync.Map
	openCircuitAttempts int64
}

func NewTopicPartitionCircuitBreakerStore(openCircuitAttempts int64) *topicPartitionCircuitBreakerStore {
	return &topicPartitionCircuitBreakerStore{
		topicCircuitBreaker: sync.Map{},
		openCircuitAttempts: openCircuitAttempts,
	}
}

func (s *topicPartitionCircuitBreakerStore) WithLogger(logger log.Logger) {
	s.logger = logger
}

func (s *topicPartitionCircuitBreakerStore) Delete(topic string, partition int) {
	cb, ok := s.topicCircuitBreaker.Load(topic)
	if !ok {
		return
	}

	if _, loaded := cb.(*sync.Map).LoadAndDelete(partition); loaded {
		s.logInfo("cb close %s-%d", topic, partition)
	}
}

func (s *topicPartitionCircuitBreakerStore) Load(topic string, partition int) (*PartitionCircuitBreaker, bool) {
	cb, ok := s.topicCircuitBreaker.Load(topic)
	if !ok {
		return nil, ok
	}

	pcb, ok := cb.(*sync.Map).Load(partition)
	if !ok {
		return nil, ok
	}

	return pcb.(*PartitionCircuitBreaker), true
}

func (s *topicPartitionCircuitBreakerStore) Store(topic string, partition int, pcb *PartitionCircuitBreaker) {
	var hMap *sync.Map

	cb, ok := s.topicCircuitBreaker.Load(topic)
	if !ok {
		hMap = &sync.Map{}
		s.topicCircuitBreaker.Store(topic, hMap)
	} else {
		hMap = cb.(*sync.Map)
	}

	hMap.Store(partition, pcb)
	s.logInfo("cb open %s-%d", topic, partition)
}

func (s *topicPartitionCircuitBreakerStore) GetActivePartitions(topic string, partitions []int) []int {
	activePartitions := []int{}

	cb, ok := s.topicCircuitBreaker.Load(topic)
	if !ok {
		s.logInfo("cb active-partitions %s %v", topic, partitions)

		return partitions
	}

	hMap := cb.(*sync.Map)
	inactivePartitions := map[int]bool{}
	hMap.Range(func(partition, val any) bool {
		if val.(*PartitionCircuitBreaker).recentFailures >= s.openCircuitAttempts {
			inactivePartitions[partition.(int)] = true
		}

		return true
	})

	for _, partition := range partitions {
		if _, ok := inactivePartitions[partition]; ok {
			continue
		}
		activePartitions = append(activePartitions, partition)
	}

	s.logInfo("cb active-partitions %s %v", topic, activePartitions)

	return activePartitions
}

func (s *topicPartitionCircuitBreakerStore) logInfo(fmt string, args ...interface{}) {
	if s.logger == nil {
		return
	}

	s.logger.Info(fmt, args...)
}
