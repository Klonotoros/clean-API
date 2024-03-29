// Code generated by MockGen. DO NOT EDIT.
// Source: user.go
//
// Generated by this command:
//
//	mockgen -source=user.go -destination=user_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	model "clean-API/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// FindByID mocks base method.
func (m *MockUserRepository) FindByID(arg0 int64) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockUserRepositoryMockRecorder) FindByID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockUserRepository)(nil).FindByID), arg0)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(arg0 string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), arg0)
}

// Save mocks base method.
func (m *MockUserRepository) Save(arg0 model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockUserRepositoryMockRecorder) Save(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepository)(nil).Save), arg0)
}
