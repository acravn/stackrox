// Code generated by MockGen. DO NOT EDIT.
// Source: store.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/store.go -source store.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockStore) Count(ctx context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockStoreMockRecorder) Count(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockStore)(nil).Count), ctx)
}

// Exists mocks base method.
func (m *MockStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockStoreMockRecorder) Exists(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockStore)(nil).Exists), ctx, id)
}

// Get mocks base method.
func (m *MockStore) Get(ctx context.Context, id string) (*storage.ClusterCVEEdge, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.ClusterCVEEdge)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), ctx, id)
}

// GetMany mocks base method.
func (m *MockStore) GetMany(ctx context.Context, ids []string) ([]*storage.ClusterCVEEdge, []int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, ids)
	ret0, _ := ret[0].([]*storage.ClusterCVEEdge)
	ret1, _ := ret[1].([]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMany indicates an expected call of GetMany.
func (mr *MockStoreMockRecorder) GetMany(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockStore)(nil).GetMany), ctx, ids)
}
