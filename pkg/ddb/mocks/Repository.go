// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import ddb "github.com/applike/gosoline/pkg/ddb"
import djoemo "github.com/adjoeio/djoemo"
import mdl "github.com/applike/gosoline/pkg/mdl"
import mock "github.com/stretchr/testify/mock"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateTable provides a mock function with given fields: model
func (_m *Repository) CreateTable(model interface{}) error {
	ret := _m.Called(model)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(model)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetItem provides a mock function with given fields: ctx, qb, result
func (_m *Repository) GetItem(ctx context.Context, qb ddb.QueryBuilder, result interface{}) (bool, error) {
	ret := _m.Called(ctx, qb, result)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, ddb.QueryBuilder, interface{}) bool); ok {
		r0 = rf(ctx, qb, result)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ddb.QueryBuilder, interface{}) error); ok {
		r1 = rf(ctx, qb, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItems provides a mock function with given fields: ctx, qb, result
func (_m *Repository) GetItems(ctx context.Context, qb ddb.QueryBuilder, result interface{}) (bool, error) {
	ret := _m.Called(ctx, qb, result)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, ddb.QueryBuilder, interface{}) bool); ok {
		r0 = rf(ctx, qb, result)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ddb.QueryBuilder, interface{}) error); ok {
		r1 = rf(ctx, qb, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModelId provides a mock function with given fields:
func (_m *Repository) GetModelId() mdl.ModelId {
	ret := _m.Called()

	var r0 mdl.ModelId
	if rf, ok := ret.Get(0).(func() mdl.ModelId); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mdl.ModelId)
	}

	return r0
}

// Query provides a mock function with given fields: ctx, qb, result
func (_m *Repository) Query(ctx context.Context, qb ddb.QueryBuilder, result interface{}) error {
	ret := _m.Called(ctx, qb, result)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.QueryBuilder, interface{}) error); ok {
		r0 = rf(ctx, qb, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryBuilder provides a mock function with given fields:
func (_m *Repository) QueryBuilder() ddb.QueryBuilder {
	ret := _m.Called()

	var r0 ddb.QueryBuilder
	if rf, ok := ret.Get(0).(func() ddb.QueryBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.QueryBuilder)
		}
	}

	return r0
}

// Save provides a mock function with given fields: ctx, item
func (_m *Repository) Save(ctx context.Context, item interface{}) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: ctx, exp, qb, values
func (_m *Repository) Update(ctx context.Context, exp djoemo.UpdateExpression, qb ddb.QueryBuilder, values map[string]interface{}) error {
	ret := _m.Called(ctx, exp, qb, values)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, djoemo.UpdateExpression, ddb.QueryBuilder, map[string]interface{}) error); ok {
		r0 = rf(ctx, exp, qb, values)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
