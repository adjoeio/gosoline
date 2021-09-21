// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	stream "github.com/justtrackio/gosoline/pkg/stream"
	mock "github.com/stretchr/testify/mock"
)

// ProducerDaemonBatcher is an autogenerated mock type for the ProducerDaemonBatcher type
type ProducerDaemonBatcher struct {
	mock.Mock
}

// Append provides a mock function with given fields: msg
func (_m *ProducerDaemonBatcher) Append(msg *stream.Message) ([]stream.WritableMessage, error) {
	ret := _m.Called(msg)

	var r0 []stream.WritableMessage
	if rf, ok := ret.Get(0).(func(*stream.Message) []stream.WritableMessage); ok {
		r0 = rf(msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stream.WritableMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*stream.Message) error); ok {
		r1 = rf(msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Flush provides a mock function with given fields:
func (_m *ProducerDaemonBatcher) Flush() []stream.WritableMessage {
	ret := _m.Called()

	var r0 []stream.WritableMessage
	if rf, ok := ret.Get(0).(func() []stream.WritableMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]stream.WritableMessage)
		}
	}

	return r0
}
