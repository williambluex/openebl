// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bu_server/auth/application.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "github.com/openebl/openebl/pkg/bu_server/auth"
	storage "github.com/openebl/openebl/pkg/bu_server/storage"
)

// MockApplicationManager is a mock of ApplicationManager interface.
type MockApplicationManager struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationManagerMockRecorder
}

// MockApplicationManagerMockRecorder is the mock recorder for MockApplicationManager.
type MockApplicationManagerMockRecorder struct {
	mock *MockApplicationManager
}

// NewMockApplicationManager creates a new mock instance.
func NewMockApplicationManager(ctrl *gomock.Controller) *MockApplicationManager {
	mock := &MockApplicationManager{ctrl: ctrl}
	mock.recorder = &MockApplicationManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationManager) EXPECT() *MockApplicationManagerMockRecorder {
	return m.recorder
}

// ActivateApplication mocks base method.
func (m *MockApplicationManager) ActivateApplication(ctx context.Context, ts int64, req auth.ActivateApplicationRequest) (auth.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateApplication", ctx, ts, req)
	ret0, _ := ret[0].(auth.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateApplication indicates an expected call of ActivateApplication.
func (mr *MockApplicationManagerMockRecorder) ActivateApplication(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateApplication", reflect.TypeOf((*MockApplicationManager)(nil).ActivateApplication), ctx, ts, req)
}

// CreateApplication mocks base method.
func (m *MockApplicationManager) CreateApplication(ctx context.Context, ts int64, req auth.CreateApplicationRequest) (auth.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateApplication", ctx, ts, req)
	ret0, _ := ret[0].(auth.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateApplication indicates an expected call of CreateApplication.
func (mr *MockApplicationManagerMockRecorder) CreateApplication(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateApplication", reflect.TypeOf((*MockApplicationManager)(nil).CreateApplication), ctx, ts, req)
}

// DeactivateApplication mocks base method.
func (m *MockApplicationManager) DeactivateApplication(ctx context.Context, ts int64, req auth.DeactivateApplicationRequest) (auth.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateApplication", ctx, ts, req)
	ret0, _ := ret[0].(auth.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateApplication indicates an expected call of DeactivateApplication.
func (mr *MockApplicationManagerMockRecorder) DeactivateApplication(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateApplication", reflect.TypeOf((*MockApplicationManager)(nil).DeactivateApplication), ctx, ts, req)
}

// ListApplications mocks base method.
func (m *MockApplicationManager) ListApplications(ctx context.Context, req auth.ListApplicationRequest) (auth.ListApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplications", ctx, req)
	ret0, _ := ret[0].(auth.ListApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockApplicationManagerMockRecorder) ListApplications(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockApplicationManager)(nil).ListApplications), ctx, req)
}

// UpdateApplication mocks base method.
func (m *MockApplicationManager) UpdateApplication(ctx context.Context, ts int64, req auth.UpdateApplicationRequest) (auth.Application, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApplication", ctx, ts, req)
	ret0, _ := ret[0].(auth.Application)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApplication indicates an expected call of UpdateApplication.
func (mr *MockApplicationManagerMockRecorder) UpdateApplication(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockApplicationManager)(nil).UpdateApplication), ctx, ts, req)
}

// MockApplicationStorage is a mock of ApplicationStorage interface.
type MockApplicationStorage struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationStorageMockRecorder
}

// MockApplicationStorageMockRecorder is the mock recorder for MockApplicationStorage.
type MockApplicationStorageMockRecorder struct {
	mock *MockApplicationStorage
}

// NewMockApplicationStorage creates a new mock instance.
func NewMockApplicationStorage(ctrl *gomock.Controller) *MockApplicationStorage {
	mock := &MockApplicationStorage{ctrl: ctrl}
	mock.recorder = &MockApplicationStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationStorage) EXPECT() *MockApplicationStorageMockRecorder {
	return m.recorder
}

// CreateTx mocks base method.
func (m *MockApplicationStorage) CreateTx(ctx context.Context, options ...storage.CreateTxOption) (storage.Tx, context.Context, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTx", varargs...)
	ret0, _ := ret[0].(storage.Tx)
	ret1, _ := ret[1].(context.Context)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockApplicationStorageMockRecorder) CreateTx(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockApplicationStorage)(nil).CreateTx), varargs...)
}

// ListApplication mocks base method.
func (m *MockApplicationStorage) ListApplication(ctx context.Context, tx storage.Tx, req auth.ListApplicationRequest) (auth.ListApplicationResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListApplication", ctx, tx, req)
	ret0, _ := ret[0].(auth.ListApplicationResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplication indicates an expected call of ListApplication.
func (mr *MockApplicationStorageMockRecorder) ListApplication(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplication", reflect.TypeOf((*MockApplicationStorage)(nil).ListApplication), ctx, tx, req)
}

// StoreApplication mocks base method.
func (m *MockApplicationStorage) StoreApplication(ctx context.Context, tx storage.Tx, app auth.Application) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreApplication", ctx, tx, app)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreApplication indicates an expected call of StoreApplication.
func (mr *MockApplicationStorageMockRecorder) StoreApplication(ctx, tx, app interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreApplication", reflect.TypeOf((*MockApplicationStorage)(nil).StoreApplication), ctx, tx, app)
}
