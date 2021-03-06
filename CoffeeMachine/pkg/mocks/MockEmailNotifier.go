// Code generated by MockGen. DO NOT EDIT.
// Source: EmailNotifier.go

// Package mock_infraestructure is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockEmailNotifier is a mock of EmailNotifier interface.
type MockEmailNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockEmailNotifierMockRecorder
}

// MockEmailNotifierMockRecorder is the mock recorder for MockEmailNotifier.
type MockEmailNotifierMockRecorder struct {
	mock *MockEmailNotifier
}

// NewMockEmailNotifier creates a new mock instance.
func NewMockEmailNotifier(ctrl *gomock.Controller) *MockEmailNotifier {
	mock := &MockEmailNotifier{ctrl: ctrl}
	mock.recorder = &MockEmailNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmailNotifier) EXPECT() *MockEmailNotifierMockRecorder {
	return m.recorder
}

// NotifyMissingDrink mocks base method.
func (m *MockEmailNotifier) NotifyMissingDrink(drink string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "NotifyMissingDrink", drink)
}

// NotifyMissingDrink indicates an expected call of NotifyMissingDrink.
func (mr *MockEmailNotifierMockRecorder) NotifyMissingDrink(drink interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyMissingDrink", reflect.TypeOf((*MockEmailNotifier)(nil).NotifyMissingDrink), drink)
}
