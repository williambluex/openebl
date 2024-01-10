// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/bu_server/auth/user.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "github.com/openebl/openebl/pkg/bu_server/auth"
	storage "github.com/openebl/openebl/pkg/bu_server/storage"
)

// MockUserManager is a mock of UserManager interface.
type MockUserManager struct {
	ctrl     *gomock.Controller
	recorder *MockUserManagerMockRecorder
}

// MockUserManagerMockRecorder is the mock recorder for MockUserManager.
type MockUserManagerMockRecorder struct {
	mock *MockUserManager
}

// NewMockUserManager creates a new mock instance.
func NewMockUserManager(ctrl *gomock.Controller) *MockUserManager {
	mock := &MockUserManager{ctrl: ctrl}
	mock.recorder = &MockUserManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserManager) EXPECT() *MockUserManagerMockRecorder {
	return m.recorder
}

// ActivateUser mocks base method.
func (m *MockUserManager) ActivateUser(ctx context.Context, ts int64, req auth.ActivateUserRequest) (auth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActivateUser", ctx, ts, req)
	ret0, _ := ret[0].(auth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateUser indicates an expected call of ActivateUser.
func (mr *MockUserManagerMockRecorder) ActivateUser(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateUser", reflect.TypeOf((*MockUserManager)(nil).ActivateUser), ctx, ts, req)
}

// Authenticate mocks base method.
func (m *MockUserManager) Authenticate(ctx context.Context, ts int64, req auth.AuthenticateUserRequest) (auth.UserToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", ctx, ts, req)
	ret0, _ := ret[0].(auth.UserToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockUserManagerMockRecorder) Authenticate(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockUserManager)(nil).Authenticate), ctx, ts, req)
}

// ChangePassword mocks base method.
func (m *MockUserManager) ChangePassword(ctx context.Context, ts int64, req auth.ChangePasswordRequest) (auth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, ts, req)
	ret0, _ := ret[0].(auth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserManagerMockRecorder) ChangePassword(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserManager)(nil).ChangePassword), ctx, ts, req)
}

// CreateUser mocks base method.
func (m *MockUserManager) CreateUser(ctx context.Context, ts int64, req auth.CreateUserRequest) (auth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, ts, req)
	ret0, _ := ret[0].(auth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserManagerMockRecorder) CreateUser(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserManager)(nil).CreateUser), ctx, ts, req)
}

// DeactivateUser mocks base method.
func (m *MockUserManager) DeactivateUser(ctx context.Context, ts int64, req auth.ActivateUserRequest) (auth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactivateUser", ctx, ts, req)
	ret0, _ := ret[0].(auth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeactivateUser indicates an expected call of DeactivateUser.
func (mr *MockUserManagerMockRecorder) DeactivateUser(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactivateUser", reflect.TypeOf((*MockUserManager)(nil).DeactivateUser), ctx, ts, req)
}

// ListUsers mocks base method.
func (m *MockUserManager) ListUsers(ctx context.Context, req auth.ListUserRequest) (auth.ListUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, req)
	ret0, _ := ret[0].(auth.ListUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockUserManagerMockRecorder) ListUsers(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserManager)(nil).ListUsers), ctx, req)
}

// ResetPassword mocks base method.
func (m *MockUserManager) ResetPassword(ctx context.Context, ts int64, req auth.ResetPasswordRequest) (auth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPassword", ctx, ts, req)
	ret0, _ := ret[0].(auth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetPassword indicates an expected call of ResetPassword.
func (mr *MockUserManagerMockRecorder) ResetPassword(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPassword", reflect.TypeOf((*MockUserManager)(nil).ResetPassword), ctx, ts, req)
}

// TokenAuthorization mocks base method.
func (m *MockUserManager) TokenAuthorization(ctx context.Context, ts int64, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenAuthorization", ctx, ts, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// TokenAuthorization indicates an expected call of TokenAuthorization.
func (mr *MockUserManagerMockRecorder) TokenAuthorization(ctx, ts, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenAuthorization", reflect.TypeOf((*MockUserManager)(nil).TokenAuthorization), ctx, ts, token)
}

// UpdateUser mocks base method.
func (m *MockUserManager) UpdateUser(ctx context.Context, ts int64, req auth.UpdateUserRequest) (auth.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, ts, req)
	ret0, _ := ret[0].(auth.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserManagerMockRecorder) UpdateUser(ctx, ts, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserManager)(nil).UpdateUser), ctx, ts, req)
}

// MockUserStorage is a mock of UserStorage interface.
type MockUserStorage struct {
	ctrl     *gomock.Controller
	recorder *MockUserStorageMockRecorder
}

// MockUserStorageMockRecorder is the mock recorder for MockUserStorage.
type MockUserStorageMockRecorder struct {
	mock *MockUserStorage
}

// NewMockUserStorage creates a new mock instance.
func NewMockUserStorage(ctrl *gomock.Controller) *MockUserStorage {
	mock := &MockUserStorage{ctrl: ctrl}
	mock.recorder = &MockUserStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserStorage) EXPECT() *MockUserStorageMockRecorder {
	return m.recorder
}

// CreateTx mocks base method.
func (m *MockUserStorage) CreateTx(ctx context.Context, options ...storage.CreateTxOption) (storage.Tx, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range options {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTx", varargs...)
	ret0, _ := ret[0].(storage.Tx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTx indicates an expected call of CreateTx.
func (mr *MockUserStorageMockRecorder) CreateTx(ctx interface{}, options ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, options...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTx", reflect.TypeOf((*MockUserStorage)(nil).CreateTx), varargs...)
}

// GetUserToken mocks base method.
func (m *MockUserStorage) GetUserToken(ctx context.Context, tx storage.Tx, token string) (auth.UserToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserToken", ctx, tx, token)
	ret0, _ := ret[0].(auth.UserToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserToken indicates an expected call of GetUserToken.
func (mr *MockUserStorageMockRecorder) GetUserToken(ctx, tx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserToken", reflect.TypeOf((*MockUserStorage)(nil).GetUserToken), ctx, tx, token)
}

// ListUsers mocks base method.
func (m *MockUserStorage) ListUsers(ctx context.Context, tx storage.Tx, req auth.ListUserRequest) (auth.ListUserResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, tx, req)
	ret0, _ := ret[0].(auth.ListUserResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockUserStorageMockRecorder) ListUsers(ctx, tx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserStorage)(nil).ListUsers), ctx, tx, req)
}

// StoreUser mocks base method.
func (m *MockUserStorage) StoreUser(ctx context.Context, tx storage.Tx, user auth.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUser", ctx, tx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUser indicates an expected call of StoreUser.
func (mr *MockUserStorageMockRecorder) StoreUser(ctx, tx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUser", reflect.TypeOf((*MockUserStorage)(nil).StoreUser), ctx, tx, user)
}

// StoreUserToken mocks base method.
func (m *MockUserStorage) StoreUserToken(ctx context.Context, tx storage.Tx, token auth.UserToken) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreUserToken", ctx, tx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreUserToken indicates an expected call of StoreUserToken.
func (mr *MockUserStorageMockRecorder) StoreUserToken(ctx, tx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreUserToken", reflect.TypeOf((*MockUserStorage)(nil).StoreUserToken), ctx, tx, token)
}
