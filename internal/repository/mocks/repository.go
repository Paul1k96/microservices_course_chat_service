// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go
//
// Generated by this command:
//
//	mockgen -source repository.go -destination mocks/repository.go -package mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/Paul1k96/microservices_course_chat_service/internal/model"
	gomock "go.uber.org/mock/gomock"
)

// MockChatRepository is a mock of ChatRepository interface.
type MockChatRepository struct {
	ctrl     *gomock.Controller
	recorder *MockChatRepositoryMockRecorder
	isgomock struct{}
}

// MockChatRepositoryMockRecorder is the mock recorder for MockChatRepository.
type MockChatRepositoryMockRecorder struct {
	mock *MockChatRepository
}

// NewMockChatRepository creates a new mock instance.
func NewMockChatRepository(ctrl *gomock.Controller) *MockChatRepository {
	mock := &MockChatRepository{ctrl: ctrl}
	mock.recorder = &MockChatRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChatRepository) EXPECT() *MockChatRepositoryMockRecorder {
	return m.recorder
}

// AddUsers mocks base method.
func (m *MockChatRepository) AddUsers(ctx context.Context, users model.Users) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUsers", ctx, users)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUsers indicates an expected call of AddUsers.
func (mr *MockChatRepositoryMockRecorder) AddUsers(ctx, users any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUsers", reflect.TypeOf((*MockChatRepository)(nil).AddUsers), ctx, users)
}

// Create mocks base method.
func (m *MockChatRepository) Create(ctx context.Context) (model.ChatID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx)
	ret0, _ := ret[0].(model.ChatID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockChatRepositoryMockRecorder) Create(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockChatRepository)(nil).Create), ctx)
}

// Delete mocks base method.
func (m *MockChatRepository) Delete(ctx context.Context, chatID model.ChatID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, chatID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockChatRepositoryMockRecorder) Delete(ctx, chatID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockChatRepository)(nil).Delete), ctx, chatID)
}

// MockMessageRepository is a mock of MessageRepository interface.
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
	isgomock struct{}
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository.
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance.
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockMessageRepository) Create(ctx context.Context, message *model.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockMessageRepositoryMockRecorder) Create(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageRepository)(nil).Create), ctx, message)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
	isgomock struct{}
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

// Get mocks base method.
func (m *MockUserRepository) Get(ctx context.Context, id model.UserID) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserRepositoryMockRecorder) Get(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRepository)(nil).Get), ctx, id)
}

// List mocks base method.
func (m *MockUserRepository) List(ctx context.Context, ids model.UserIDs) (model.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, ids)
	ret0, _ := ret[0].(model.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserRepositoryMockRecorder) List(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserRepository)(nil).List), ctx, ids)
}