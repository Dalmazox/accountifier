// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repositories/tx.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITx is a mock of ITx interface.
type MockITx struct {
	ctrl     *gomock.Controller
	recorder *MockITxMockRecorder
}

// MockITxMockRecorder is the mock recorder for MockITx.
type MockITxMockRecorder struct {
	mock *MockITx
}

// NewMockITx creates a new mock instance.
func NewMockITx(ctrl *gomock.Controller) *MockITx {
	mock := &MockITx{ctrl: ctrl}
	mock.recorder = &MockITxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITx) EXPECT() *MockITxMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockITx) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockITxMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockITx)(nil).Commit))
}

// ExecContext mocks base method.
func (m *MockITx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ExecContext", varargs...)
	ret0, _ := ret[0].(sql.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecContext indicates an expected call of ExecContext.
func (mr *MockITxMockRecorder) ExecContext(ctx, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecContext", reflect.TypeOf((*MockITx)(nil).ExecContext), varargs...)
}

// GetContext mocks base method.
func (m *MockITx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, dest, query}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetContext", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetContext indicates an expected call of GetContext.
func (mr *MockITxMockRecorder) GetContext(ctx, dest, query interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, dest, query}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContext", reflect.TypeOf((*MockITx)(nil).GetContext), varargs...)
}

// Rollback mocks base method.
func (m *MockITx) Rollback() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rollback")
	ret0, _ := ret[0].(error)
	return ret0
}

// Rollback indicates an expected call of Rollback.
func (mr *MockITxMockRecorder) Rollback() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rollback", reflect.TypeOf((*MockITx)(nil).Rollback))
}
