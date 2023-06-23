package producer

import "github.com/segmentio/kafka-go"

type kafkaHashBalancer struct {
	*kafka.Hash
}

func NewKafkaHashBalancer() *kafkaHashBalancer {
	return &kafkaHashBalancer{
		Hash: &kafka.Hash{},
	}
}

func (b *kafkaHashBalancer) OnSuccess(kafka.Message) {}

func (b *kafkaHashBalancer) OnError(_ kafka.Message, _ error) {}
