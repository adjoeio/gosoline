package producer

import "github.com/segmentio/kafka-go"

const defaultKafkaBalancer = "default"

type KafkaBalancer interface {
	kafka.Balancer

	OnSuccess(msg kafka.Message)
	OnError(kafka.Message, error)
}

var kafkaBalancers = map[string]KafkaBalancer{
	defaultKafkaBalancer:        NewKafkaHashBalancer(),
	"active_partition_balancer": NewActivePartitionHashBalancer(),
}

func AddBalancer(key string, balancer KafkaBalancer) {
	kafkaBalancers[key] = balancer
}
