// Code generated by mockery v1.0.0. DO NOT EDIT.

package utils

import mock "github.com/stretchr/testify/mock"

// mockClosable is an autogenerated mock type for the closable type
type mockClosable struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockClosable) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}