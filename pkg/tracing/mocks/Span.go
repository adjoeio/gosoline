// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import tracing "github.com/applike/gosoline/pkg/tracing"

// Span is an autogenerated mock type for the Span type
type Span struct {
	mock.Mock
}

// AddAnnotation provides a mock function with given fields: key, value
func (_m *Span) AddAnnotation(key string, value string) {
	_m.Called(key, value)
}

// AddError provides a mock function with given fields: err
func (_m *Span) AddError(err error) {
	_m.Called(err)
}

// AddMetadata provides a mock function with given fields: key, value
func (_m *Span) AddMetadata(key string, value interface{}) {
	_m.Called(key, value)
}

// Finish provides a mock function with given fields:
func (_m *Span) Finish() {
	_m.Called()
}

// GetId provides a mock function with given fields:
func (_m *Span) GetId() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetTrace provides a mock function with given fields:
func (_m *Span) GetTrace() *tracing.Trace {
	ret := _m.Called()

	var r0 *tracing.Trace
	if rf, ok := ret.Get(0).(func() *tracing.Trace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*tracing.Trace)
		}
	}

	return r0
}
