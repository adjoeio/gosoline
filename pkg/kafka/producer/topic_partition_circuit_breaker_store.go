package producer

import "sync"

type topicPartitionCircuitBreakerStore struct {
	topicCircuitBreaker sync.Map
	openCircuitAttempts int64
}

func NewTopicPartitionCircuitBreakerStore(openCircuitAttempts int64) *topicPartitionCircuitBreakerStore {
	return &topicPartitionCircuitBreakerStore{
		topicCircuitBreaker: sync.Map{},
		openCircuitAttempts: openCircuitAttempts,
	}
}

func (s *topicPartitionCircuitBreakerStore) Delete(topic string, partition int) {
	cb, ok := s.topicCircuitBreaker.Load(topic)
	if !ok {
		return
	}

	cb.(*sync.Map).Delete(partition)
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
}

func (s *topicPartitionCircuitBreakerStore) GetActivePartitions(topic string, partitions []int) []int {
	activePartitions := []int{}

	cb, ok := s.topicCircuitBreaker.Load(topic)
	if !ok {
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

	return activePartitions
}
