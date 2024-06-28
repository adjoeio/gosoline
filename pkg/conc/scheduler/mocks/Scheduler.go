// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Scheduler is an autogenerated mock type for the Scheduler type
type Scheduler[T interface{}] struct {
	mock.Mock
}

type Scheduler_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *Scheduler[T]) EXPECT() *Scheduler_Expecter[T] {
	return &Scheduler_Expecter[T]{mock: &_m.Mock}
}

// Run provides a mock function with given fields: ctx
func (_m *Scheduler[T]) Run(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Scheduler_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type Scheduler_Run_Call[T interface{}] struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Scheduler_Expecter[T]) Run(ctx interface{}) *Scheduler_Run_Call[T] {
	return &Scheduler_Run_Call[T]{Call: _e.mock.On("Run", ctx)}
}

func (_c *Scheduler_Run_Call[T]) Run(run func(ctx context.Context)) *Scheduler_Run_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Scheduler_Run_Call[T]) Return(_a0 error) *Scheduler_Run_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Scheduler_Run_Call[T]) RunAndReturn(run func(context.Context) error) *Scheduler_Run_Call[T] {
	_c.Call.Return(run)
	return _c
}

// ScheduleJob provides a mock function with given fields: key, provider
func (_m *Scheduler[T]) ScheduleJob(key string, provider func() (T, error)) (T, error) {
	ret := _m.Called(key, provider)

	if len(ret) == 0 {
		panic("no return value specified for ScheduleJob")
	}

	var r0 T
	var r1 error
	if rf, ok := ret.Get(0).(func(string, func() (T, error)) (T, error)); ok {
		return rf(key, provider)
	}
	if rf, ok := ret.Get(0).(func(string, func() (T, error)) T); ok {
		r0 = rf(key, provider)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(string, func() (T, error)) error); ok {
		r1 = rf(key, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Scheduler_ScheduleJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScheduleJob'
type Scheduler_ScheduleJob_Call[T interface{}] struct {
	*mock.Call
}

// ScheduleJob is a helper method to define mock.On call
//   - key string
//   - provider func()(T , error)
func (_e *Scheduler_Expecter[T]) ScheduleJob(key interface{}, provider interface{}) *Scheduler_ScheduleJob_Call[T] {
	return &Scheduler_ScheduleJob_Call[T]{Call: _e.mock.On("ScheduleJob", key, provider)}
}

func (_c *Scheduler_ScheduleJob_Call[T]) Run(run func(key string, provider func() (T, error))) *Scheduler_ScheduleJob_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func() (T, error)))
	})
	return _c
}

func (_c *Scheduler_ScheduleJob_Call[T]) Return(_a0 T, _a1 error) *Scheduler_ScheduleJob_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Scheduler_ScheduleJob_Call[T]) RunAndReturn(run func(string, func() (T, error)) (T, error)) *Scheduler_ScheduleJob_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewScheduler creates a new instance of Scheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewScheduler[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Scheduler[T] {
	mock := &Scheduler[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
