// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	ddb "github.com/justtrackio/gosoline/pkg/ddb"
	mdl "github.com/justtrackio/gosoline/pkg/mdl"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// BatchDeleteItems provides a mock function with given fields: ctx, value
func (_m *Repository) BatchDeleteItems(ctx context.Context, value interface{}) (*ddb.OperationResult, error) {
	ret := _m.Called(ctx, value)

	if len(ret) == 0 {
		panic("no return value specified for BatchDeleteItems")
	}

	var r0 *ddb.OperationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*ddb.OperationResult, error)); ok {
		return rf(ctx, value)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *ddb.OperationResult); ok {
		r0 = rf(ctx, value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.OperationResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_BatchDeleteItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchDeleteItems'
type Repository_BatchDeleteItems_Call struct {
	*mock.Call
}

// BatchDeleteItems is a helper method to define mock.On call
//   - ctx context.Context
//   - value interface{}
func (_e *Repository_Expecter) BatchDeleteItems(ctx interface{}, value interface{}) *Repository_BatchDeleteItems_Call {
	return &Repository_BatchDeleteItems_Call{Call: _e.mock.On("BatchDeleteItems", ctx, value)}
}

func (_c *Repository_BatchDeleteItems_Call) Run(run func(ctx context.Context, value interface{})) *Repository_BatchDeleteItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *Repository_BatchDeleteItems_Call) Return(_a0 *ddb.OperationResult, _a1 error) *Repository_BatchDeleteItems_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_BatchDeleteItems_Call) RunAndReturn(run func(context.Context, interface{}) (*ddb.OperationResult, error)) *Repository_BatchDeleteItems_Call {
	_c.Call.Return(run)
	return _c
}

// BatchGetItems provides a mock function with given fields: ctx, qb, result
func (_m *Repository) BatchGetItems(ctx context.Context, qb ddb.BatchGetItemsBuilder, result interface{}) (*ddb.OperationResult, error) {
	ret := _m.Called(ctx, qb, result)

	if len(ret) == 0 {
		panic("no return value specified for BatchGetItems")
	}

	var r0 *ddb.OperationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.BatchGetItemsBuilder, interface{}) (*ddb.OperationResult, error)); ok {
		return rf(ctx, qb, result)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.BatchGetItemsBuilder, interface{}) *ddb.OperationResult); ok {
		r0 = rf(ctx, qb, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.OperationResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.BatchGetItemsBuilder, interface{}) error); ok {
		r1 = rf(ctx, qb, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_BatchGetItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGetItems'
type Repository_BatchGetItems_Call struct {
	*mock.Call
}

// BatchGetItems is a helper method to define mock.On call
//   - ctx context.Context
//   - qb ddb.BatchGetItemsBuilder
//   - result interface{}
func (_e *Repository_Expecter) BatchGetItems(ctx interface{}, qb interface{}, result interface{}) *Repository_BatchGetItems_Call {
	return &Repository_BatchGetItems_Call{Call: _e.mock.On("BatchGetItems", ctx, qb, result)}
}

func (_c *Repository_BatchGetItems_Call) Run(run func(ctx context.Context, qb ddb.BatchGetItemsBuilder, result interface{})) *Repository_BatchGetItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.BatchGetItemsBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_BatchGetItems_Call) Return(_a0 *ddb.OperationResult, _a1 error) *Repository_BatchGetItems_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_BatchGetItems_Call) RunAndReturn(run func(context.Context, ddb.BatchGetItemsBuilder, interface{}) (*ddb.OperationResult, error)) *Repository_BatchGetItems_Call {
	_c.Call.Return(run)
	return _c
}

// BatchGetItemsBuilder provides a mock function with given fields:
func (_m *Repository) BatchGetItemsBuilder() ddb.BatchGetItemsBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BatchGetItemsBuilder")
	}

	var r0 ddb.BatchGetItemsBuilder
	if rf, ok := ret.Get(0).(func() ddb.BatchGetItemsBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.BatchGetItemsBuilder)
		}
	}

	return r0
}

// Repository_BatchGetItemsBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchGetItemsBuilder'
type Repository_BatchGetItemsBuilder_Call struct {
	*mock.Call
}

// BatchGetItemsBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) BatchGetItemsBuilder() *Repository_BatchGetItemsBuilder_Call {
	return &Repository_BatchGetItemsBuilder_Call{Call: _e.mock.On("BatchGetItemsBuilder")}
}

func (_c *Repository_BatchGetItemsBuilder_Call) Run(run func()) *Repository_BatchGetItemsBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_BatchGetItemsBuilder_Call) Return(_a0 ddb.BatchGetItemsBuilder) *Repository_BatchGetItemsBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_BatchGetItemsBuilder_Call) RunAndReturn(run func() ddb.BatchGetItemsBuilder) *Repository_BatchGetItemsBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// BatchPutItems provides a mock function with given fields: ctx, items
func (_m *Repository) BatchPutItems(ctx context.Context, items interface{}) (*ddb.OperationResult, error) {
	ret := _m.Called(ctx, items)

	if len(ret) == 0 {
		panic("no return value specified for BatchPutItems")
	}

	var r0 *ddb.OperationResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*ddb.OperationResult, error)); ok {
		return rf(ctx, items)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *ddb.OperationResult); ok {
		r0 = rf(ctx, items)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.OperationResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, items)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_BatchPutItems_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BatchPutItems'
type Repository_BatchPutItems_Call struct {
	*mock.Call
}

// BatchPutItems is a helper method to define mock.On call
//   - ctx context.Context
//   - items interface{}
func (_e *Repository_Expecter) BatchPutItems(ctx interface{}, items interface{}) *Repository_BatchPutItems_Call {
	return &Repository_BatchPutItems_Call{Call: _e.mock.On("BatchPutItems", ctx, items)}
}

func (_c *Repository_BatchPutItems_Call) Run(run func(ctx context.Context, items interface{})) *Repository_BatchPutItems_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(interface{}))
	})
	return _c
}

func (_c *Repository_BatchPutItems_Call) Return(_a0 *ddb.OperationResult, _a1 error) *Repository_BatchPutItems_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_BatchPutItems_Call) RunAndReturn(run func(context.Context, interface{}) (*ddb.OperationResult, error)) *Repository_BatchPutItems_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteItem provides a mock function with given fields: ctx, db, item
func (_m *Repository) DeleteItem(ctx context.Context, db ddb.DeleteItemBuilder, item interface{}) (*ddb.DeleteItemResult, error) {
	ret := _m.Called(ctx, db, item)

	if len(ret) == 0 {
		panic("no return value specified for DeleteItem")
	}

	var r0 *ddb.DeleteItemResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.DeleteItemBuilder, interface{}) (*ddb.DeleteItemResult, error)); ok {
		return rf(ctx, db, item)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.DeleteItemBuilder, interface{}) *ddb.DeleteItemResult); ok {
		r0 = rf(ctx, db, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.DeleteItemResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.DeleteItemBuilder, interface{}) error); ok {
		r1 = rf(ctx, db, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_DeleteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItem'
type Repository_DeleteItem_Call struct {
	*mock.Call
}

// DeleteItem is a helper method to define mock.On call
//   - ctx context.Context
//   - db ddb.DeleteItemBuilder
//   - item interface{}
func (_e *Repository_Expecter) DeleteItem(ctx interface{}, db interface{}, item interface{}) *Repository_DeleteItem_Call {
	return &Repository_DeleteItem_Call{Call: _e.mock.On("DeleteItem", ctx, db, item)}
}

func (_c *Repository_DeleteItem_Call) Run(run func(ctx context.Context, db ddb.DeleteItemBuilder, item interface{})) *Repository_DeleteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.DeleteItemBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_DeleteItem_Call) Return(_a0 *ddb.DeleteItemResult, _a1 error) *Repository_DeleteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_DeleteItem_Call) RunAndReturn(run func(context.Context, ddb.DeleteItemBuilder, interface{}) (*ddb.DeleteItemResult, error)) *Repository_DeleteItem_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteItemBuilder provides a mock function with given fields:
func (_m *Repository) DeleteItemBuilder() ddb.DeleteItemBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeleteItemBuilder")
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

// Repository_DeleteItemBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItemBuilder'
type Repository_DeleteItemBuilder_Call struct {
	*mock.Call
}

// DeleteItemBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) DeleteItemBuilder() *Repository_DeleteItemBuilder_Call {
	return &Repository_DeleteItemBuilder_Call{Call: _e.mock.On("DeleteItemBuilder")}
}

func (_c *Repository_DeleteItemBuilder_Call) Run(run func()) *Repository_DeleteItemBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_DeleteItemBuilder_Call) Return(_a0 ddb.DeleteItemBuilder) *Repository_DeleteItemBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_DeleteItemBuilder_Call) RunAndReturn(run func() ddb.DeleteItemBuilder) *Repository_DeleteItemBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// GetItem provides a mock function with given fields: ctx, qb, result
func (_m *Repository) GetItem(ctx context.Context, qb ddb.GetItemBuilder, result interface{}) (*ddb.GetItemResult, error) {
	ret := _m.Called(ctx, qb, result)

	if len(ret) == 0 {
		panic("no return value specified for GetItem")
	}

	var r0 *ddb.GetItemResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.GetItemBuilder, interface{}) (*ddb.GetItemResult, error)); ok {
		return rf(ctx, qb, result)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.GetItemBuilder, interface{}) *ddb.GetItemResult); ok {
		r0 = rf(ctx, qb, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.GetItemResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.GetItemBuilder, interface{}) error); ok {
		r1 = rf(ctx, qb, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItem'
type Repository_GetItem_Call struct {
	*mock.Call
}

// GetItem is a helper method to define mock.On call
//   - ctx context.Context
//   - qb ddb.GetItemBuilder
//   - result interface{}
func (_e *Repository_Expecter) GetItem(ctx interface{}, qb interface{}, result interface{}) *Repository_GetItem_Call {
	return &Repository_GetItem_Call{Call: _e.mock.On("GetItem", ctx, qb, result)}
}

func (_c *Repository_GetItem_Call) Run(run func(ctx context.Context, qb ddb.GetItemBuilder, result interface{})) *Repository_GetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.GetItemBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_GetItem_Call) Return(_a0 *ddb.GetItemResult, _a1 error) *Repository_GetItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetItem_Call) RunAndReturn(run func(context.Context, ddb.GetItemBuilder, interface{}) (*ddb.GetItemResult, error)) *Repository_GetItem_Call {
	_c.Call.Return(run)
	return _c
}

// GetItemBuilder provides a mock function with given fields:
func (_m *Repository) GetItemBuilder() ddb.GetItemBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetItemBuilder")
	}

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

// Repository_GetItemBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItemBuilder'
type Repository_GetItemBuilder_Call struct {
	*mock.Call
}

// GetItemBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) GetItemBuilder() *Repository_GetItemBuilder_Call {
	return &Repository_GetItemBuilder_Call{Call: _e.mock.On("GetItemBuilder")}
}

func (_c *Repository_GetItemBuilder_Call) Run(run func()) *Repository_GetItemBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_GetItemBuilder_Call) Return(_a0 ddb.GetItemBuilder) *Repository_GetItemBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_GetItemBuilder_Call) RunAndReturn(run func() ddb.GetItemBuilder) *Repository_GetItemBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// GetModelId provides a mock function with given fields:
func (_m *Repository) GetModelId() mdl.ModelId {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetModelId")
	}

	var r0 mdl.ModelId
	if rf, ok := ret.Get(0).(func() mdl.ModelId); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(mdl.ModelId)
	}

	return r0
}

// Repository_GetModelId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetModelId'
type Repository_GetModelId_Call struct {
	*mock.Call
}

// GetModelId is a helper method to define mock.On call
func (_e *Repository_Expecter) GetModelId() *Repository_GetModelId_Call {
	return &Repository_GetModelId_Call{Call: _e.mock.On("GetModelId")}
}

func (_c *Repository_GetModelId_Call) Run(run func()) *Repository_GetModelId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_GetModelId_Call) Return(_a0 mdl.ModelId) *Repository_GetModelId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_GetModelId_Call) RunAndReturn(run func() mdl.ModelId) *Repository_GetModelId_Call {
	_c.Call.Return(run)
	return _c
}

// PutItem provides a mock function with given fields: ctx, qb, item
func (_m *Repository) PutItem(ctx context.Context, qb ddb.PutItemBuilder, item interface{}) (*ddb.PutItemResult, error) {
	ret := _m.Called(ctx, qb, item)

	if len(ret) == 0 {
		panic("no return value specified for PutItem")
	}

	var r0 *ddb.PutItemResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.PutItemBuilder, interface{}) (*ddb.PutItemResult, error)); ok {
		return rf(ctx, qb, item)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.PutItemBuilder, interface{}) *ddb.PutItemResult); ok {
		r0 = rf(ctx, qb, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.PutItemResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.PutItemBuilder, interface{}) error); ok {
		r1 = rf(ctx, qb, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_PutItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItem'
type Repository_PutItem_Call struct {
	*mock.Call
}

// PutItem is a helper method to define mock.On call
//   - ctx context.Context
//   - qb ddb.PutItemBuilder
//   - item interface{}
func (_e *Repository_Expecter) PutItem(ctx interface{}, qb interface{}, item interface{}) *Repository_PutItem_Call {
	return &Repository_PutItem_Call{Call: _e.mock.On("PutItem", ctx, qb, item)}
}

func (_c *Repository_PutItem_Call) Run(run func(ctx context.Context, qb ddb.PutItemBuilder, item interface{})) *Repository_PutItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.PutItemBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_PutItem_Call) Return(_a0 *ddb.PutItemResult, _a1 error) *Repository_PutItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_PutItem_Call) RunAndReturn(run func(context.Context, ddb.PutItemBuilder, interface{}) (*ddb.PutItemResult, error)) *Repository_PutItem_Call {
	_c.Call.Return(run)
	return _c
}

// PutItemBuilder provides a mock function with given fields:
func (_m *Repository) PutItemBuilder() ddb.PutItemBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PutItemBuilder")
	}

	var r0 ddb.PutItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.PutItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.PutItemBuilder)
		}
	}

	return r0
}

// Repository_PutItemBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItemBuilder'
type Repository_PutItemBuilder_Call struct {
	*mock.Call
}

// PutItemBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) PutItemBuilder() *Repository_PutItemBuilder_Call {
	return &Repository_PutItemBuilder_Call{Call: _e.mock.On("PutItemBuilder")}
}

func (_c *Repository_PutItemBuilder_Call) Run(run func()) *Repository_PutItemBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_PutItemBuilder_Call) Return(_a0 ddb.PutItemBuilder) *Repository_PutItemBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_PutItemBuilder_Call) RunAndReturn(run func() ddb.PutItemBuilder) *Repository_PutItemBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, qb, result
func (_m *Repository) Query(ctx context.Context, qb ddb.QueryBuilder, result interface{}) (*ddb.QueryResult, error) {
	ret := _m.Called(ctx, qb, result)

	if len(ret) == 0 {
		panic("no return value specified for Query")
	}

	var r0 *ddb.QueryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.QueryBuilder, interface{}) (*ddb.QueryResult, error)); ok {
		return rf(ctx, qb, result)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.QueryBuilder, interface{}) *ddb.QueryResult); ok {
		r0 = rf(ctx, qb, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.QueryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.QueryBuilder, interface{}) error); ok {
		r1 = rf(ctx, qb, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type Repository_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - qb ddb.QueryBuilder
//   - result interface{}
func (_e *Repository_Expecter) Query(ctx interface{}, qb interface{}, result interface{}) *Repository_Query_Call {
	return &Repository_Query_Call{Call: _e.mock.On("Query", ctx, qb, result)}
}

func (_c *Repository_Query_Call) Run(run func(ctx context.Context, qb ddb.QueryBuilder, result interface{})) *Repository_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.QueryBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_Query_Call) Return(_a0 *ddb.QueryResult, _a1 error) *Repository_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Query_Call) RunAndReturn(run func(context.Context, ddb.QueryBuilder, interface{}) (*ddb.QueryResult, error)) *Repository_Query_Call {
	_c.Call.Return(run)
	return _c
}

// QueryBuilder provides a mock function with given fields:
func (_m *Repository) QueryBuilder() ddb.QueryBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for QueryBuilder")
	}

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

// Repository_QueryBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryBuilder'
type Repository_QueryBuilder_Call struct {
	*mock.Call
}

// QueryBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) QueryBuilder() *Repository_QueryBuilder_Call {
	return &Repository_QueryBuilder_Call{Call: _e.mock.On("QueryBuilder")}
}

func (_c *Repository_QueryBuilder_Call) Run(run func()) *Repository_QueryBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_QueryBuilder_Call) Return(_a0 ddb.QueryBuilder) *Repository_QueryBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_QueryBuilder_Call) RunAndReturn(run func() ddb.QueryBuilder) *Repository_QueryBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// Scan provides a mock function with given fields: ctx, sb, result
func (_m *Repository) Scan(ctx context.Context, sb ddb.ScanBuilder, result interface{}) (*ddb.ScanResult, error) {
	ret := _m.Called(ctx, sb, result)

	if len(ret) == 0 {
		panic("no return value specified for Scan")
	}

	var r0 *ddb.ScanResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.ScanBuilder, interface{}) (*ddb.ScanResult, error)); ok {
		return rf(ctx, sb, result)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.ScanBuilder, interface{}) *ddb.ScanResult); ok {
		r0 = rf(ctx, sb, result)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.ScanResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.ScanBuilder, interface{}) error); ok {
		r1 = rf(ctx, sb, result)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_Scan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Scan'
type Repository_Scan_Call struct {
	*mock.Call
}

// Scan is a helper method to define mock.On call
//   - ctx context.Context
//   - sb ddb.ScanBuilder
//   - result interface{}
func (_e *Repository_Expecter) Scan(ctx interface{}, sb interface{}, result interface{}) *Repository_Scan_Call {
	return &Repository_Scan_Call{Call: _e.mock.On("Scan", ctx, sb, result)}
}

func (_c *Repository_Scan_Call) Run(run func(ctx context.Context, sb ddb.ScanBuilder, result interface{})) *Repository_Scan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.ScanBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_Scan_Call) Return(_a0 *ddb.ScanResult, _a1 error) *Repository_Scan_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_Scan_Call) RunAndReturn(run func(context.Context, ddb.ScanBuilder, interface{}) (*ddb.ScanResult, error)) *Repository_Scan_Call {
	_c.Call.Return(run)
	return _c
}

// ScanBuilder provides a mock function with given fields:
func (_m *Repository) ScanBuilder() ddb.ScanBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ScanBuilder")
	}

	var r0 ddb.ScanBuilder
	if rf, ok := ret.Get(0).(func() ddb.ScanBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.ScanBuilder)
		}
	}

	return r0
}

// Repository_ScanBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScanBuilder'
type Repository_ScanBuilder_Call struct {
	*mock.Call
}

// ScanBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) ScanBuilder() *Repository_ScanBuilder_Call {
	return &Repository_ScanBuilder_Call{Call: _e.mock.On("ScanBuilder")}
}

func (_c *Repository_ScanBuilder_Call) Run(run func()) *Repository_ScanBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_ScanBuilder_Call) Return(_a0 ddb.ScanBuilder) *Repository_ScanBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_ScanBuilder_Call) RunAndReturn(run func() ddb.ScanBuilder) *Repository_ScanBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItem provides a mock function with given fields: ctx, ub, item
func (_m *Repository) UpdateItem(ctx context.Context, ub ddb.UpdateItemBuilder, item interface{}) (*ddb.UpdateItemResult, error) {
	ret := _m.Called(ctx, ub, item)

	if len(ret) == 0 {
		panic("no return value specified for UpdateItem")
	}

	var r0 *ddb.UpdateItemResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ddb.UpdateItemBuilder, interface{}) (*ddb.UpdateItemResult, error)); ok {
		return rf(ctx, ub, item)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ddb.UpdateItemBuilder, interface{}) *ddb.UpdateItemResult); ok {
		r0 = rf(ctx, ub, item)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ddb.UpdateItemResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ddb.UpdateItemBuilder, interface{}) error); ok {
		r1 = rf(ctx, ub, item)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_UpdateItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItem'
type Repository_UpdateItem_Call struct {
	*mock.Call
}

// UpdateItem is a helper method to define mock.On call
//   - ctx context.Context
//   - ub ddb.UpdateItemBuilder
//   - item interface{}
func (_e *Repository_Expecter) UpdateItem(ctx interface{}, ub interface{}, item interface{}) *Repository_UpdateItem_Call {
	return &Repository_UpdateItem_Call{Call: _e.mock.On("UpdateItem", ctx, ub, item)}
}

func (_c *Repository_UpdateItem_Call) Run(run func(ctx context.Context, ub ddb.UpdateItemBuilder, item interface{})) *Repository_UpdateItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ddb.UpdateItemBuilder), args[2].(interface{}))
	})
	return _c
}

func (_c *Repository_UpdateItem_Call) Return(_a0 *ddb.UpdateItemResult, _a1 error) *Repository_UpdateItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_UpdateItem_Call) RunAndReturn(run func(context.Context, ddb.UpdateItemBuilder, interface{}) (*ddb.UpdateItemResult, error)) *Repository_UpdateItem_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItemBuilder provides a mock function with given fields:
func (_m *Repository) UpdateItemBuilder() ddb.UpdateItemBuilder {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UpdateItemBuilder")
	}

	var r0 ddb.UpdateItemBuilder
	if rf, ok := ret.Get(0).(func() ddb.UpdateItemBuilder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ddb.UpdateItemBuilder)
		}
	}

	return r0
}

// Repository_UpdateItemBuilder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItemBuilder'
type Repository_UpdateItemBuilder_Call struct {
	*mock.Call
}

// UpdateItemBuilder is a helper method to define mock.On call
func (_e *Repository_Expecter) UpdateItemBuilder() *Repository_UpdateItemBuilder_Call {
	return &Repository_UpdateItemBuilder_Call{Call: _e.mock.On("UpdateItemBuilder")}
}

func (_c *Repository_UpdateItemBuilder_Call) Run(run func()) *Repository_UpdateItemBuilder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Repository_UpdateItemBuilder_Call) Return(_a0 ddb.UpdateItemBuilder) *Repository_UpdateItemBuilder_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_UpdateItemBuilder_Call) RunAndReturn(run func() ddb.UpdateItemBuilder) *Repository_UpdateItemBuilder_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
