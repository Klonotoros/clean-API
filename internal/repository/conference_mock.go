// Code generated by MockGen. DO NOT EDIT.
// Source: conference.go
//
// Generated by this command:
//
//	mockgen -source=conference.go -destination=conference_mock.go -package repository
//

// Package repository is a generated GoMock package.
package repository

import (
	model "clean-API/internal/model"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockConferenceRepository is a mock of ConferenceRepository interface.
type MockConferenceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockConferenceRepositoryMockRecorder
}

// MockConferenceRepositoryMockRecorder is the mock recorder for MockConferenceRepository.
type MockConferenceRepositoryMockRecorder struct {
	mock *MockConferenceRepository
}

// NewMockConferenceRepository creates a new mock instance.
func NewMockConferenceRepository(ctrl *gomock.Controller) *MockConferenceRepository {
	mock := &MockConferenceRepository{ctrl: ctrl}
	mock.recorder = &MockConferenceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConferenceRepository) EXPECT() *MockConferenceRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockConferenceRepository) Delete(conference model.Conference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", conference)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockConferenceRepositoryMockRecorder) Delete(conference any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockConferenceRepository)(nil).Delete), conference)
}

// GetAllConferences mocks base method.
func (m *MockConferenceRepository) GetAllConferences() ([]model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllConferences")
	ret0, _ := ret[0].([]model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllConferences indicates an expected call of GetAllConferences.
func (mr *MockConferenceRepositoryMockRecorder) GetAllConferences() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllConferences", reflect.TypeOf((*MockConferenceRepository)(nil).GetAllConferences))
}

// GetConferenceByID mocks base method.
func (m *MockConferenceRepository) GetConferenceByID(id int64) (model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConferenceByID", id)
	ret0, _ := ret[0].(model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConferenceByID indicates an expected call of GetConferenceByID.
func (mr *MockConferenceRepositoryMockRecorder) GetConferenceByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConferenceByID", reflect.TypeOf((*MockConferenceRepository)(nil).GetConferenceByID), id)
}

// Save mocks base method.
func (m *MockConferenceRepository) Save(conference model.Conference) (model.Conference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", conference)
	ret0, _ := ret[0].(model.Conference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockConferenceRepositoryMockRecorder) Save(conference any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockConferenceRepository)(nil).Save), conference)
}

// Update mocks base method.
func (m *MockConferenceRepository) Update(conference model.Conference) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", conference)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockConferenceRepositoryMockRecorder) Update(conference any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockConferenceRepository)(nil).Update), conference)
}
