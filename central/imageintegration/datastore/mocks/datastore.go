// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddImageIntegration mocks base method.
func (m *MockDataStore) AddImageIntegration(ctx context.Context, integration *storage.ImageIntegration) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddImageIntegration", ctx, integration)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddImageIntegration indicates an expected call of AddImageIntegration.
func (mr *MockDataStoreMockRecorder) AddImageIntegration(ctx, integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddImageIntegration", reflect.TypeOf((*MockDataStore)(nil).AddImageIntegration), ctx, integration)
}

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// GetImageIntegration mocks base method.
func (m *MockDataStore) GetImageIntegration(ctx context.Context, id string) (*storage.ImageIntegration, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageIntegration", ctx, id)
	ret0, _ := ret[0].(*storage.ImageIntegration)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetImageIntegration indicates an expected call of GetImageIntegration.
func (mr *MockDataStoreMockRecorder) GetImageIntegration(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageIntegration", reflect.TypeOf((*MockDataStore)(nil).GetImageIntegration), ctx, id)
}

// GetImageIntegrations mocks base method.
func (m *MockDataStore) GetImageIntegrations(ctx context.Context, integration *v1.GetImageIntegrationsRequest) ([]*storage.ImageIntegration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImageIntegrations", ctx, integration)
	ret0, _ := ret[0].([]*storage.ImageIntegration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImageIntegrations indicates an expected call of GetImageIntegrations.
func (mr *MockDataStoreMockRecorder) GetImageIntegrations(ctx, integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageIntegrations", reflect.TypeOf((*MockDataStore)(nil).GetImageIntegrations), ctx, integration)
}

// RemoveImageIntegration mocks base method.
func (m *MockDataStore) RemoveImageIntegration(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveImageIntegration", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImageIntegration indicates an expected call of RemoveImageIntegration.
func (mr *MockDataStoreMockRecorder) RemoveImageIntegration(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImageIntegration", reflect.TypeOf((*MockDataStore)(nil).RemoveImageIntegration), ctx, id)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchImageIntegrations mocks base method.
func (m *MockDataStore) SearchImageIntegrations(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchImageIntegrations", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchImageIntegrations indicates an expected call of SearchImageIntegrations.
func (mr *MockDataStoreMockRecorder) SearchImageIntegrations(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchImageIntegrations", reflect.TypeOf((*MockDataStore)(nil).SearchImageIntegrations), ctx, q)
}

// UpdateImageIntegration mocks base method.
func (m *MockDataStore) UpdateImageIntegration(ctx context.Context, integration *storage.ImageIntegration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateImageIntegration", ctx, integration)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateImageIntegration indicates an expected call of UpdateImageIntegration.
func (mr *MockDataStoreMockRecorder) UpdateImageIntegration(ctx, integration any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateImageIntegration", reflect.TypeOf((*MockDataStore)(nil).UpdateImageIntegration), ctx, integration)
}
