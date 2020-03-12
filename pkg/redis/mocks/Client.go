// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import go_redisredis "github.com/go-redis/redis"
import mock "github.com/stretchr/testify/mock"

import time "time"

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BLPop provides a mock function with given fields: _a0, _a1
func (_m *Client) BLPop(_a0 time.Duration, _a1 ...string) ([]string, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	if rf, ok := ret.Get(0).(func(time.Duration, ...string) []string); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Duration, ...string) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Decr provides a mock function with given fields: key
func (_m *Client) Decr(key string) (int64, error) {
	ret := _m.Called(key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecrBy provides a mock function with given fields: key, amount
func (_m *Client) DecrBy(key string, amount int64) (int64, error) {
	ret := _m.Called(key, amount)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, int64) int64); ok {
		r0 = rf(key, amount)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(key, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Del provides a mock function with given fields: _a0
func (_m *Client) Del(_a0 string) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: keys
func (_m *Client) Exists(keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(...string) int64); ok {
		r0 = rf(keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Expire provides a mock function with given fields: key, ttl
func (_m *Client) Expire(key string, ttl time.Duration) (bool, error) {
	ret := _m.Called(key, ttl)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, time.Duration) bool); ok {
		r0 = rf(key, ttl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, time.Duration) error); ok {
		r1 = rf(key, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlushDB provides a mock function with given fields:
func (_m *Client) FlushDB() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: _a0
func (_m *Client) Get(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HExists provides a mock function with given fields: _a0, _a1
func (_m *Client) HExists(_a0 string, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HGet provides a mock function with given fields: _a0, _a1
func (_m *Client) HGet(_a0 string, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HKeys provides a mock function with given fields: _a0
func (_m *Client) HKeys(_a0 string) ([]string, error) {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HMGet provides a mock function with given fields: key, fields
func (_m *Client) HMGet(key string, fields ...string) ([]interface{}, error) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(string, ...string) []interface{}); ok {
		r0 = rf(key, fields...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...string) error); ok {
		r1 = rf(key, fields...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HMSet provides a mock function with given fields: key, pairs
func (_m *Client) HMSet(key string, pairs map[string]interface{}) error {
	ret := _m.Called(key, pairs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]interface{}) error); ok {
		r0 = rf(key, pairs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HSet provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) HSet(_a0 string, _a1 string, _a2 interface{}) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, interface{}) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HSetNX provides a mock function with given fields: key, field, value
func (_m *Client) HSetNX(key string, field string, value interface{}) (bool, error) {
	ret := _m.Called(key, field, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, interface{}) bool); ok {
		r0 = rf(key, field, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, interface{}) error); ok {
		r1 = rf(key, field, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Incr provides a mock function with given fields: key
func (_m *Client) Incr(key string) (int64, error) {
	ret := _m.Called(key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrBy provides a mock function with given fields: key, amount
func (_m *Client) IncrBy(key string, amount int64) (int64, error) {
	ret := _m.Called(key, amount)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, int64) int64); ok {
		r0 = rf(key, amount)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int64) error); ok {
		r1 = rf(key, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAlive provides a mock function with given fields:
func (_m *Client) IsAlive() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LLen provides a mock function with given fields: _a0
func (_m *Client) LLen(_a0 string) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LPop provides a mock function with given fields: _a0
func (_m *Client) LPop(_a0 string) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MGet provides a mock function with given fields: keys
func (_m *Client) MGet(keys ...string) ([]interface{}, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(...string) []interface{}); ok {
		r0 = rf(keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MSet provides a mock function with given fields: pairs
func (_m *Client) MSet(pairs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, pairs...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(pairs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Pipeline provides a mock function with given fields:
func (_m *Client) Pipeline() go_redisredis.Pipeliner {
	ret := _m.Called()

	var r0 go_redisredis.Pipeliner
	if rf, ok := ret.Get(0).(func() go_redisredis.Pipeliner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(go_redisredis.Pipeliner)
		}
	}

	return r0
}

// RPush provides a mock function with given fields: _a0, _a1
func (_m *Client) RPush(_a0 string, _a1 ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string, ...interface{}) int64); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) Set(_a0 string, _a1 interface{}, _a2 time.Duration) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, interface{}, time.Duration) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
