// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// Cache is an autogenerated mock type for the Cache type
type Cache[T interface{}] struct {
	mock.Mock
}

type Cache_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *Cache[T]) EXPECT() *Cache_Expecter[T] {
	return &Cache_Expecter[T]{mock: &_m.Mock}
}

// Contains provides a mock function with given fields: key
func (_m *Cache[T]) Contains(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Contains")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Cache_Contains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Contains'
type Cache_Contains_Call[T interface{}] struct {
	*mock.Call
}

// Contains is a helper method to define mock.On call
//   - key string
func (_e *Cache_Expecter[T]) Contains(key interface{}) *Cache_Contains_Call[T] {
	return &Cache_Contains_Call[T]{Call: _e.mock.On("Contains", key)}
}

func (_c *Cache_Contains_Call[T]) Run(run func(key string)) *Cache_Contains_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Cache_Contains_Call[T]) Return(_a0 bool) *Cache_Contains_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Contains_Call[T]) RunAndReturn(run func(string) bool) *Cache_Contains_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Expire provides a mock function with given fields: key
func (_m *Cache[T]) Expire(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Expire")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Cache_Expire_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Expire'
type Cache_Expire_Call[T interface{}] struct {
	*mock.Call
}

// Expire is a helper method to define mock.On call
//   - key string
func (_e *Cache_Expecter[T]) Expire(key interface{}) *Cache_Expire_Call[T] {
	return &Cache_Expire_Call[T]{Call: _e.mock.On("Expire", key)}
}

func (_c *Cache_Expire_Call[T]) Run(run func(key string)) *Cache_Expire_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Cache_Expire_Call[T]) Return(_a0 bool) *Cache_Expire_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Expire_Call[T]) RunAndReturn(run func(string) bool) *Cache_Expire_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: key
func (_m *Cache[T]) Get(key string) (T, bool) {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 T
	var r1 bool
	if rf, ok := ret.Get(0).(func(string) (T, bool)); ok {
		return rf(key)
	}
	if rf, ok := ret.Get(0).(func(string) T); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(T)
	}

	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Cache_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type Cache_Get_Call[T interface{}] struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - key string
func (_e *Cache_Expecter[T]) Get(key interface{}) *Cache_Get_Call[T] {
	return &Cache_Get_Call[T]{Call: _e.mock.On("Get", key)}
}

func (_c *Cache_Get_Call[T]) Run(run func(key string)) *Cache_Get_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Cache_Get_Call[T]) Return(_a0 T, _a1 bool) *Cache_Get_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Cache_Get_Call[T]) RunAndReturn(run func(string) (T, bool)) *Cache_Get_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Provide provides a mock function with given fields: key, provider
func (_m *Cache[T]) Provide(key string, provider func() T) T {
	ret := _m.Called(key, provider)

	if len(ret) == 0 {
		panic("no return value specified for Provide")
	}

	var r0 T
	if rf, ok := ret.Get(0).(func(string, func() T) T); ok {
		r0 = rf(key, provider)
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

// Cache_Provide_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Provide'
type Cache_Provide_Call[T interface{}] struct {
	*mock.Call
}

// Provide is a helper method to define mock.On call
//   - key string
//   - provider func() T
func (_e *Cache_Expecter[T]) Provide(key interface{}, provider interface{}) *Cache_Provide_Call[T] {
	return &Cache_Provide_Call[T]{Call: _e.mock.On("Provide", key, provider)}
}

func (_c *Cache_Provide_Call[T]) Run(run func(key string, provider func() T)) *Cache_Provide_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func() T))
	})
	return _c
}

func (_c *Cache_Provide_Call[T]) Return(_a0 T) *Cache_Provide_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_Provide_Call[T]) RunAndReturn(run func(string, func() T) T) *Cache_Provide_Call[T] {
	_c.Call.Return(run)
	return _c
}

// ProvideWithError provides a mock function with given fields: key, provider
func (_m *Cache[T]) ProvideWithError(key string, provider func() (T, error)) (T, error) {
	ret := _m.Called(key, provider)

	if len(ret) == 0 {
		panic("no return value specified for ProvideWithError")
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

// Cache_ProvideWithError_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProvideWithError'
type Cache_ProvideWithError_Call[T interface{}] struct {
	*mock.Call
}

// ProvideWithError is a helper method to define mock.On call
//   - key string
//   - provider func()(T , error)
func (_e *Cache_Expecter[T]) ProvideWithError(key interface{}, provider interface{}) *Cache_ProvideWithError_Call[T] {
	return &Cache_ProvideWithError_Call[T]{Call: _e.mock.On("ProvideWithError", key, provider)}
}

func (_c *Cache_ProvideWithError_Call[T]) Run(run func(key string, provider func() (T, error))) *Cache_ProvideWithError_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(func() (T, error)))
	})
	return _c
}

func (_c *Cache_ProvideWithError_Call[T]) Return(_a0 T, _a1 error) *Cache_ProvideWithError_Call[T] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Cache_ProvideWithError_Call[T]) RunAndReturn(run func(string, func() (T, error)) (T, error)) *Cache_ProvideWithError_Call[T] {
	_c.Call.Return(run)
	return _c
}

// Set provides a mock function with given fields: key, value
func (_m *Cache[T]) Set(key string, value T) {
	_m.Called(key, value)
}

// Cache_Set_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Set'
type Cache_Set_Call[T interface{}] struct {
	*mock.Call
}

// Set is a helper method to define mock.On call
//   - key string
//   - value T
func (_e *Cache_Expecter[T]) Set(key interface{}, value interface{}) *Cache_Set_Call[T] {
	return &Cache_Set_Call[T]{Call: _e.mock.On("Set", key, value)}
}

func (_c *Cache_Set_Call[T]) Run(run func(key string, value T)) *Cache_Set_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(T))
	})
	return _c
}

func (_c *Cache_Set_Call[T]) Return() *Cache_Set_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *Cache_Set_Call[T]) RunAndReturn(run func(string, T)) *Cache_Set_Call[T] {
	_c.Call.Return(run)
	return _c
}

// SetX provides a mock function with given fields: key, value, ttl
func (_m *Cache[T]) SetX(key string, value T, ttl time.Duration) {
	_m.Called(key, value, ttl)
}

// Cache_SetX_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetX'
type Cache_SetX_Call[T interface{}] struct {
	*mock.Call
}

// SetX is a helper method to define mock.On call
//   - key string
//   - value T
//   - ttl time.Duration
func (_e *Cache_Expecter[T]) SetX(key interface{}, value interface{}, ttl interface{}) *Cache_SetX_Call[T] {
	return &Cache_SetX_Call[T]{Call: _e.mock.On("SetX", key, value, ttl)}
}

func (_c *Cache_SetX_Call[T]) Run(run func(key string, value T, ttl time.Duration)) *Cache_SetX_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(T), args[2].(time.Duration))
	})
	return _c
}

func (_c *Cache_SetX_Call[T]) Return() *Cache_SetX_Call[T] {
	_c.Call.Return()
	return _c
}

func (_c *Cache_SetX_Call[T]) RunAndReturn(run func(string, T, time.Duration)) *Cache_SetX_Call[T] {
	_c.Call.Return(run)
	return _c
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache[T interface{}](t interface {
	mock.TestingT
	Cleanup(func())
}) *Cache[T] {
	mock := &Cache[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
