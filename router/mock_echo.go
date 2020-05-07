// Code generated by mockery v1.0.0. DO NOT EDIT.

package router

import context "context"
import echo "github.com/labstack/echo/v4"
import mock "github.com/stretchr/testify/mock"

// MockEcho is an autogenerated mock type for the Echo type
type MockEcho struct {
	mock.Mock
}

// GET provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockEcho) GET(_a0 string, _a1 echo.HandlerFunc, _a2 ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// Group provides a mock function with given fields: prefix, m
func (_m *MockEcho) Group(prefix string, m ...echo.MiddlewareFunc) *echo.Group {
	_va := make([]interface{}, len(m))
	for _i := range m {
		_va[_i] = m[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, prefix)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Group
	if rf, ok := ret.Get(0).(func(string, ...echo.MiddlewareFunc) *echo.Group); ok {
		r0 = rf(prefix, m...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Group)
		}
	}

	return r0
}

// POST provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockEcho) POST(_a0 string, _a1 echo.HandlerFunc, _a2 ...echo.MiddlewareFunc) *echo.Route {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...echo.MiddlewareFunc) *echo.Route); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// Shutdown provides a mock function with given fields: ctx
func (_m *MockEcho) Shutdown(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: address
func (_m *MockEcho) Start(address string) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Static provides a mock function with given fields: prefix, root
func (_m *MockEcho) Static(prefix string, root string) *echo.Route {
	ret := _m.Called(prefix, root)

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, string) *echo.Route); ok {
		r0 = rf(prefix, root)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}
