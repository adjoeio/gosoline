package connection

import (
	"crypto/tls"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

// NewDialer is a dialer factory.
func NewDialer(conf *Settings) (*kafka.Dialer, error) {
	mechanism, err := scram.Mechanism(scram.SHA512, conf.Username, conf.Password)
	if err != nil {
		return nil, err
	}

	dialer := &kafka.Dialer{
		DualStack: true,

		SASLMechanism:   mechanism,
		KeepAlive:       conf.KeepAlive,
		Timeout:         conf.Timeout,
		TransactionalID: uuid.New().String(),
	}
	if conf.UseTLS {
		dialer.TLS = &tls.Config{
			InsecureSkipVerify: conf.InsecureSkipVerify,
			MinVersion:         tls.VersionTLS12,
		}
	}

	return dialer, nil
}
