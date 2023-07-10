package producer_test

import (
	"errors"
	"testing"
	"time"

	"github.com/justtrackio/gosoline/pkg/clock"
	"github.com/justtrackio/gosoline/pkg/kafka/producer"
	logMocks "github.com/justtrackio/gosoline/pkg/log/mocks"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

var failureErr = errors.New("fail")

type FakeBalancer struct {
	ReturnedPartitionMap map[string]int
	WhenMissing          int
}

func (f *FakeBalancer) Balance(msg kafka.Message, partitions ...int) (partition int) {
	partition = f.ReturnedPartitionMap[string(msg.Key)]

	for _, p := range partitions {
		if p == partition {
			return partition
		}
	}

	return f.WhenMissing
}

func Test_activePartitionHashBalancer_Balance(t *testing.T) {
	balancer := &FakeBalancer{
		ReturnedPartitionMap: map[string]int{
			"0": 0,
			"1": 1,
			"2": 2,
			"3": 3,
			"4": 4,
			"5": 5,
		},
		WhenMissing: 3,
	}
	fakeClock := clock.NewFakeClock()
	activePartitionBalancer := producer.NewActivePartitionHashBalancerWithInterfaces(
		balancer,
		fakeClock,
		&producer.PartitionCircuitBreakerSettings{
			Retries: 2,
			Delay:   time.Minute,
		},
	)
	activePartitionBalancer.Init(logMocks.NewLoggerMockedAll())

	partitions := []int{0, 1, 2, 4, 5}
	msg1 := kafkaMessageWithId("1")
	msg2 := kafkaMessageWithId("2")
	assert.Equal(t, 1, activePartitionBalancer.Balance(msg1, partitions...))
	activePartitionBalancer.OnError(msg1, failureErr)
	assert.Equal(t, 1, activePartitionBalancer.Balance(msg1, partitions...))
	activePartitionBalancer.OnError(msg1, failureErr)
	assert.Equal(t, 2, activePartitionBalancer.Balance(msg2, partitions...))
	activePartitionBalancer.OnError(kafkaMessageWithId("2"), failureErr)
	assert.NotEqual(t, 1, activePartitionBalancer.Balance(msg1, partitions...))
	activePartitionBalancer.OnSuccess(msg1)

	fakeClock.Advance(time.Minute * 2)
	assert.Equal(t, 1, activePartitionBalancer.Balance(msg1, partitions...))
	activePartitionBalancer.OnError(msg1, failureErr)
}

func kafkaMessageWithId(id string) kafka.Message {
	return kafka.Message{Topic: "example", Key: []byte(id)}
}
