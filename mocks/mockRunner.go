// Code generated by MockGen. DO NOT EDIT.
// Source: sample/mocktest/Test_User (interfaces: TestUserInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTestUserInterface is a mock of TestUserInterface interface
type MockTestUserInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTestUserInterfaceMockRecorder
}

// MockTestUserInterfaceMockRecorder is the mock recorder for MockTestUserInterface
type MockTestUserInterfaceMockRecorder struct {
	mock *MockTestUserInterface
}

// NewMockTestUserInterface creates a new mock instance
func NewMockTestUserInterface(ctrl *gomock.Controller) *MockTestUserInterface {
	mock := &MockTestUserInterface{ctrl: ctrl}
	mock.recorder = &MockTestUserInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTestUserInterface) EXPECT() *MockTestUserInterfaceMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockTestUserInterface) AddUser(arg0 int, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockTestUserInterfaceMockRecorder) AddUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockTestUserInterface)(nil).AddUser), arg0, arg1)
}