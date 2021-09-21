// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	db "github.com/justtrackio/gosoline/pkg/db"
	mock "github.com/stretchr/testify/mock"
)

// QueryBuilder is an autogenerated mock type for the QueryBuilder type
type QueryBuilder struct {
	mock.Mock
}

// GroupBy provides a mock function with given fields: field
func (_m *QueryBuilder) GroupBy(field ...string) db.QueryBuilder {
	_va := make([]interface{}, len(field))
	for _i := range field {
		_va[_i] = field[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.QueryBuilder
	if rf, ok := ret.Get(0).(func(...string) db.QueryBuilder); ok {
		r0 = rf(field...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.QueryBuilder)
		}
	}

	return r0
}

// Joins provides a mock function with given fields: joins
func (_m *QueryBuilder) Joins(joins []string) db.QueryBuilder {
	ret := _m.Called(joins)

	var r0 db.QueryBuilder
	if rf, ok := ret.Get(0).(func([]string) db.QueryBuilder); ok {
		r0 = rf(joins)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.QueryBuilder)
		}
	}

	return r0
}

// OrderBy provides a mock function with given fields: field, direction
func (_m *QueryBuilder) OrderBy(field string, direction string) db.QueryBuilder {
	ret := _m.Called(field, direction)

	var r0 db.QueryBuilder
	if rf, ok := ret.Get(0).(func(string, string) db.QueryBuilder); ok {
		r0 = rf(field, direction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.QueryBuilder)
		}
	}

	return r0
}

// Page provides a mock function with given fields: offset, size
func (_m *QueryBuilder) Page(offset int, size int) db.QueryBuilder {
	ret := _m.Called(offset, size)

	var r0 db.QueryBuilder
	if rf, ok := ret.Get(0).(func(int, int) db.QueryBuilder); ok {
		r0 = rf(offset, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.QueryBuilder)
		}
	}

	return r0
}

// Table provides a mock function with given fields: table
func (_m *QueryBuilder) Table(table string) db.QueryBuilder {
	ret := _m.Called(table)

	var r0 db.QueryBuilder
	if rf, ok := ret.Get(0).(func(string) db.QueryBuilder); ok {
		r0 = rf(table)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.QueryBuilder)
		}
	}

	return r0
}

// Where provides a mock function with given fields: query, args
func (_m *QueryBuilder) Where(query interface{}, args ...interface{}) db.QueryBuilder {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 db.QueryBuilder
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) db.QueryBuilder); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.QueryBuilder)
		}
	}

	return r0
}
