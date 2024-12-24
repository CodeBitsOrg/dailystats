// Code generated by mockery v2.50.0. DO NOT EDIT.

package telegram

import mock "github.com/stretchr/testify/mock"

// TelegramProvider is an autogenerated mock type for the TelegramProvider type
type TelegramProvider struct {
	mock.Mock
}

type TelegramProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *TelegramProvider) EXPECT() *TelegramProvider_Expecter {
	return &TelegramProvider_Expecter{mock: &_m.Mock}
}

// Send provides a mock function with given fields: chat_id, message
func (_m *TelegramProvider) Send(chat_id string, message string) error {
	ret := _m.Called(chat_id, message)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(chat_id, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TelegramProvider_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type TelegramProvider_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - chat_id string
//   - message string
func (_e *TelegramProvider_Expecter) Send(chat_id interface{}, message interface{}) *TelegramProvider_Send_Call {
	return &TelegramProvider_Send_Call{Call: _e.mock.On("Send", chat_id, message)}
}

func (_c *TelegramProvider_Send_Call) Run(run func(chat_id string, message string)) *TelegramProvider_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *TelegramProvider_Send_Call) Return(_a0 error) *TelegramProvider_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *TelegramProvider_Send_Call) RunAndReturn(run func(string, string) error) *TelegramProvider_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewTelegramProvider creates a new instance of TelegramProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTelegramProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *TelegramProvider {
	mock := &TelegramProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
