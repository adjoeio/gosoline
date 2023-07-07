package stream

import (
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

const (
	AttributeKafkaMessageMetadata = "KafkaMetadata"
	AttributeKafkaKey             = "KafkaKey"
)

type KafkaMessageMetadata struct {
	Topic         string
	Partition     int
	Offset        int64
	HighWaterMark int64
	Key           []byte
	Time          time.Time
}

func NewKafkaMessageMetadataFromKafkaMessage(message kafka.Message) KafkaMessageMetadata {
	return KafkaMessageMetadata{
		Topic:         message.Topic,
		Partition:     message.Partition,
		Offset:        message.Offset,
		HighWaterMark: message.HighWaterMark,
		Key:           message.Key,
		Time:          message.Time,
	}
}

func (k KafkaMessageMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Time":      k.Time,
		"Partition": k.Partition,
		"Offset":    k.Offset,
		"Key":       string(k.Key),
	})
}

func (k KafkaMessageMetadata) ToKafkaMessageForAck() kafka.Message {
	return kafka.Message{
		Topic:         k.Topic,
		Partition:     k.Partition,
		Offset:        k.Offset,
		HighWaterMark: k.HighWaterMark,
		Key:           k.Key,
		Time:          k.Time,
	}
}

func NewKafkaMessageAttrs(key string) map[string]interface{} {
	return map[string]interface{}{AttributeKafkaKey: key}
}

func KafkaToGosoAttributes(headers []kafka.Header, attributes map[string]interface{}) map[string]interface{} {
	for _, v := range headers {
		attributes[v.Key] = string(v.Value)
	}

	return attributes
}

func KafkaToGosoMessage(k kafka.Message) *Message {
	attributes := KafkaToGosoAttributes(k.Headers, map[string]interface{}{
		AttributeKafkaMessageMetadata: NewKafkaMessageMetadataFromKafkaMessage(k),
	})

	return &Message{Body: string(k.Value), Attributes: attributes}
}

func GosoToAckKafkaMessages(msgs ...*Message) []kafka.Message {
	ks := []kafka.Message{}

	for _, m := range msgs {
		ks = append(ks, m.Attributes[AttributeKafkaMessageMetadata].(KafkaMessageMetadata).ToKafkaMessageForAck())
	}

	return ks
}

func GosoToKafkaMessage(msg *Message) kafka.Message {
	return GosoToAckKafkaMessages(msg)[0]
}

func NewKafkaMessage(writable WritableMessage) kafka.Message {
	gMessage := writable.(*Message)
	kMessage := kafka.Message{Value: []byte(gMessage.Body)}

	key, ok := gMessage.GetAttributes()[AttributeKafkaKey].(string)
	if ok {
		kMessage.Key = []byte(key)
	}

	for k, v := range gMessage.Attributes {
		if k == AttributeKafkaKey {
			continue
		}
		vStr, ok := v.(string)
		if !ok {
			continue
		}

		kMessage.Headers = append(
			kMessage.Headers,
			protocol.Header{Key: k, Value: []byte(vStr)},
		)
	}

	return kMessage
}

func NewKafkaMessages(ms []WritableMessage) []kafka.Message {
	out := []kafka.Message{}
	for _, m := range ms {
		out = append(out, NewKafkaMessage(m))
	}

	return out
}
