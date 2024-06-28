// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	db_repo "github.com/justtrackio/gosoline/pkg/db-repo"
	mock "github.com/stretchr/testify/mock"
)

// RepositoryReadOnly is an autogenerated mock type for the RepositoryReadOnly type
type RepositoryReadOnly struct {
	mock.Mock
}

type RepositoryReadOnly_Expecter struct {
	mock *mock.Mock
}

func (_m *RepositoryReadOnly) EXPECT() *RepositoryReadOnly_Expecter {
	return &RepositoryReadOnly_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, qb, model
func (_m *RepositoryReadOnly) Count(ctx context.Context, qb *db_repo.QueryBuilder, model db_repo.ModelBased) (int, error) {
	ret := _m.Called(ctx, qb, model)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) (int, error)); ok {
		return rf(ctx, qb, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) int); ok {
		r0 = rf(ctx, qb, model)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) error); ok {
		r1 = rf(ctx, qb, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RepositoryReadOnly_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type RepositoryReadOnly_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - qb *db_repo.QueryBuilder
//   - model db_repo.ModelBased
func (_e *RepositoryReadOnly_Expecter) Count(ctx interface{}, qb interface{}, model interface{}) *RepositoryReadOnly_Count_Call {
	return &RepositoryReadOnly_Count_Call{Call: _e.mock.On("Count", ctx, qb, model)}
}

func (_c *RepositoryReadOnly_Count_Call) Run(run func(ctx context.Context, qb *db_repo.QueryBuilder, model db_repo.ModelBased)) *RepositoryReadOnly_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*db_repo.QueryBuilder), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *RepositoryReadOnly_Count_Call) Return(_a0 int, _a1 error) *RepositoryReadOnly_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *RepositoryReadOnly_Count_Call) RunAndReturn(run func(context.Context, *db_repo.QueryBuilder, db_repo.ModelBased) (int, error)) *RepositoryReadOnly_Count_Call {
	_c.Call.Return(run)
	return _c
}

// GetMetadata provides a mock function with given fields:
func (_m *RepositoryReadOnly) GetMetadata() db_repo.Metadata {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 db_repo.Metadata
	if rf, ok := ret.Get(0).(func() db_repo.Metadata); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(db_repo.Metadata)
	}

	return r0
}

// RepositoryReadOnly_GetMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMetadata'
type RepositoryReadOnly_GetMetadata_Call struct {
	*mock.Call
}

// GetMetadata is a helper method to define mock.On call
func (_e *RepositoryReadOnly_Expecter) GetMetadata() *RepositoryReadOnly_GetMetadata_Call {
	return &RepositoryReadOnly_GetMetadata_Call{Call: _e.mock.On("GetMetadata")}
}

func (_c *RepositoryReadOnly_GetMetadata_Call) Run(run func()) *RepositoryReadOnly_GetMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RepositoryReadOnly_GetMetadata_Call) Return(_a0 db_repo.Metadata) *RepositoryReadOnly_GetMetadata_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryReadOnly_GetMetadata_Call) RunAndReturn(run func() db_repo.Metadata) *RepositoryReadOnly_GetMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// GetModelId provides a mock function with given fields:
func (_m *RepositoryReadOnly) GetModelId() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetModelId")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RepositoryReadOnly_GetModelId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModelId'
type RepositoryReadOnly_GetModelId_Call struct {
	*mock.Call
}

// GetModelId is a helper method to define mock.On call
func (_e *RepositoryReadOnly_Expecter) GetModelId() *RepositoryReadOnly_GetModelId_Call {
	return &RepositoryReadOnly_GetModelId_Call{Call: _e.mock.On("GetModelId")}
}

func (_c *RepositoryReadOnly_GetModelId_Call) Run(run func()) *RepositoryReadOnly_GetModelId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RepositoryReadOnly_GetModelId_Call) Return(_a0 string) *RepositoryReadOnly_GetModelId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryReadOnly_GetModelId_Call) RunAndReturn(run func() string) *RepositoryReadOnly_GetModelId_Call {
	_c.Call.Return(run)
	return _c
}

// GetModelName provides a mock function with given fields:
func (_m *RepositoryReadOnly) GetModelName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetModelName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RepositoryReadOnly_GetModelName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModelName'
type RepositoryReadOnly_GetModelName_Call struct {
	*mock.Call
}

// GetModelName is a helper method to define mock.On call
func (_e *RepositoryReadOnly_Expecter) GetModelName() *RepositoryReadOnly_GetModelName_Call {
	return &RepositoryReadOnly_GetModelName_Call{Call: _e.mock.On("GetModelName")}
}

func (_c *RepositoryReadOnly_GetModelName_Call) Run(run func()) *RepositoryReadOnly_GetModelName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *RepositoryReadOnly_GetModelName_Call) Return(_a0 string) *RepositoryReadOnly_GetModelName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryReadOnly_GetModelName_Call) RunAndReturn(run func() string) *RepositoryReadOnly_GetModelName_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, qb, result
func (_m *RepositoryReadOnly) Query(ctx context.Context, qb *db_repo.QueryBuilder, result interface{}) error {
	ret := _m.Called(ctx, qb, result)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *db_repo.QueryBuilder, interface{}) error); ok {
		r0 = rf(ctx, qb, result)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryReadOnly_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type RepositoryReadOnly_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - qb *db_repo.QueryBuilder
//   - result interface{}
func (_e *RepositoryReadOnly_Expecter) Query(ctx interface{}, qb interface{}, result interface{}) *RepositoryReadOnly_Query_Call {
	return &RepositoryReadOnly_Query_Call{Call: _e.mock.On("Query", ctx, qb, result)}
}

func (_c *RepositoryReadOnly_Query_Call) Run(run func(ctx context.Context, qb *db_repo.QueryBuilder, result interface{})) *RepositoryReadOnly_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*db_repo.QueryBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *RepositoryReadOnly_Query_Call) Return(_a0 error) *RepositoryReadOnly_Query_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryReadOnly_Query_Call) RunAndReturn(run func(context.Context, *db_repo.QueryBuilder, interface{}) error) *RepositoryReadOnly_Query_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: ctx, id, out
func (_m *RepositoryReadOnly) Read(ctx context.Context, id *uint, out db_repo.ModelBased) error {
	ret := _m.Called(ctx, id, out)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *uint, db_repo.ModelBased) error); ok {
		r0 = rf(ctx, id, out)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RepositoryReadOnly_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type RepositoryReadOnly_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - ctx context.Context
//   - id *uint
//   - out db_repo.ModelBased
func (_e *RepositoryReadOnly_Expecter) Read(ctx interface{}, id interface{}, out interface{}) *RepositoryReadOnly_Read_Call {
	return &RepositoryReadOnly_Read_Call{Call: _e.mock.On("Read", ctx, id, out)}
}

func (_c *RepositoryReadOnly_Read_Call) Run(run func(ctx context.Context, id *uint, out db_repo.ModelBased)) *RepositoryReadOnly_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*uint), args[2].(db_repo.ModelBased))
	})
	return _c
}

func (_c *RepositoryReadOnly_Read_Call) Return(_a0 error) *RepositoryReadOnly_Read_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *RepositoryReadOnly_Read_Call) RunAndReturn(run func(context.Context, *uint, db_repo.ModelBased) error) *RepositoryReadOnly_Read_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepositoryReadOnly creates a new instance of RepositoryReadOnly. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepositoryReadOnly(t interface {
	mock.TestingT
	Cleanup(func())
}) *RepositoryReadOnly {
	mock := &RepositoryReadOnly{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
