package stream

import (
	"fmt"

	"github.com/spf13/cast"
)

type EncodingType string

const (
	EncodingJson     EncodingType = "application/json"
	EncodingText     EncodingType = "text/plain"
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
	Encode(data interface{}) ([]byte, error)
	Decode(data []byte, out interface{}) error
}

var messageBodyEncoders = map[EncodingType]MessageBodyEncoder{
	EncodingJson:     new(jsonEncoder),
	EncodingProtobuf: new(protobufEncoder),
	EncodingText:     new(textEncoder),
}

func AddMessageBodyEncoder(encoding EncodingType, encoder MessageBodyEncoder) {
	messageBodyEncoders[encoding] = encoder
}

func EncodeMessage(encoding EncodingType, data interface{}) ([]byte, error) {
	if encoding == "" {
		return nil, fmt.Errorf("no encoding provided to encode message")
	}

	encoder, ok := messageBodyEncoders[encoding]

	if !ok {
		return nil, fmt.Errorf("there is no message body encoder available for encoding '%s'", encoding)
	}

	body, err := encoder.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("can not encode message body with encoding '%s': %w", encoding, err)
	}

	return body, nil
}

func DecodeMessage(encoding EncodingType, data []byte, out interface{}) error {
	encoder, ok := messageBodyEncoders[encoding]

	if !ok {
		return fmt.Errorf("there is no message body decoder available for encoding '%s'", encoding)
	}

	err := encoder.Decode(data, out)
	if err != nil {
		return fmt.Errorf("can not decode message body with encoding '%s': %w", encoding, err)
	}

	return nil
}

type textEncoder struct{}

func (e textEncoder) Encode(data interface{}) ([]byte, error) {
	if bts, ok := data.([]byte); ok {
		return bts, nil
	}

	str, err := cast.ToStringE(data)

	return []byte(str), err
}

func (e textEncoder) Decode(data []byte, out interface{}) error {
	bts, ok := out.(*[]byte)
	if !ok {
		return fmt.Errorf("the out parameter of the text decode has to be a pointer to byte slice")
	}

	*bts = append(*bts, data...)

	return nil
}
