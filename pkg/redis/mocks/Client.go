// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	redis "github.com/justtrackio/gosoline/pkg/redis"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// BLPop provides a mock function with given fields: ctx, timeout, keys
func (_m *Client) BLPop(ctx context.Context, timeout time.Duration, keys ...string) ([]string, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, timeout)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration, ...string) []string); ok {
		r0 = rf(ctx, timeout, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, time.Duration, ...string) error); ok {
		r1 = rf(ctx, timeout, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DBSize provides a mock function with given fields: ctx
func (_m *Client) DBSize(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Decr provides a mock function with given fields: ctx, key
func (_m *Client) Decr(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecrBy provides a mock function with given fields: ctx, key, amount
func (_m *Client) DecrBy(ctx context.Context, key string, amount int64) (int64, error) {
	ret := _m.Called(ctx, key, amount)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) int64); ok {
		r0 = rf(ctx, key, amount)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, key, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Del provides a mock function with given fields: ctx, keys
func (_m *Client) Del(ctx context.Context, keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, ...string) int64); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields: ctx, keys
func (_m *Client) Exists(ctx context.Context, keys ...string) (int64, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, ...string) int64); ok {
		r0 = rf(ctx, keys...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Expire provides a mock function with given fields: ctx, key, ttl
func (_m *Client) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	ret := _m.Called(ctx, key, ttl)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Duration) bool); ok {
		r0 = rf(ctx, key, ttl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Duration) error); ok {
		r1 = rf(ctx, key, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FlushDB provides a mock function with given fields: ctx
func (_m *Client) FlushDB(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, key
func (_m *Client) Get(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HDel provides a mock function with given fields: ctx, key, fields
func (_m *Client) HDel(ctx context.Context, key string, fields ...string) (int64, error) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) int64); ok {
		r0 = rf(ctx, key, fields...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, key, fields...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HExists provides a mock function with given fields: ctx, key, field
func (_m *Client) HExists(ctx context.Context, key string, field string) (bool, error) {
	ret := _m.Called(ctx, key, field)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, key, field)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, field)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HGet provides a mock function with given fields: ctx, key, field
func (_m *Client) HGet(ctx context.Context, key string, field string) (string, error) {
	ret := _m.Called(ctx, key, field)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, key, field)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, key, field)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HGetAll provides a mock function with given fields: ctx, key
func (_m *Client) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	ret := _m.Called(ctx, key)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(context.Context, string) map[string]string); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HKeys provides a mock function with given fields: ctx, key
func (_m *Client) HKeys(ctx context.Context, key string) ([]string, error) {
	ret := _m.Called(ctx, key)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HMGet provides a mock function with given fields: ctx, key, fields
func (_m *Client) HMGet(ctx context.Context, key string, fields ...string) ([]interface{}, error) {
	_va := make([]interface{}, len(fields))
	for _i := range fields {
		_va[_i] = fields[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, string, ...string) []interface{}); ok {
		r0 = rf(ctx, key, fields...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...string) error); ok {
		r1 = rf(ctx, key, fields...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HMSet provides a mock function with given fields: ctx, key, pairs
func (_m *Client) HMSet(ctx context.Context, key string, pairs map[string]interface{}) error {
	ret := _m.Called(ctx, key, pairs)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]interface{}) error); ok {
		r0 = rf(ctx, key, pairs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HSet provides a mock function with given fields: ctx, key, field, value
func (_m *Client) HSet(ctx context.Context, key string, field string, value interface{}) error {
	ret := _m.Called(ctx, key, field, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) error); ok {
		r0 = rf(ctx, key, field, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HSetNX provides a mock function with given fields: ctx, key, field, value
func (_m *Client) HSetNX(ctx context.Context, key string, field string, value interface{}) (bool, error) {
	ret := _m.Called(ctx, key, field, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string, interface{}) bool); ok {
		r0 = rf(ctx, key, field, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, interface{}) error); ok {
		r1 = rf(ctx, key, field, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Incr provides a mock function with given fields: ctx, key
func (_m *Client) Incr(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrBy provides a mock function with given fields: ctx, key, amount
func (_m *Client) IncrBy(ctx context.Context, key string, amount int64) (int64, error) {
	ret := _m.Called(ctx, key, amount)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) int64); ok {
		r0 = rf(ctx, key, amount)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int64) error); ok {
		r1 = rf(ctx, key, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAlive provides a mock function with given fields: ctx
func (_m *Client) IsAlive(ctx context.Context) bool {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// LLen provides a mock function with given fields: ctx, key
func (_m *Client) LLen(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LPop provides a mock function with given fields: ctx, key
func (_m *Client) LPop(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MGet provides a mock function with given fields: ctx, keys
func (_m *Client) MGet(ctx context.Context, keys ...string) ([]interface{}, error) {
	_va := make([]interface{}, len(keys))
	for _i := range keys {
		_va[_i] = keys[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, ...string) []interface{}); ok {
		r0 = rf(ctx, keys...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...string) error); ok {
		r1 = rf(ctx, keys...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MSet provides a mock function with given fields: ctx, pairs
func (_m *Client) MSet(ctx context.Context, pairs ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, pairs...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...interface{}) error); ok {
		r0 = rf(ctx, pairs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Pipeline provides a mock function with given fields:
func (_m *Client) Pipeline() redis.Pipeliner {
	ret := _m.Called()

	var r0 redis.Pipeliner
	if rf, ok := ret.Get(0).(func() redis.Pipeliner); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(redis.Pipeliner)
		}
	}

	return r0
}

// RPush provides a mock function with given fields: ctx, key, values
func (_m *Client) RPush(ctx context.Context, key string, values ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) int64); ok {
		r0 = rf(ctx, key, values...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, key, values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SAdd provides a mock function with given fields: ctx, key, values
func (_m *Client) SAdd(ctx context.Context, key string, values ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, ctx, key)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, ...interface{}) int64); ok {
		r0 = rf(ctx, key, values...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...interface{}) error); ok {
		r1 = rf(ctx, key, values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SCard provides a mock function with given fields: ctx, key
func (_m *Client) SCard(ctx context.Context, key string) (int64, error) {
	ret := _m.Called(ctx, key)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SIsMember provides a mock function with given fields: ctx, key, value
func (_m *Client) SIsMember(ctx context.Context, key string, value interface{}) (bool, error) {
	ret := _m.Called(ctx, key, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) bool); ok {
		r0 = rf(ctx, key, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, key, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, key, value, ttl
func (_m *Client) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	ret := _m.Called(ctx, key, value, ttl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetNX provides a mock function with given fields: ctx, key, value, ttl
func (_m *Client) SetNX(ctx context.Context, key string, value interface{}, ttl time.Duration) (bool, error) {
	ret := _m.Called(ctx, key, value, ttl)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) bool); ok {
		r0 = rf(ctx, key, value, ttl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}, time.Duration) error); ok {
		r1 = rf(ctx, key, value, ttl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
