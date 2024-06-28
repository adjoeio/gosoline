// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddb "github.com/justtrackio/gosoline/pkg/ddb"

	expression "github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"

	mock "github.com/stretchr/testify/mock"
)

// DeleteItemBuilder is an autogenerated mock type for the DeleteItemBuilder type
type DeleteItemBuilder struct {
	mock.Mock
}

type DeleteItemBuilder_Expecter struct {
	mock *mock.Mock
}

func (_m *DeleteItemBuilder) EXPECT() *DeleteItemBuilder_Expecter {
	return &DeleteItemBuilder_Expecter{mock: &_m.Mock}
}

// Build provides a mock function with given fields: item
func (_m *DeleteItemBuilder) Build(item interface{}) (*dynamodb.DeleteItemInput, error) {
	ret := _m.Called(item)

	if len(ret) == 0 {
		panic("no return value specified for Build")
	}

	var r0 *dynamodb.DeleteItemInput
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (*dynamodb.DeleteItemInput, error)); ok {
		return rf(item)
	}
	if rf, ok := ret.Get(0).(func(interface{}) *dynamodb.DeleteItemInput); ok {
		r0 = rf(item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.DeleteItemInput)
		}
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteItemBuilder_Build_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Build'
type DeleteItemBuilder_Build_Call struct {
	*mock.Call
}

// Build is a helper method to define mock.On call
//   - item interface{}
func (_e *DeleteItemBuilder_Expecter) Build(item interface{}) *DeleteItemBuilder_Build_Call {
	return &DeleteItemBuilder_Build_Call{Call: _e.mock.On("Build", item)}
}

func (_c *DeleteItemBuilder_Build_Call) Run(run func(item interface{})) *DeleteItemBuilder_Build_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DeleteItemBuilder_Build_Call) Return(_a0 *dynamodb.DeleteItemInput, _a1 error) *DeleteItemBuilder_Build_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DeleteItemBuilder_Build_Call) RunAndReturn(run func(interface{}) (*dynamodb.DeleteItemInput, error)) *DeleteItemBuilder_Build_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnAllOld provides a mock function with given fields:
func (_m *DeleteItemBuilder) ReturnAllOld() ddb.DeleteItemBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnAllOld")
	}

	var r0 ddb.DeleteItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.DeleteItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.DeleteItemBuilder)
		}
	}

	return r0
}

// DeleteItemBuilder_ReturnAllOld_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnAllOld'
type DeleteItemBuilder_ReturnAllOld_Call struct {
	*mock.Call
}

// ReturnAllOld is a helper method to define mock.On call
func (_e *DeleteItemBuilder_Expecter) ReturnAllOld() *DeleteItemBuilder_ReturnAllOld_Call {
	return &DeleteItemBuilder_ReturnAllOld_Call{Call: _e.mock.On("ReturnAllOld")}
}

func (_c *DeleteItemBuilder_ReturnAllOld_Call) Run(run func()) *DeleteItemBuilder_ReturnAllOld_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeleteItemBuilder_ReturnAllOld_Call) Return(_a0 ddb.DeleteItemBuilder) *DeleteItemBuilder_ReturnAllOld_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteItemBuilder_ReturnAllOld_Call) RunAndReturn(run func() ddb.DeleteItemBuilder) *DeleteItemBuilder_ReturnAllOld_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnNone provides a mock function with given fields:
func (_m *DeleteItemBuilder) ReturnNone() ddb.DeleteItemBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnNone")
	}

	var r0 ddb.DeleteItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.DeleteItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.DeleteItemBuilder)
		}
	}

	return r0
}

// DeleteItemBuilder_ReturnNone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnNone'
type DeleteItemBuilder_ReturnNone_Call struct {
	*mock.Call
}

// ReturnNone is a helper method to define mock.On call
func (_e *DeleteItemBuilder_Expecter) ReturnNone() *DeleteItemBuilder_ReturnNone_Call {
	return &DeleteItemBuilder_ReturnNone_Call{Call: _e.mock.On("ReturnNone")}
}

func (_c *DeleteItemBuilder_ReturnNone_Call) Run(run func()) *DeleteItemBuilder_ReturnNone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *DeleteItemBuilder_ReturnNone_Call) Return(_a0 ddb.DeleteItemBuilder) *DeleteItemBuilder_ReturnNone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteItemBuilder_ReturnNone_Call) RunAndReturn(run func() ddb.DeleteItemBuilder) *DeleteItemBuilder_ReturnNone_Call {
	_c.Call.Return(run)
	return _c
}

// WithCondition provides a mock function with given fields: cond
func (_m *DeleteItemBuilder) WithCondition(cond expression.ConditionBuilder) ddb.DeleteItemBuilder {
	ret := _m.Called(cond)

	if len(ret) == 0 {
		panic("no return value specified for WithCondition")
	}

	var r0 ddb.DeleteItemBuilder
	if rf, ok := ret.Get(0).(func(expression.ConditionBuilder) ddb.DeleteItemBuilder); ok {
		r0 = rf(cond)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.DeleteItemBuilder)
		}
	}

	return r0
}

// DeleteItemBuilder_WithCondition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithCondition'
type DeleteItemBuilder_WithCondition_Call struct {
	*mock.Call
}

// WithCondition is a helper method to define mock.On call
//   - cond expression.ConditionBuilder
func (_e *DeleteItemBuilder_Expecter) WithCondition(cond interface{}) *DeleteItemBuilder_WithCondition_Call {
	return &DeleteItemBuilder_WithCondition_Call{Call: _e.mock.On("WithCondition", cond)}
}

func (_c *DeleteItemBuilder_WithCondition_Call) Run(run func(cond expression.ConditionBuilder)) *DeleteItemBuilder_WithCondition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(expression.ConditionBuilder))
	})
	return _c
}

func (_c *DeleteItemBuilder_WithCondition_Call) Return(_a0 ddb.DeleteItemBuilder) *DeleteItemBuilder_WithCondition_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteItemBuilder_WithCondition_Call) RunAndReturn(run func(expression.ConditionBuilder) ddb.DeleteItemBuilder) *DeleteItemBuilder_WithCondition_Call {
	_c.Call.Return(run)
	return _c
}

// WithHash provides a mock function with given fields: hashValue
func (_m *DeleteItemBuilder) WithHash(hashValue interface{}) ddb.DeleteItemBuilder {
	ret := _m.Called(hashValue)

	if len(ret) == 0 {
		panic("no return value specified for WithHash")
	}

	var r0 ddb.DeleteItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.DeleteItemBuilder); ok {
		r0 = rf(hashValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.DeleteItemBuilder)
		}
	}

	return r0
}

// DeleteItemBuilder_WithHash_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithHash'
type DeleteItemBuilder_WithHash_Call struct {
	*mock.Call
}

// WithHash is a helper method to define mock.On call
//   - hashValue interface{}
func (_e *DeleteItemBuilder_Expecter) WithHash(hashValue interface{}) *DeleteItemBuilder_WithHash_Call {
	return &DeleteItemBuilder_WithHash_Call{Call: _e.mock.On("WithHash", hashValue)}
}

func (_c *DeleteItemBuilder_WithHash_Call) Run(run func(hashValue interface{})) *DeleteItemBuilder_WithHash_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DeleteItemBuilder_WithHash_Call) Return(_a0 ddb.DeleteItemBuilder) *DeleteItemBuilder_WithHash_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteItemBuilder_WithHash_Call) RunAndReturn(run func(interface{}) ddb.DeleteItemBuilder) *DeleteItemBuilder_WithHash_Call {
	_c.Call.Return(run)
	return _c
}

// WithRange provides a mock function with given fields: rangeValue
func (_m *DeleteItemBuilder) WithRange(rangeValue interface{}) ddb.DeleteItemBuilder {
	ret := _m.Called(rangeValue)

	if len(ret) == 0 {
		panic("no return value specified for WithRange")
	}

	var r0 ddb.DeleteItemBuilder
	if rf, ok := ret.Get(0).(func(interface{}) ddb.DeleteItemBuilder); ok {
		r0 = rf(rangeValue)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.DeleteItemBuilder)
		}
	}

	return r0
}

// DeleteItemBuilder_WithRange_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithRange'
type DeleteItemBuilder_WithRange_Call struct {
	*mock.Call
}

// WithRange is a helper method to define mock.On call
//   - rangeValue interface{}
func (_e *DeleteItemBuilder_Expecter) WithRange(rangeValue interface{}) *DeleteItemBuilder_WithRange_Call {
	return &DeleteItemBuilder_WithRange_Call{Call: _e.mock.On("WithRange", rangeValue)}
}

func (_c *DeleteItemBuilder_WithRange_Call) Run(run func(rangeValue interface{})) *DeleteItemBuilder_WithRange_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *DeleteItemBuilder_WithRange_Call) Return(_a0 ddb.DeleteItemBuilder) *DeleteItemBuilder_WithRange_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DeleteItemBuilder_WithRange_Call) RunAndReturn(run func(interface{}) ddb.DeleteItemBuilder) *DeleteItemBuilder_WithRange_Call {
	_c.Call.Return(run)
	return _c
}

// NewDeleteItemBuilder creates a new instance of DeleteItemBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDeleteItemBuilder(t interface {
	mock.TestingT
	Cleanup(func())
}) *DeleteItemBuilder {
	mock := &DeleteItemBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
