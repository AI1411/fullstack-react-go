// Code generated by MockGen. DO NOT EDIT.
// Source: notification.go
//
// Generated by this command:
//
//	mockgen -source=notification.go -destination=../../../tests/mock/domain/notification.mock.go
//

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	reflect "reflect"

	model "github.com/AI1411/fullstack-react-go/internal/domain/model"
	gomock "go.uber.org/mock/gomock"
)

// MockNotificationRepository is a mock of NotificationRepository interface.
type MockNotificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNotificationRepositoryMockRecorder
	isgomock struct{}
}

// MockNotificationRepositoryMockRecorder is the mock recorder for MockNotificationRepository.
type MockNotificationRepositoryMockRecorder struct {
	mock *MockNotificationRepository
}

// NewMockNotificationRepository creates a new mock instance.
func NewMockNotificationRepository(ctrl *gomock.Controller) *MockNotificationRepository {
	mock := &MockNotificationRepository{ctrl: ctrl}
	mock.recorder = &MockNotificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotificationRepository) EXPECT() *MockNotificationRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockNotificationRepository) Create(ctx context.Context, notification *model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockNotificationRepositoryMockRecorder) Create(ctx, notification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNotificationRepository)(nil).Create), ctx, notification)
}

// Delete mocks base method.
func (m *MockNotificationRepository) Delete(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNotificationRepositoryMockRecorder) Delete(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNotificationRepository)(nil).Delete), ctx, id)
}

// Find mocks base method.
func (m *MockNotificationRepository) Find(ctx context.Context) ([]*model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx)
	ret0, _ := ret[0].([]*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockNotificationRepositoryMockRecorder) Find(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockNotificationRepository)(nil).Find), ctx)
}

// FindByID mocks base method.
func (m *MockNotificationRepository) FindByID(ctx context.Context, id int32) (*model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockNotificationRepositoryMockRecorder) FindByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockNotificationRepository)(nil).FindByID), ctx, id)
}

// FindByUserID mocks base method.
func (m *MockNotificationRepository) FindByUserID(ctx context.Context, userID int32) ([]*model.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUserID", ctx, userID)
	ret0, _ := ret[0].([]*model.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUserID indicates an expected call of FindByUserID.
func (mr *MockNotificationRepositoryMockRecorder) FindByUserID(ctx, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUserID", reflect.TypeOf((*MockNotificationRepository)(nil).FindByUserID), ctx, userID)
}

// MarkAsRead mocks base method.
func (m *MockNotificationRepository) MarkAsRead(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkAsRead", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkAsRead indicates an expected call of MarkAsRead.
func (mr *MockNotificationRepositoryMockRecorder) MarkAsRead(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkAsRead", reflect.TypeOf((*MockNotificationRepository)(nil).MarkAsRead), ctx, id)
}

// Update mocks base method.
func (m *MockNotificationRepository) Update(ctx context.Context, notification *model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNotificationRepositoryMockRecorder) Update(ctx, notification any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNotificationRepository)(nil).Update), ctx, notification)
}
