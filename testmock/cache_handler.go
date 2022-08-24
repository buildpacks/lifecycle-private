// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpacks/lifecycle (interfaces: CacheHandler)

// Package testmock is a generated GoMock package.
package testmock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	lifecycle "github.com/buildpacks/lifecycle"
)

// MockCacheHandler is a mock of CacheHandler interface.
type MockCacheHandler struct {
	ctrl     *gomock.Controller
	recorder *MockCacheHandlerMockRecorder
}

// MockCacheHandlerMockRecorder is the mock recorder for MockCacheHandler.
type MockCacheHandlerMockRecorder struct {
	mock *MockCacheHandler
}

// NewMockCacheHandler creates a new mock instance.
func NewMockCacheHandler(ctrl *gomock.Controller) *MockCacheHandler {
	mock := &MockCacheHandler{ctrl: ctrl}
	mock.recorder = &MockCacheHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCacheHandler) EXPECT() *MockCacheHandlerMockRecorder {
	return m.recorder
}

// InitCache mocks base method.
func (m *MockCacheHandler) InitCache(arg0, arg1 string) (lifecycle.Cache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitCache", arg0, arg1)
	ret0, _ := ret[0].(lifecycle.Cache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitCache indicates an expected call of InitCache.
func (mr *MockCacheHandlerMockRecorder) InitCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitCache", reflect.TypeOf((*MockCacheHandler)(nil).InitCache), arg0, arg1)
}