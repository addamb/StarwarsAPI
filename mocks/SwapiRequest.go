// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// SwapiRequest is an autogenerated mock type for the SwapiRequest type
type SwapiRequest struct {
	mock.Mock
}

type SwapiRequest_Expecter struct {
	mock *mock.Mock
}

func (_m *SwapiRequest) EXPECT() *SwapiRequest_Expecter {
	return &SwapiRequest_Expecter{mock: &_m.Mock}
}

// SendSwapiRequest provides a mock function with given fields: reqURL
func (_m *SwapiRequest) SendSwapiRequest(reqURL string) (io.ReadCloser, error) {
	ret := _m.Called(reqURL)

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (io.ReadCloser, error)); ok {
		return rf(reqURL)
	}
	if rf, ok := ret.Get(0).(func(string) io.ReadCloser); ok {
		r0 = rf(reqURL)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(reqURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SwapiRequest_SendSwapiRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendSwapiRequest'
type SwapiRequest_SendSwapiRequest_Call struct {
	*mock.Call
}

// SendSwapiRequest is a helper method to define mock.On call
//   - reqURL string
func (_e *SwapiRequest_Expecter) SendSwapiRequest(reqURL interface{}) *SwapiRequest_SendSwapiRequest_Call {
	return &SwapiRequest_SendSwapiRequest_Call{Call: _e.mock.On("SendSwapiRequest", reqURL)}
}

func (_c *SwapiRequest_SendSwapiRequest_Call) Run(run func(reqURL string)) *SwapiRequest_SendSwapiRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *SwapiRequest_SendSwapiRequest_Call) Return(_a0 io.ReadCloser, _a1 error) *SwapiRequest_SendSwapiRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SwapiRequest_SendSwapiRequest_Call) RunAndReturn(run func(string) (io.ReadCloser, error)) *SwapiRequest_SendSwapiRequest_Call {
	_c.Call.Return(run)
	return _c
}

// NewSwapiRequest creates a new instance of SwapiRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSwapiRequest(t interface {
	mock.TestingT
	Cleanup(func())
}) *SwapiRequest {
	mock := &SwapiRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
