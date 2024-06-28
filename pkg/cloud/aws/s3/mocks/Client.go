// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// AbortMultipartUpload provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) AbortMultipartUpload(_a0 context.Context, _a1 *s3.AbortMultipartUploadInput, _a2 ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AbortMultipartUpload")
	}

	var r0 *s3.AbortMultipartUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.AbortMultipartUploadInput, ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.AbortMultipartUploadInput, ...func(*s3.Options)) *s3.AbortMultipartUploadOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.AbortMultipartUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.AbortMultipartUploadInput, ...func(*s3.Options)) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_AbortMultipartUpload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AbortMultipartUpload'
type Client_AbortMultipartUpload_Call struct {
	*mock.Call
}

// AbortMultipartUpload is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *s3.AbortMultipartUploadInput
//   - _a2 ...func(*s3.Options)
func (_e *Client_Expecter) AbortMultipartUpload(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *Client_AbortMultipartUpload_Call {
	return &Client_AbortMultipartUpload_Call{Call: _e.mock.On("AbortMultipartUpload",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *Client_AbortMultipartUpload_Call) Run(run func(_a0 context.Context, _a1 *s3.AbortMultipartUploadInput, _a2 ...func(*s3.Options))) *Client_AbortMultipartUpload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.AbortMultipartUploadInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_AbortMultipartUpload_Call) Return(_a0 *s3.AbortMultipartUploadOutput, _a1 error) *Client_AbortMultipartUpload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_AbortMultipartUpload_Call) RunAndReturn(run func(context.Context, *s3.AbortMultipartUploadInput, ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error)) *Client_AbortMultipartUpload_Call {
	_c.Call.Return(run)
	return _c
}

// CompleteMultipartUpload provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) CompleteMultipartUpload(_a0 context.Context, _a1 *s3.CompleteMultipartUploadInput, _a2 ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CompleteMultipartUpload")
	}

	var r0 *s3.CompleteMultipartUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CompleteMultipartUploadInput, ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CompleteMultipartUploadInput, ...func(*s3.Options)) *s3.CompleteMultipartUploadOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CompleteMultipartUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.CompleteMultipartUploadInput, ...func(*s3.Options)) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_CompleteMultipartUpload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CompleteMultipartUpload'
type Client_CompleteMultipartUpload_Call struct {
	*mock.Call
}

// CompleteMultipartUpload is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *s3.CompleteMultipartUploadInput
//   - _a2 ...func(*s3.Options)
func (_e *Client_Expecter) CompleteMultipartUpload(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *Client_CompleteMultipartUpload_Call {
	return &Client_CompleteMultipartUpload_Call{Call: _e.mock.On("CompleteMultipartUpload",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *Client_CompleteMultipartUpload_Call) Run(run func(_a0 context.Context, _a1 *s3.CompleteMultipartUploadInput, _a2 ...func(*s3.Options))) *Client_CompleteMultipartUpload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.CompleteMultipartUploadInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_CompleteMultipartUpload_Call) Return(_a0 *s3.CompleteMultipartUploadOutput, _a1 error) *Client_CompleteMultipartUpload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_CompleteMultipartUpload_Call) RunAndReturn(run func(context.Context, *s3.CompleteMultipartUploadInput, ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error)) *Client_CompleteMultipartUpload_Call {
	_c.Call.Return(run)
	return _c
}

// CopyObject provides a mock function with given fields: ctx, params, optFns
func (_m *Client) CopyObject(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options)) (*s3.CopyObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CopyObject")
	}

	var r0 *s3.CopyObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CopyObjectInput, ...func(*s3.Options)) (*s3.CopyObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CopyObjectInput, ...func(*s3.Options)) *s3.CopyObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CopyObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.CopyObjectInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_CopyObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CopyObject'
type Client_CopyObject_Call struct {
	*mock.Call
}

// CopyObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.CopyObjectInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) CopyObject(ctx interface{}, params interface{}, optFns ...interface{}) *Client_CopyObject_Call {
	return &Client_CopyObject_Call{Call: _e.mock.On("CopyObject",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_CopyObject_Call) Run(run func(ctx context.Context, params *s3.CopyObjectInput, optFns ...func(*s3.Options))) *Client_CopyObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.CopyObjectInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_CopyObject_Call) Return(_a0 *s3.CopyObjectOutput, _a1 error) *Client_CopyObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_CopyObject_Call) RunAndReturn(run func(context.Context, *s3.CopyObjectInput, ...func(*s3.Options)) (*s3.CopyObjectOutput, error)) *Client_CopyObject_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBucket provides a mock function with given fields: ctx, params, optFns
func (_m *Client) CreateBucket(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options)) (*s3.CreateBucketOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateBucket")
	}

	var r0 *s3.CreateBucketOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CreateBucketInput, ...func(*s3.Options)) (*s3.CreateBucketOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CreateBucketInput, ...func(*s3.Options)) *s3.CreateBucketOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateBucketOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.CreateBucketInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_CreateBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBucket'
type Client_CreateBucket_Call struct {
	*mock.Call
}

// CreateBucket is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.CreateBucketInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) CreateBucket(ctx interface{}, params interface{}, optFns ...interface{}) *Client_CreateBucket_Call {
	return &Client_CreateBucket_Call{Call: _e.mock.On("CreateBucket",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_CreateBucket_Call) Run(run func(ctx context.Context, params *s3.CreateBucketInput, optFns ...func(*s3.Options))) *Client_CreateBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.CreateBucketInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_CreateBucket_Call) Return(_a0 *s3.CreateBucketOutput, _a1 error) *Client_CreateBucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_CreateBucket_Call) RunAndReturn(run func(context.Context, *s3.CreateBucketInput, ...func(*s3.Options)) (*s3.CreateBucketOutput, error)) *Client_CreateBucket_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMultipartUpload provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) CreateMultipartUpload(_a0 context.Context, _a1 *s3.CreateMultipartUploadInput, _a2 ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMultipartUpload")
	}

	var r0 *s3.CreateMultipartUploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CreateMultipartUploadInput, ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.CreateMultipartUploadInput, ...func(*s3.Options)) *s3.CreateMultipartUploadOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.CreateMultipartUploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.CreateMultipartUploadInput, ...func(*s3.Options)) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_CreateMultipartUpload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMultipartUpload'
type Client_CreateMultipartUpload_Call struct {
	*mock.Call
}

// CreateMultipartUpload is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *s3.CreateMultipartUploadInput
//   - _a2 ...func(*s3.Options)
func (_e *Client_Expecter) CreateMultipartUpload(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *Client_CreateMultipartUpload_Call {
	return &Client_CreateMultipartUpload_Call{Call: _e.mock.On("CreateMultipartUpload",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *Client_CreateMultipartUpload_Call) Run(run func(_a0 context.Context, _a1 *s3.CreateMultipartUploadInput, _a2 ...func(*s3.Options))) *Client_CreateMultipartUpload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.CreateMultipartUploadInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_CreateMultipartUpload_Call) Return(_a0 *s3.CreateMultipartUploadOutput, _a1 error) *Client_CreateMultipartUpload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_CreateMultipartUpload_Call) RunAndReturn(run func(context.Context, *s3.CreateMultipartUploadInput, ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error)) *Client_CreateMultipartUpload_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBucket provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DeleteBucket(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.Options)) (*s3.DeleteBucketOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBucket")
	}

	var r0 *s3.DeleteBucketOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.DeleteBucketInput, ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.DeleteBucketInput, ...func(*s3.Options)) *s3.DeleteBucketOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteBucketOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.DeleteBucketInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DeleteBucket_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBucket'
type Client_DeleteBucket_Call struct {
	*mock.Call
}

// DeleteBucket is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.DeleteBucketInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) DeleteBucket(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DeleteBucket_Call {
	return &Client_DeleteBucket_Call{Call: _e.mock.On("DeleteBucket",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DeleteBucket_Call) Run(run func(ctx context.Context, params *s3.DeleteBucketInput, optFns ...func(*s3.Options))) *Client_DeleteBucket_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.DeleteBucketInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DeleteBucket_Call) Return(_a0 *s3.DeleteBucketOutput, _a1 error) *Client_DeleteBucket_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DeleteBucket_Call) RunAndReturn(run func(context.Context, *s3.DeleteBucketInput, ...func(*s3.Options)) (*s3.DeleteBucketOutput, error)) *Client_DeleteBucket_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteObject provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteObject")
	}

	var r0 *s3.DeleteObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.DeleteObjectInput, ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.DeleteObjectInput, ...func(*s3.Options)) *s3.DeleteObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.DeleteObjectInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DeleteObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteObject'
type Client_DeleteObject_Call struct {
	*mock.Call
}

// DeleteObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.DeleteObjectInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) DeleteObject(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DeleteObject_Call {
	return &Client_DeleteObject_Call{Call: _e.mock.On("DeleteObject",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DeleteObject_Call) Run(run func(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options))) *Client_DeleteObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.DeleteObjectInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DeleteObject_Call) Return(_a0 *s3.DeleteObjectOutput, _a1 error) *Client_DeleteObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DeleteObject_Call) RunAndReturn(run func(context.Context, *s3.DeleteObjectInput, ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)) *Client_DeleteObject_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteObjects provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DeleteObjects(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeleteObjects")
	}

	var r0 *s3.DeleteObjectsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.DeleteObjectsInput, ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.DeleteObjectsInput, ...func(*s3.Options)) *s3.DeleteObjectsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.DeleteObjectsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.DeleteObjectsInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DeleteObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteObjects'
type Client_DeleteObjects_Call struct {
	*mock.Call
}

// DeleteObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.DeleteObjectsInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) DeleteObjects(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DeleteObjects_Call {
	return &Client_DeleteObjects_Call{Call: _e.mock.On("DeleteObjects",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DeleteObjects_Call) Run(run func(ctx context.Context, params *s3.DeleteObjectsInput, optFns ...func(*s3.Options))) *Client_DeleteObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.DeleteObjectsInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DeleteObjects_Call) Return(_a0 *s3.DeleteObjectsOutput, _a1 error) *Client_DeleteObjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DeleteObjects_Call) RunAndReturn(run func(context.Context, *s3.DeleteObjectsInput, ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error)) *Client_DeleteObjects_Call {
	_c.Call.Return(run)
	return _c
}

// GetObject provides a mock function with given fields: ctx, params, optFns
func (_m *Client) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetObject")
	}

	var r0 *s3.GetObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) *s3.GetObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.GetObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_GetObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetObject'
type Client_GetObject_Call struct {
	*mock.Call
}

// GetObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.GetObjectInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) GetObject(ctx interface{}, params interface{}, optFns ...interface{}) *Client_GetObject_Call {
	return &Client_GetObject_Call{Call: _e.mock.On("GetObject",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_GetObject_Call) Run(run func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options))) *Client_GetObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.GetObjectInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_GetObject_Call) Return(_a0 *s3.GetObjectOutput, _a1 error) *Client_GetObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_GetObject_Call) RunAndReturn(run func(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)) *Client_GetObject_Call {
	_c.Call.Return(run)
	return _c
}

// HeadObject provides a mock function with given fields: ctx, params, optFns
func (_m *Client) HeadObject(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for HeadObject")
	}

	var r0 *s3.HeadObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.HeadObjectInput, ...func(*s3.Options)) (*s3.HeadObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.HeadObjectInput, ...func(*s3.Options)) *s3.HeadObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.HeadObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.HeadObjectInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_HeadObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeadObject'
type Client_HeadObject_Call struct {
	*mock.Call
}

// HeadObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.HeadObjectInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) HeadObject(ctx interface{}, params interface{}, optFns ...interface{}) *Client_HeadObject_Call {
	return &Client_HeadObject_Call{Call: _e.mock.On("HeadObject",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_HeadObject_Call) Run(run func(ctx context.Context, params *s3.HeadObjectInput, optFns ...func(*s3.Options))) *Client_HeadObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.HeadObjectInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_HeadObject_Call) Return(_a0 *s3.HeadObjectOutput, _a1 error) *Client_HeadObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_HeadObject_Call) RunAndReturn(run func(context.Context, *s3.HeadObjectInput, ...func(*s3.Options)) (*s3.HeadObjectOutput, error)) *Client_HeadObject_Call {
	_c.Call.Return(run)
	return _c
}

// ListObjects provides a mock function with given fields: ctx, params, optFns
func (_m *Client) ListObjects(ctx context.Context, params *s3.ListObjectsInput, optFns ...func(*s3.Options)) (*s3.ListObjectsOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListObjects")
	}

	var r0 *s3.ListObjectsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) (*s3.ListObjectsOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) *s3.ListObjectsOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListObjects_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListObjects'
type Client_ListObjects_Call struct {
	*mock.Call
}

// ListObjects is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.ListObjectsInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) ListObjects(ctx interface{}, params interface{}, optFns ...interface{}) *Client_ListObjects_Call {
	return &Client_ListObjects_Call{Call: _e.mock.On("ListObjects",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_ListObjects_Call) Run(run func(ctx context.Context, params *s3.ListObjectsInput, optFns ...func(*s3.Options))) *Client_ListObjects_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.ListObjectsInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_ListObjects_Call) Return(_a0 *s3.ListObjectsOutput, _a1 error) *Client_ListObjects_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListObjects_Call) RunAndReturn(run func(context.Context, *s3.ListObjectsInput, ...func(*s3.Options)) (*s3.ListObjectsOutput, error)) *Client_ListObjects_Call {
	_c.Call.Return(run)
	return _c
}

// ListObjectsV2 provides a mock function with given fields: ctx, params, optFns
func (_m *Client) ListObjectsV2(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options)) (*s3.ListObjectsV2Output, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListObjectsV2")
	}

	var r0 *s3.ListObjectsV2Output
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) *s3.ListObjectsV2Output); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.ListObjectsV2Output)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListObjectsV2_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListObjectsV2'
type Client_ListObjectsV2_Call struct {
	*mock.Call
}

// ListObjectsV2 is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.ListObjectsV2Input
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) ListObjectsV2(ctx interface{}, params interface{}, optFns ...interface{}) *Client_ListObjectsV2_Call {
	return &Client_ListObjectsV2_Call{Call: _e.mock.On("ListObjectsV2",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_ListObjectsV2_Call) Run(run func(ctx context.Context, params *s3.ListObjectsV2Input, optFns ...func(*s3.Options))) *Client_ListObjectsV2_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.ListObjectsV2Input), variadicArgs...)
	})
	return _c
}

func (_c *Client_ListObjectsV2_Call) Return(_a0 *s3.ListObjectsV2Output, _a1 error) *Client_ListObjectsV2_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListObjectsV2_Call) RunAndReturn(run func(context.Context, *s3.ListObjectsV2Input, ...func(*s3.Options)) (*s3.ListObjectsV2Output, error)) *Client_ListObjectsV2_Call {
	_c.Call.Return(run)
	return _c
}

// PutObject provides a mock function with given fields: ctx, params, optFns
func (_m *Client) PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutObject")
	}

	var r0 *s3.PutObjectOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) *s3.PutObjectOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_PutObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutObject'
type Client_PutObject_Call struct {
	*mock.Call
}

// PutObject is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.PutObjectInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) PutObject(ctx interface{}, params interface{}, optFns ...interface{}) *Client_PutObject_Call {
	return &Client_PutObject_Call{Call: _e.mock.On("PutObject",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_PutObject_Call) Run(run func(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options))) *Client_PutObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.PutObjectInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_PutObject_Call) Return(_a0 *s3.PutObjectOutput, _a1 error) *Client_PutObject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_PutObject_Call) RunAndReturn(run func(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)) *Client_PutObject_Call {
	_c.Call.Return(run)
	return _c
}

// PutObjectTagging provides a mock function with given fields: ctx, params, optFns
func (_m *Client) PutObjectTagging(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for PutObjectTagging")
	}

	var r0 *s3.PutObjectTaggingOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectTaggingInput, ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectTaggingInput, ...func(*s3.Options)) *s3.PutObjectTaggingOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.PutObjectTaggingOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.PutObjectTaggingInput, ...func(*s3.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_PutObjectTagging_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutObjectTagging'
type Client_PutObjectTagging_Call struct {
	*mock.Call
}

// PutObjectTagging is a helper method to define mock.On call
//   - ctx context.Context
//   - params *s3.PutObjectTaggingInput
//   - optFns ...func(*s3.Options)
func (_e *Client_Expecter) PutObjectTagging(ctx interface{}, params interface{}, optFns ...interface{}) *Client_PutObjectTagging_Call {
	return &Client_PutObjectTagging_Call{Call: _e.mock.On("PutObjectTagging",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_PutObjectTagging_Call) Run(run func(ctx context.Context, params *s3.PutObjectTaggingInput, optFns ...func(*s3.Options))) *Client_PutObjectTagging_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.PutObjectTaggingInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_PutObjectTagging_Call) Return(_a0 *s3.PutObjectTaggingOutput, _a1 error) *Client_PutObjectTagging_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_PutObjectTagging_Call) RunAndReturn(run func(context.Context, *s3.PutObjectTaggingInput, ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error)) *Client_PutObjectTagging_Call {
	_c.Call.Return(run)
	return _c
}

// UploadPart provides a mock function with given fields: _a0, _a1, _a2
func (_m *Client) UploadPart(_a0 context.Context, _a1 *s3.UploadPartInput, _a2 ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UploadPart")
	}

	var r0 *s3.UploadPartOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.UploadPartInput, ...func(*s3.Options)) (*s3.UploadPartOutput, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.UploadPartInput, ...func(*s3.Options)) *s3.UploadPartOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*s3.UploadPartOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.UploadPartInput, ...func(*s3.Options)) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_UploadPart_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadPart'
type Client_UploadPart_Call struct {
	*mock.Call
}

// UploadPart is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *s3.UploadPartInput
//   - _a2 ...func(*s3.Options)
func (_e *Client_Expecter) UploadPart(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *Client_UploadPart_Call {
	return &Client_UploadPart_Call{Call: _e.mock.On("UploadPart",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *Client_UploadPart_Call) Run(run func(_a0 context.Context, _a1 *s3.UploadPartInput, _a2 ...func(*s3.Options))) *Client_UploadPart_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*s3.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*s3.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.UploadPartInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_UploadPart_Call) Return(_a0 *s3.UploadPartOutput, _a1 error) *Client_UploadPart_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_UploadPart_Call) RunAndReturn(run func(context.Context, *s3.UploadPartInput, ...func(*s3.Options)) (*s3.UploadPartOutput, error)) *Client_UploadPart_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
