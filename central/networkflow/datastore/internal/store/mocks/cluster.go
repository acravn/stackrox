// Code generated by MockGen. DO NOT EDIT.
// Source: cluster.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	store "github.com/stackrox/rox/central/networkflow/datastore/internal/store"
	reflect "reflect"
)

// MockClusterStore is a mock of ClusterStore interface
type MockClusterStore struct {
	ctrl     *gomock.Controller
	recorder *MockClusterStoreMockRecorder
}

// MockClusterStoreMockRecorder is the mock recorder for MockClusterStore
type MockClusterStoreMockRecorder struct {
	mock *MockClusterStore
}

// NewMockClusterStore creates a new mock instance
func NewMockClusterStore(ctrl *gomock.Controller) *MockClusterStore {
	mock := &MockClusterStore{ctrl: ctrl}
	mock.recorder = &MockClusterStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterStore) EXPECT() *MockClusterStoreMockRecorder {
	return m.recorder
}

// GetFlowStore mocks base method
func (m *MockClusterStore) GetFlowStore(clusterID string) store.FlowStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlowStore", clusterID)
	ret0, _ := ret[0].(store.FlowStore)
	return ret0
}

// GetFlowStore indicates an expected call of GetFlowStore
func (mr *MockClusterStoreMockRecorder) GetFlowStore(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlowStore", reflect.TypeOf((*MockClusterStore)(nil).GetFlowStore), clusterID)
}

// CreateFlowStore mocks base method
func (m *MockClusterStore) CreateFlowStore(clusterID string) (store.FlowStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFlowStore", clusterID)
	ret0, _ := ret[0].(store.FlowStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFlowStore indicates an expected call of CreateFlowStore
func (mr *MockClusterStoreMockRecorder) CreateFlowStore(clusterID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFlowStore", reflect.TypeOf((*MockClusterStore)(nil).CreateFlowStore), clusterID)
}
