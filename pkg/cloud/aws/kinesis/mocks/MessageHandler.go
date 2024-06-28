// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// MessageHandler is an autogenerated mock type for the MessageHandler type
type MessageHandler struct {
	mock.Mock
}

type MessageHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MessageHandler) EXPECT() *MessageHandler_Expecter {
	return &MessageHandler_Expecter{mock: &_m.Mock}
}

// Done provides a mock function with given fields:
func (_m *MessageHandler) Done() {
	_m.Called()
}

// MessageHandler_Done_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Done'
type MessageHandler_Done_Call struct {
	*mock.Call
}

// Done is a helper method to define mock.On call
func (_e *MessageHandler_Expecter) Done() *MessageHandler_Done_Call {
	return &MessageHandler_Done_Call{Call: _e.mock.On("Done")}
}

func (_c *MessageHandler_Done_Call) Run(run func()) *MessageHandler_Done_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MessageHandler_Done_Call) Return() *MessageHandler_Done_Call {
	_c.Call.Return()
	return _c
}

func (_c *MessageHandler_Done_Call) RunAndReturn(run func()) *MessageHandler_Done_Call {
	_c.Call.Return(run)
	return _c
}

// Handle provides a mock function with given fields: rawMessage
func (_m *MessageHandler) Handle(rawMessage []byte) error {
	ret := _m.Called(rawMessage)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(rawMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MessageHandler_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MessageHandler_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - rawMessage []byte
func (_e *MessageHandler_Expecter) Handle(rawMessage interface{}) *MessageHandler_Handle_Call {
	return &MessageHandler_Handle_Call{Call: _e.mock.On("Handle", rawMessage)}
}

func (_c *MessageHandler_Handle_Call) Run(run func(rawMessage []byte)) *MessageHandler_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *MessageHandler_Handle_Call) Return(_a0 error) *MessageHandler_Handle_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MessageHandler_Handle_Call) RunAndReturn(run func([]byte) error) *MessageHandler_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMessageHandler creates a new instance of MessageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageHandler {
	mock := &MessageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
