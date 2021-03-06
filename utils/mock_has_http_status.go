// Code generated by mockery v1.0.0. DO NOT EDIT.

package utils

import mock "github.com/stretchr/testify/mock"

// MockHasHTTPStatus is an autogenerated mock type for the HasHTTPStatus type
type MockHasHTTPStatus struct {
	mock.Mock
}

// ErrorType provides a mock function with given fields:
func (_m *MockHasHTTPStatus) ErrorType() ErrorType {
	ret := _m.Called()

	var r0 ErrorType
	if rf, ok := ret.Get(0).(func() ErrorType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(ErrorType)
	}

	return r0
}

// HTTPStatusCode provides a mock function with given fields:
func (_m *MockHasHTTPStatus) HTTPStatusCode() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
