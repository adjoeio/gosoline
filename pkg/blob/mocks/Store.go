// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	blob "github.com/justtrackio/gosoline/pkg/blob"

	mock "github.com/stretchr/testify/mock"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// BucketName provides a mock function with given fields:
func (_m *Store) BucketName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BucketName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Store_BucketName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BucketName'
type Store_BucketName_Call struct {
	*mock.Call
}

// BucketName is a helper method to define mock.On call
func (_e *Store_Expecter) BucketName() *Store_BucketName_Call {
	return &Store_BucketName_Call{Call: _e.mock.On("BucketName")}
}

func (_c *Store_BucketName_Call) Run(run func()) *Store_BucketName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Store_BucketName_Call) Return(_a0 string) *Store_BucketName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_BucketName_Call) RunAndReturn(run func() string) *Store_BucketName_Call {
	_c.Call.Return(run)
	return _c
}

// Copy provides a mock function with given fields: batch
func (_m *Store) Copy(batch blob.CopyBatch) {
	_m.Called(batch)
}

// Store_Copy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Copy'
type Store_Copy_Call struct {
	*mock.Call
}

// Copy is a helper method to define mock.On call
//   - batch blob.CopyBatch
func (_e *Store_Expecter) Copy(batch interface{}) *Store_Copy_Call {
	return &Store_Copy_Call{Call: _e.mock.On("Copy", batch)}
}

func (_c *Store_Copy_Call) Run(run func(batch blob.CopyBatch)) *Store_Copy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(blob.CopyBatch))
	})
	return _c
}

func (_c *Store_Copy_Call) Return() *Store_Copy_Call {
	_c.Call.Return()
	return _c
}

func (_c *Store_Copy_Call) RunAndReturn(run func(blob.CopyBatch)) *Store_Copy_Call {
	_c.Call.Return(run)
	return _c
}

// CopyOne provides a mock function with given fields: obj
func (_m *Store) CopyOne(obj *blob.CopyObject) error {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for CopyOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.CopyObject) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_CopyOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyOne'
type Store_CopyOne_Call struct {
	*mock.Call
}

// CopyOne is a helper method to define mock.On call
//   - obj *blob.CopyObject
func (_e *Store_Expecter) CopyOne(obj interface{}) *Store_CopyOne_Call {
	return &Store_CopyOne_Call{Call: _e.mock.On("CopyOne", obj)}
}

func (_c *Store_CopyOne_Call) Run(run func(obj *blob.CopyObject)) *Store_CopyOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*blob.CopyObject))
	})
	return _c
}

func (_c *Store_CopyOne_Call) Return(_a0 error) *Store_CopyOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_CopyOne_Call) RunAndReturn(run func(*blob.CopyObject) error) *Store_CopyOne_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBucket provides a mock function with given fields: ctx
func (_m *Store) CreateBucket(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CreateBucket")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_CreateBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBucket'
type Store_CreateBucket_Call struct {
	*mock.Call
}

// CreateBucket is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Store_Expecter) CreateBucket(ctx interface{}) *Store_CreateBucket_Call {
	return &Store_CreateBucket_Call{Call: _e.mock.On("CreateBucket", ctx)}
}

func (_c *Store_CreateBucket_Call) Run(run func(ctx context.Context)) *Store_CreateBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Store_CreateBucket_Call) Return(_a0 error) *Store_CreateBucket_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_CreateBucket_Call) RunAndReturn(run func(context.Context) error) *Store_CreateBucket_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: batch
func (_m *Store) Delete(batch blob.Batch) {
	_m.Called(batch)
}

// Store_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Store_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - batch blob.Batch
func (_e *Store_Expecter) Delete(batch interface{}) *Store_Delete_Call {
	return &Store_Delete_Call{Call: _e.mock.On("Delete", batch)}
}

func (_c *Store_Delete_Call) Run(run func(batch blob.Batch)) *Store_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(blob.Batch))
	})
	return _c
}

func (_c *Store_Delete_Call) Return() *Store_Delete_Call {
	_c.Call.Return()
	return _c
}

func (_c *Store_Delete_Call) RunAndReturn(run func(blob.Batch)) *Store_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBucket provides a mock function with given fields: ctx
func (_m *Store) DeleteBucket(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBucket")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_DeleteBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBucket'
type Store_DeleteBucket_Call struct {
	*mock.Call
}

// DeleteBucket is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Store_Expecter) DeleteBucket(ctx interface{}) *Store_DeleteBucket_Call {
	return &Store_DeleteBucket_Call{Call: _e.mock.On("DeleteBucket", ctx)}
}

func (_c *Store_DeleteBucket_Call) Run(run func(ctx context.Context)) *Store_DeleteBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Store_DeleteBucket_Call) Return(_a0 error) *Store_DeleteBucket_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_DeleteBucket_Call) RunAndReturn(run func(context.Context) error) *Store_DeleteBucket_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteOne provides a mock function with given fields: obj
func (_m *Store) DeleteOne(obj *blob.Object) error {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.Object) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_DeleteOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteOne'
type Store_DeleteOne_Call struct {
	*mock.Call
}

// DeleteOne is a helper method to define mock.On call
//   - obj *blob.Object
func (_e *Store_Expecter) DeleteOne(obj interface{}) *Store_DeleteOne_Call {
	return &Store_DeleteOne_Call{Call: _e.mock.On("DeleteOne", obj)}
}

func (_c *Store_DeleteOne_Call) Run(run func(obj *blob.Object)) *Store_DeleteOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*blob.Object))
	})
	return _c
}

func (_c *Store_DeleteOne_Call) Return(_a0 error) *Store_DeleteOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_DeleteOne_Call) RunAndReturn(run func(*blob.Object) error) *Store_DeleteOne_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: batch
func (_m *Store) Read(batch blob.Batch) {
	_m.Called(batch)
}

// Store_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type Store_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - batch blob.Batch
func (_e *Store_Expecter) Read(batch interface{}) *Store_Read_Call {
	return &Store_Read_Call{Call: _e.mock.On("Read", batch)}
}

func (_c *Store_Read_Call) Run(run func(batch blob.Batch)) *Store_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(blob.Batch))
	})
	return _c
}

func (_c *Store_Read_Call) Return() *Store_Read_Call {
	_c.Call.Return()
	return _c
}

func (_c *Store_Read_Call) RunAndReturn(run func(blob.Batch)) *Store_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadOne provides a mock function with given fields: obj
func (_m *Store) ReadOne(obj *blob.Object) error {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for ReadOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.Object) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_ReadOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadOne'
type Store_ReadOne_Call struct {
	*mock.Call
}

// ReadOne is a helper method to define mock.On call
//   - obj *blob.Object
func (_e *Store_Expecter) ReadOne(obj interface{}) *Store_ReadOne_Call {
	return &Store_ReadOne_Call{Call: _e.mock.On("ReadOne", obj)}
}

func (_c *Store_ReadOne_Call) Run(run func(obj *blob.Object)) *Store_ReadOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*blob.Object))
	})
	return _c
}

func (_c *Store_ReadOne_Call) Return(_a0 error) *Store_ReadOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_ReadOne_Call) RunAndReturn(run func(*blob.Object) error) *Store_ReadOne_Call {
	_c.Call.Return(run)
	return _c
}

// Write provides a mock function with given fields: batch
func (_m *Store) Write(batch blob.Batch) error {
	ret := _m.Called(batch)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(blob.Batch) error); ok {
		r0 = rf(batch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type Store_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - batch blob.Batch
func (_e *Store_Expecter) Write(batch interface{}) *Store_Write_Call {
	return &Store_Write_Call{Call: _e.mock.On("Write", batch)}
}

func (_c *Store_Write_Call) Run(run func(batch blob.Batch)) *Store_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(blob.Batch))
	})
	return _c
}

func (_c *Store_Write_Call) Return(_a0 error) *Store_Write_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_Write_Call) RunAndReturn(run func(blob.Batch) error) *Store_Write_Call {
	_c.Call.Return(run)
	return _c
}

// WriteOne provides a mock function with given fields: obj
func (_m *Store) WriteOne(obj *blob.Object) error {
	ret := _m.Called(obj)

	if len(ret) == 0 {
		panic("no return value specified for WriteOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.Object) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Store_WriteOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteOne'
type Store_WriteOne_Call struct {
	*mock.Call
}

// WriteOne is a helper method to define mock.On call
//   - obj *blob.Object
func (_e *Store_Expecter) WriteOne(obj interface{}) *Store_WriteOne_Call {
	return &Store_WriteOne_Call{Call: _e.mock.On("WriteOne", obj)}
}

func (_c *Store_WriteOne_Call) Run(run func(obj *blob.Object)) *Store_WriteOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*blob.Object))
	})
	return _c
}

func (_c *Store_WriteOne_Call) Return(_a0 error) *Store_WriteOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Store_WriteOne_Call) RunAndReturn(run func(*blob.Object) error) *Store_WriteOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
