// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/incli (interfaces: Handler)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	cli "github.com/eneskzlcn/incli"
	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// CommandRun mocks base method.
func (m *MockHandler) CommandRun(arg0 *cli.Command, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CommandRun", arg0, arg1)
}

// CommandRun indicates an expected call of CommandRun.
func (mr *MockHandlerMockRecorder) CommandRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommandRun", reflect.TypeOf((*MockHandler)(nil).CommandRun), arg0, arg1)
}

// ExecuteCommand mocks base method.
func (m *MockHandler) ExecuteCommand() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteCommand")
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteCommand indicates an expected call of ExecuteCommand.
func (mr *MockHandlerMockRecorder) ExecuteCommand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteCommand", reflect.TypeOf((*MockHandler)(nil).ExecuteCommand))
}

// GetCommand mocks base method.
func (m *MockHandler) GetCommand() *cli.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommand")
	ret0, _ := ret[0].(*cli.Command)
	return ret0
}

// GetCommand indicates an expected call of GetCommand.
func (mr *MockHandlerMockRecorder) GetCommand() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommand", reflect.TypeOf((*MockHandler)(nil).GetCommand))
}

// GetName mocks base method.
func (m *MockHandler) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockHandlerMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockHandler)(nil).GetName))
}