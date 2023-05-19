// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/pkg/notifiers (interfaces: AlertNotifier)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockAlertNotifier is a mock of AlertNotifier interface.
type MockAlertNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockAlertNotifierMockRecorder
}

// MockAlertNotifierMockRecorder is the mock recorder for MockAlertNotifier.
type MockAlertNotifierMockRecorder struct {
	mock *MockAlertNotifier
}

// NewMockAlertNotifier creates a new mock instance.
func NewMockAlertNotifier(ctrl *gomock.Controller) *MockAlertNotifier {
	mock := &MockAlertNotifier{ctrl: ctrl}
	mock.recorder = &MockAlertNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlertNotifier) EXPECT() *MockAlertNotifierMockRecorder {
	return m.recorder
}

// AlertNotify mocks base method.
func (m *MockAlertNotifier) AlertNotify(arg0 context.Context, arg1 *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertNotify", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlertNotify indicates an expected call of AlertNotify.
func (mr *MockAlertNotifierMockRecorder) AlertNotify(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertNotify", reflect.TypeOf((*MockAlertNotifier)(nil).AlertNotify), arg0, arg1)
}

// Close mocks base method.
func (m *MockAlertNotifier) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAlertNotifierMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAlertNotifier)(nil).Close), arg0)
}

// ProtoNotifier mocks base method.
func (m *MockAlertNotifier) ProtoNotifier() *storage.Notifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoNotifier")
	ret0, _ := ret[0].(*storage.Notifier)
	return ret0
}

// ProtoNotifier indicates an expected call of ProtoNotifier.
func (mr *MockAlertNotifierMockRecorder) ProtoNotifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoNotifier", reflect.TypeOf((*MockAlertNotifier)(nil).ProtoNotifier))
}

// Test mocks base method.
func (m *MockAlertNotifier) Test(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Test", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Test indicates an expected call of Test.
func (mr *MockAlertNotifierMockRecorder) Test(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockAlertNotifier)(nil).Test), arg0)
}
