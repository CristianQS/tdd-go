// Code generated by MockGen. DO NOT EDIT.
// Source: ../Infraestructure/Logger.go

// Package mock_Infraestructure is a generated GoMock package.
package Mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Print mocks base method.
func (m *MockLogger) Print(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Print", message)
}

// Print indicates an expected call of Print.
func (mr *MockLoggerMockRecorder) Print(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockLogger)(nil).Print), message)
}
