// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// ListImage mocks base method
func (m *MockStore) ListImage(sha string) (*storage.ListImage, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImage", sha)
	ret0, _ := ret[0].(*storage.ListImage)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListImage indicates an expected call of ListImage
func (mr *MockStoreMockRecorder) ListImage(sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImage", reflect.TypeOf((*MockStore)(nil).ListImage), sha)
}

// GetImages mocks base method
func (m *MockStore) GetImages() ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImages")
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImages indicates an expected call of GetImages
func (mr *MockStoreMockRecorder) GetImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImages", reflect.TypeOf((*MockStore)(nil).GetImages))
}

// CountImages mocks base method
func (m *MockStore) CountImages() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountImages")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountImages indicates an expected call of CountImages
func (mr *MockStoreMockRecorder) CountImages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountImages", reflect.TypeOf((*MockStore)(nil).CountImages))
}

// GetImage mocks base method
func (m *MockStore) GetImage(sha string) (*storage.Image, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImage", sha)
	ret0, _ := ret[0].(*storage.Image)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetImage indicates an expected call of GetImage
func (mr *MockStoreMockRecorder) GetImage(sha interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImage", reflect.TypeOf((*MockStore)(nil).GetImage), sha)
}

// GetImagesBatch mocks base method
func (m *MockStore) GetImagesBatch(shas []string) ([]*storage.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagesBatch", shas)
	ret0, _ := ret[0].([]*storage.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetImagesBatch indicates an expected call of GetImagesBatch
func (mr *MockStoreMockRecorder) GetImagesBatch(shas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagesBatch", reflect.TypeOf((*MockStore)(nil).GetImagesBatch), shas)
}

// Exists mocks base method
func (m *MockStore) Exists(id string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists
func (mr *MockStoreMockRecorder) Exists(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockStore)(nil).Exists), id)
}

// UpsertImage mocks base method
func (m *MockStore) UpsertImage(image *storage.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertImage", image)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertImage indicates an expected call of UpsertImage
func (mr *MockStoreMockRecorder) UpsertImage(image interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertImage", reflect.TypeOf((*MockStore)(nil).UpsertImage), image)
}

// DeleteImage mocks base method
func (m *MockStore) DeleteImage(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteImage", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteImage indicates an expected call of DeleteImage
func (mr *MockStoreMockRecorder) DeleteImage(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteImage", reflect.TypeOf((*MockStore)(nil).DeleteImage), id)
}

// GetTxnCount mocks base method
func (m *MockStore) GetTxnCount() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxnCount")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTxnCount indicates an expected call of GetTxnCount
func (mr *MockStoreMockRecorder) GetTxnCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxnCount", reflect.TypeOf((*MockStore)(nil).GetTxnCount))
}

// IncTxnCount mocks base method
func (m *MockStore) IncTxnCount() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncTxnCount")
	ret0, _ := ret[0].(error)
	return ret0
}

// IncTxnCount indicates an expected call of IncTxnCount
func (mr *MockStoreMockRecorder) IncTxnCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncTxnCount", reflect.TypeOf((*MockStore)(nil).IncTxnCount))
}
