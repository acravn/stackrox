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

// AddReportSnapshot mocks base method.
func (m *MockDataStore) AddReportSnapshot(ctx context.Context, snapshot *storage.ReportSnapshot) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddReportSnapshot", ctx, snapshot)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddReportSnapshot indicates an expected call of AddReportSnapshot.
func (mr *MockDataStoreMockRecorder) AddReportSnapshot(ctx, snapshot any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddReportSnapshot", reflect.TypeOf((*MockDataStore)(nil).AddReportSnapshot), ctx, snapshot)
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

// DeleteReportSnapshot mocks base method.
func (m *MockDataStore) DeleteReportSnapshot(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReportSnapshot", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteReportSnapshot indicates an expected call of DeleteReportSnapshot.
func (mr *MockDataStoreMockRecorder) DeleteReportSnapshot(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReportSnapshot", reflect.TypeOf((*MockDataStore)(nil).DeleteReportSnapshot), ctx, id)
}

// Exists mocks base method.
func (m *MockDataStore) Exists(ctx context.Context, id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDataStoreMockRecorder) Exists(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDataStore)(nil).Exists), ctx, id)
}

// Get mocks base method.
func (m *MockDataStore) Get(ctx context.Context, id string) (*storage.ReportSnapshot, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*storage.ReportSnapshot)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockDataStoreMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDataStore)(nil).Get), ctx, id)
}

// GetMany mocks base method.
func (m *MockDataStore) GetMany(ctx context.Context, ids []string) ([]*storage.ReportSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMany", ctx, ids)
	ret0, _ := ret[0].([]*storage.ReportSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMany indicates an expected call of GetMany.
func (mr *MockDataStoreMockRecorder) GetMany(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMany", reflect.TypeOf((*MockDataStore)(nil).GetMany), ctx, ids)
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

// SearchReportSnapshots mocks base method.
func (m *MockDataStore) SearchReportSnapshots(ctx context.Context, q *v1.Query) ([]*storage.ReportSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchReportSnapshots", ctx, q)
	ret0, _ := ret[0].([]*storage.ReportSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchReportSnapshots indicates an expected call of SearchReportSnapshots.
func (mr *MockDataStoreMockRecorder) SearchReportSnapshots(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchReportSnapshots", reflect.TypeOf((*MockDataStore)(nil).SearchReportSnapshots), ctx, q)
}

// SearchResults mocks base method.
func (m *MockDataStore) SearchResults(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchResults", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchResults indicates an expected call of SearchResults.
func (mr *MockDataStoreMockRecorder) SearchResults(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchResults", reflect.TypeOf((*MockDataStore)(nil).SearchResults), ctx, q)
}

// UpdateReportSnapshot mocks base method.
func (m *MockDataStore) UpdateReportSnapshot(ctx context.Context, snapshot *storage.ReportSnapshot) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReportSnapshot", ctx, snapshot)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateReportSnapshot indicates an expected call of UpdateReportSnapshot.
func (mr *MockDataStoreMockRecorder) UpdateReportSnapshot(ctx, snapshot any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReportSnapshot", reflect.TypeOf((*MockDataStore)(nil).UpdateReportSnapshot), ctx, snapshot)
}

// Walk mocks base method.
func (m *MockDataStore) Walk(ctx context.Context, fn func(*storage.ReportSnapshot) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Walk", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// Walk indicates an expected call of Walk.
func (mr *MockDataStoreMockRecorder) Walk(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Walk", reflect.TypeOf((*MockDataStore)(nil).Walk), ctx, fn)
}
