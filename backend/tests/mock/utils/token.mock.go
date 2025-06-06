// Code generated by MockGen. DO NOT EDIT.
// Source: token.go
//
// Generated by this command:
//
//	mockgen -source=token.go -destination=../../tests/mock/utils/token.mock.go
//

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTokenGenerator is a mock of TokenGenerator interface.
type MockTokenGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockTokenGeneratorMockRecorder
	isgomock struct{}
}

// MockTokenGeneratorMockRecorder is the mock recorder for MockTokenGenerator.
type MockTokenGeneratorMockRecorder struct {
	mock *MockTokenGenerator
}

// NewMockTokenGenerator creates a new mock instance.
func NewMockTokenGenerator(ctrl *gomock.Controller) *MockTokenGenerator {
	mock := &MockTokenGenerator{ctrl: ctrl}
	mock.recorder = &MockTokenGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenGenerator) EXPECT() *MockTokenGeneratorMockRecorder {
	return m.recorder
}

// GenerateEmailVerificationToken mocks base method.
func (m *MockTokenGenerator) GenerateEmailVerificationToken() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateEmailVerificationToken")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateEmailVerificationToken indicates an expected call of GenerateEmailVerificationToken.
func (mr *MockTokenGeneratorMockRecorder) GenerateEmailVerificationToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateEmailVerificationToken", reflect.TypeOf((*MockTokenGenerator)(nil).GenerateEmailVerificationToken))
}
