// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddb "github.com/justtrackio/gosoline/pkg/ddb"

	mock "github.com/stretchr/testify/mock"
)

// GetItemBuilder is an autogenerated mock type for the GetItemBuilder type
type GetItemBuilder struct {
	mock.Mock
}

// Build provides a mock function with given fields: result
func (_m *GetItemBuilder) Build(result interface{}) (*dynamodb.GetItemInput, error) {
	ret := _m.Called(result)

	var r0 *dynamodb.GetItemInput
	if rf, ok := ret.Get(0).(func(interface{}) *dynamodb.GetItemInput); ok {
		r0 = rf(result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.GetItemInput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisableTtlFilter provides a mock function with given fields:
func (_m *GetItemBuilder) DisableTtlFilter() ddb.GetItemBuilder {
	ret := _m.Called()

	var r0 ddb.GetItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.GetItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.GetItemBuilder)
		}
	}

	return r0
}

// WithConsistentRead provides a mock function with given fields: consistentRead
func (_m *GetItemBuilder) WithConsistentRead(consistentRead bool) ddb.GetItemBuilder {
	ret := _m.Called(consistentRead)

	var r0 ddb.GetItemBuilder
	if rf, ok := ret.Get(0).(func(bool) ddb.GetItemBuilder); ok {
		r0 = rf(consistentRead)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.GetItemBuilder)
		}
	}

	return r0
}

// WithHash provides a mock function with given fields: hashValue
func (_m *GetItemBuilder) WithHash(hashValue interface{}) ddb.GetItemBuilder {
	ret := _m.Called(hashValue)

	var r0 ddb.GetItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.GetItemBuilder); ok {
		r0 = rf(hashValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.GetItemBuilder)
		}
	}

	return r0
}

// WithKeys provides a mock function with given fields: keys
func (_m *GetItemBuilder) WithKeys(keys ...interface{}) ddb.GetItemBuilder {
	var _ca []interface{}
	_ca = append(_ca, keys...)
	ret := _m.Called(_ca...)

	var r0 ddb.GetItemBuilder
	if rf, ok := ret.Get(0).(func(...interface{}) ddb.GetItemBuilder); ok {
		r0 = rf(keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.GetItemBuilder)
		}
	}

	return r0
}

// WithProjection provides a mock function with given fields: rangeValue
func (_m *GetItemBuilder) WithProjection(rangeValue interface{}) ddb.GetItemBuilder {
	ret := _m.Called(rangeValue)

	var r0 ddb.GetItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.GetItemBuilder); ok {
		r0 = rf(rangeValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.GetItemBuilder)
		}
	}

	return r0
}

// WithRange provides a mock function with given fields: rangeValue
func (_m *GetItemBuilder) WithRange(rangeValue interface{}) ddb.GetItemBuilder {
	ret := _m.Called(rangeValue)

	var r0 ddb.GetItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.GetItemBuilder); ok {
		r0 = rf(rangeValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.GetItemBuilder)
		}
	}

	return r0
}
