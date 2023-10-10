// Code generated by MockGen. DO NOT EDIT.
// Source: searcher.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/searcher.go -source searcher.go
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

// MockSearcher is a mock of Searcher interface.
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher.
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance.
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockSearcher) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSearcherMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSearcher)(nil).Count), ctx, q)
}

// Search mocks base method.
func (m *MockSearcher) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockSearcherMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSearcher)(nil).Search), ctx, q)
}

// SearchNodes mocks base method.
func (m *MockSearcher) SearchNodes(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchNodes", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchNodes indicates an expected call of SearchNodes.
func (mr *MockSearcherMockRecorder) SearchNodes(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchNodes", reflect.TypeOf((*MockSearcher)(nil).SearchNodes), ctx, q)
}

// SearchRawNodes mocks base method.
func (m *MockSearcher) SearchRawNodes(ctx context.Context, q *v1.Query) ([]*storage.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawNodes", ctx, q)
	ret0, _ := ret[0].([]*storage.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawNodes indicates an expected call of SearchRawNodes.
func (mr *MockSearcherMockRecorder) SearchRawNodes(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawNodes", reflect.TypeOf((*MockSearcher)(nil).SearchRawNodes), ctx, q)
}
