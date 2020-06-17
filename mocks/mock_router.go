// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/reverted/router (interfaces: Router)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// Route mocks base method
func (m *MockRouter) Route(arg0 *http.Request) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Route indicates an expected call of Route
func (mr *MockRouterMockRecorder) Route(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouter)(nil).Route), arg0)
}
