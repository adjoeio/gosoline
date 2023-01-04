package connection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	username = "username"
	password = "password"
)

func Test_NewDialer_TLS(t *testing.T) {
	dialer, err := NewDialer(
		&Settings{
			Username:           username,
			Password:           password,
			UseTLS:             true,
			InsecureSkipVerify: true,
		})

	assert.Nil(t, err)

	assert.NotNil(t, dialer.TLS)
	assert.True(t, dialer.TLS.InsecureSkipVerify)
	assert.NotNil(t, dialer.SASLMechanism)
	assert.NotEmpty(t, dialer.TransactionalID)
}

func Test_NewDialer(t *testing.T) {
	dialer, err := NewDialer(
		&Settings{
			Username: username,
			Password: password,
			UseTLS:   false,
		})

	assert.Nil(t, err)
	assert.Nil(t, dialer.TLS)
	assert.NotNil(t, dialer.SASLMechanism)
	assert.NotEmpty(t, dialer.TransactionalID)
}
