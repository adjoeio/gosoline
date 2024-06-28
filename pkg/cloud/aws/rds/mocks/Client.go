// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	rds "github.com/aws/aws-sdk-go-v2/service/rds"
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

// DescribeDBInstances provides a mock function with given fields: ctx, params, optFns
func (_m *Client) DescribeDBInstances(ctx context.Context, params *rds.DescribeDBInstancesInput, optFns ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DescribeDBInstances")
	}

	var r0 *rds.DescribeDBInstancesOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *rds.DescribeDBInstancesInput, ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *rds.DescribeDBInstancesInput, ...func(*rds.Options)) *rds.DescribeDBInstancesOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*rds.DescribeDBInstancesOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *rds.DescribeDBInstancesInput, ...func(*rds.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_DescribeDBInstances_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeDBInstances'
type Client_DescribeDBInstances_Call struct {
	*mock.Call
}

// DescribeDBInstances is a helper method to define mock.On call
//   - ctx context.Context
//   - params *rds.DescribeDBInstancesInput
//   - optFns ...func(*rds.Options)
func (_e *Client_Expecter) DescribeDBInstances(ctx interface{}, params interface{}, optFns ...interface{}) *Client_DescribeDBInstances_Call {
	return &Client_DescribeDBInstances_Call{Call: _e.mock.On("DescribeDBInstances",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *Client_DescribeDBInstances_Call) Run(run func(ctx context.Context, params *rds.DescribeDBInstancesInput, optFns ...func(*rds.Options))) *Client_DescribeDBInstances_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*rds.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*rds.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*rds.DescribeDBInstancesInput), variadicArgs...)
	})
	return _c
}

func (_c *Client_DescribeDBInstances_Call) Return(_a0 *rds.DescribeDBInstancesOutput, _a1 error) *Client_DescribeDBInstances_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_DescribeDBInstances_Call) RunAndReturn(run func(context.Context, *rds.DescribeDBInstancesInput, ...func(*rds.Options)) (*rds.DescribeDBInstancesOutput, error)) *Client_DescribeDBInstances_Call {
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
