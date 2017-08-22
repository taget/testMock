// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/taget/testMock/doer (interfaces: Doer)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDoer is a mock of Doer interface
type MockDoer struct {
	ctrl     *gomock.Controller
	recorder *MockDoerMockRecorder
}

// MockDoerMockRecorder is the mock recorder for MockDoer
type MockDoerMockRecorder struct {
	mock *MockDoer
}

// NewMockDoer creates a new mock instance
func NewMockDoer(ctrl *gomock.Controller) *MockDoer {
	mock := &MockDoer{ctrl: ctrl}
	mock.recorder = &MockDoerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDoer) EXPECT() *MockDoerMockRecorder {
	return _m.recorder
}

// DoSomething mocks base method
func (_m *MockDoer) DoSomething(_param0 int, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DoSomething", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoSomething indicates an expected call of DoSomething
func (_mr *MockDoerMockRecorder) DoSomething(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DoSomething", reflect.TypeOf((*MockDoer)(nil).DoSomething), arg0, arg1)
}
