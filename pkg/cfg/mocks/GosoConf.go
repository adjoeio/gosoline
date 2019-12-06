// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import cfg "github.com/applike/gosoline/pkg/cfg"
import mock "github.com/stretchr/testify/mock"
import time "time"

// GosoConf is an autogenerated mock type for the GosoConf type
type GosoConf struct {
	mock.Mock
}

// AllKeys provides a mock function with given fields:
func (_m *GosoConf) AllKeys() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// AllSettings provides a mock function with given fields:
func (_m *GosoConf) AllSettings() map[string]interface{} {
	ret := _m.Called()

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func() map[string]interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *GosoConf) Get(_a0 string) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// GetBool provides a mock function with given fields: _a0
func (_m *GosoConf) GetBool(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// GetDuration provides a mock function with given fields: _a0
func (_m *GosoConf) GetDuration(_a0 string) time.Duration {
	ret := _m.Called(_a0)

	var r0 time.Duration
	if rf, ok := ret.Get(0).(func(string) time.Duration); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(time.Duration)
	}

	return r0
}

// GetFloat64 provides a mock function with given fields: _a0
func (_m *GosoConf) GetFloat64(_a0 string) float64 {
	ret := _m.Called(_a0)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetInt provides a mock function with given fields: _a0
func (_m *GosoConf) GetInt(_a0 string) int {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetIntSlice provides a mock function with given fields: _a0
func (_m *GosoConf) GetIntSlice(_a0 string) []int {
	ret := _m.Called(_a0)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// GetString provides a mock function with given fields: _a0
func (_m *GosoConf) GetString(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetStringMap provides a mock function with given fields: key
func (_m *GosoConf) GetStringMap(key string) map[string]interface{} {
	ret := _m.Called(key)

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(string) map[string]interface{}); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// GetStringMapString provides a mock function with given fields: _a0
func (_m *GosoConf) GetStringMapString(_a0 string) map[string]string {
	ret := _m.Called(_a0)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// GetStringSlice provides a mock function with given fields: _a0
func (_m *GosoConf) GetStringSlice(_a0 string) []string {
	ret := _m.Called(_a0)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetTime provides a mock function with given fields: key
func (_m *GosoConf) GetTime(key string) time.Time {
	ret := _m.Called(key)

	var r0 time.Time
	if rf, ok := ret.Get(0).(func(string) time.Time); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// IsSet provides a mock function with given fields: _a0
func (_m *GosoConf) IsSet(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Option provides a mock function with given fields: options
func (_m *GosoConf) Option(options ...cfg.Option) error {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...cfg.Option) error); ok {
		r0 = rf(options...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnmarshalKey provides a mock function with given fields: key, val, opts
func (_m *GosoConf) UnmarshalKey(key string, val interface{}, opts ...cfg.DecoderConfigOption) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, val)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}
