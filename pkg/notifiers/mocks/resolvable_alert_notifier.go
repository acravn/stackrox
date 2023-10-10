// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/rox/pkg/notifiers (interfaces: ResolvableAlertNotifier)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/resolvable_alert_notifier.go github.com/stackrox/rox/pkg/notifiers ResolvableAlertNotifier
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockResolvableAlertNotifier is a mock of ResolvableAlertNotifier interface.
type MockResolvableAlertNotifier struct {
	ctrl     *gomock.Controller
	recorder *MockResolvableAlertNotifierMockRecorder
}

// MockResolvableAlertNotifierMockRecorder is the mock recorder for MockResolvableAlertNotifier.
type MockResolvableAlertNotifierMockRecorder struct {
	mock *MockResolvableAlertNotifier
}

// NewMockResolvableAlertNotifier creates a new mock instance.
func NewMockResolvableAlertNotifier(ctrl *gomock.Controller) *MockResolvableAlertNotifier {
	mock := &MockResolvableAlertNotifier{ctrl: ctrl}
	mock.recorder = &MockResolvableAlertNotifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResolvableAlertNotifier) EXPECT() *MockResolvableAlertNotifierMockRecorder {
	return m.recorder
}

// AckAlert mocks base method.
func (m *MockResolvableAlertNotifier) AckAlert(arg0 context.Context, arg1 *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AckAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AckAlert indicates an expected call of AckAlert.
func (mr *MockResolvableAlertNotifierMockRecorder) AckAlert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckAlert", reflect.TypeOf((*MockResolvableAlertNotifier)(nil).AckAlert), arg0, arg1)
}

// AlertNotify mocks base method.
func (m *MockResolvableAlertNotifier) AlertNotify(arg0 context.Context, arg1 *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertNotify", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AlertNotify indicates an expected call of AlertNotify.
func (mr *MockResolvableAlertNotifierMockRecorder) AlertNotify(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertNotify", reflect.TypeOf((*MockResolvableAlertNotifier)(nil).AlertNotify), arg0, arg1)
}

// Close mocks base method.
func (m *MockResolvableAlertNotifier) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockResolvableAlertNotifierMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockResolvableAlertNotifier)(nil).Close), arg0)
}

// ProtoNotifier mocks base method.
func (m *MockResolvableAlertNotifier) ProtoNotifier() *storage.Notifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProtoNotifier")
	ret0, _ := ret[0].(*storage.Notifier)
	return ret0
}

// ProtoNotifier indicates an expected call of ProtoNotifier.
func (mr *MockResolvableAlertNotifierMockRecorder) ProtoNotifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProtoNotifier", reflect.TypeOf((*MockResolvableAlertNotifier)(nil).ProtoNotifier))
}

// ResolveAlert mocks base method.
func (m *MockResolvableAlertNotifier) ResolveAlert(arg0 context.Context, arg1 *storage.Alert) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResolveAlert indicates an expected call of ResolveAlert.
func (mr *MockResolvableAlertNotifierMockRecorder) ResolveAlert(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveAlert", reflect.TypeOf((*MockResolvableAlertNotifier)(nil).ResolveAlert), arg0, arg1)
}

// Test mocks base method.
func (m *MockResolvableAlertNotifier) Test(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Test", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Test indicates an expected call of Test.
func (mr *MockResolvableAlertNotifierMockRecorder) Test(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Test", reflect.TypeOf((*MockResolvableAlertNotifier)(nil).Test), arg0)
}
