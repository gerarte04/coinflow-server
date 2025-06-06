// Code generated by MockGen. DO NOT EDIT.
// Source: users.go
//
// Generated by this command:
//
//	mockgen -source users.go -destination mocks/users_mock.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	models "coinflow/coinflow-server/storage-service/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUsersRepo is a mock of UsersRepo interface.
type MockUsersRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUsersRepoMockRecorder
	isgomock struct{}
}

// MockUsersRepoMockRecorder is the mock recorder for MockUsersRepo.
type MockUsersRepoMockRecorder struct {
	mock *MockUsersRepo
}

// NewMockUsersRepo creates a new mock instance.
func NewMockUsersRepo(ctrl *gomock.Controller) *MockUsersRepo {
	mock := &MockUsersRepo{ctrl: ctrl}
	mock.recorder = &MockUsersRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersRepo) EXPECT() *MockUsersRepoMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockUsersRepo) GetUser(usrId string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", usrId)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUsersRepoMockRecorder) GetUser(usrId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsersRepo)(nil).GetUser), usrId)
}

// GetUserByCred mocks base method.
func (m *MockUsersRepo) GetUserByCred(login, password string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByCred", login, password)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByCred indicates an expected call of GetUserByCred.
func (mr *MockUsersRepoMockRecorder) GetUserByCred(login, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByCred", reflect.TypeOf((*MockUsersRepo)(nil).GetUserByCred), login, password)
}

// PostUser mocks base method.
func (m *MockUsersRepo) PostUser(usr *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostUser", usr)
	ret0, _ := ret[0].(error)
	return ret0
}

// PostUser indicates an expected call of PostUser.
func (mr *MockUsersRepoMockRecorder) PostUser(usr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostUser", reflect.TypeOf((*MockUsersRepo)(nil).PostUser), usr)
}
