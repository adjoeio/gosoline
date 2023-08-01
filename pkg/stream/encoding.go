package stream

import (
	"fmt"
)

type EncodingType string

const (
	EncodingJson     EncodingType = "application/json"
	EncodingProtobuf EncodingType = "application/x-protobuf"
)

func (s EncodingType) String() string {
	return string(s)
}

var _ fmt.Stringer = EncodingType("")

var defaultMessageBodyEncoding = EncodingJson

func WithDefaultMessageBodyEncoding(encoding EncodingType) {
	defaultMessageBodyEncoding = encoding
}

type MessageBodyEncoder interface {
	Encode(data interface{}, attributes map[string]interface{}) ([]byte, error)
	Decode(data []byte, attributes map[string]interface{}, out interface{}) error
}

var messageBodyEncoders = map[EncodingType]MessageBodyEncoder{
	EncodingJson:     new(jsonEncoder),
	EncodingProtobuf: new(protobufEncoder),
}

func AddMessageBodyEncoder(encoding EncodingType, encoder MessageBodyEncoder) {
	messageBodyEncoders[encoding] = encoder
}

func EncodeMessage(encoding EncodingType, data interface{}, attributes map[string]interface{}) ([]byte, error) {
	if encoding == "" {
		return nil, fmt.Errorf("no encoding provided to encode message")
	}

	encoder, ok := messageBodyEncoders[encoding]

	if !ok {
		return nil, fmt.Errorf("there is no message body encoder available for encoding '%s'", encoding)
	}

	body, err := encoder.Encode(data, attributes)
	if err != nil {
		return nil, fmt.Errorf("can not encode message body with encoding '%s': %w", encoding, err)
	}

	return body, nil
}

func DecodeMessage(encoding EncodingType, data []byte, attributes map[string]interface{}, out interface{}) error {
	encoder, ok := messageBodyEncoders[encoding]

	if !ok {
		return fmt.Errorf("there is no message body decoder available for encoding '%s'", encoding)
	}

	err := encoder.Decode(data, attributes, out)
	if err != nil {
		return fmt.Errorf("can not decode message body with encoding '%s': %w", encoding, err)
	}

	return nil
}
