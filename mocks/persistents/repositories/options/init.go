// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package options is a generated GoMock package.
package options

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/naufalfmm/cryptocurrency-price-api/model/dao"
)

// MockRepositories is a mock of Repositories interface.
type MockRepositories struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesMockRecorder
}

// MockRepositoriesMockRecorder is the mock recorder for MockRepositories.
type MockRepositoriesMockRecorder struct {
	mock *MockRepositories
}

// NewMockRepositories creates a new mock instance.
func NewMockRepositories(ctrl *gomock.Controller) *MockRepositories {
	mock := &MockRepositories{ctrl: ctrl}
	mock.recorder = &MockRepositoriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositories) EXPECT() *MockRepositoriesMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRepositories) Create(ctx context.Context, opt dao.Option) (dao.Option, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, opt)
	ret0, _ := ret[0].(dao.Option)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoriesMockRecorder) Create(ctx, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositories)(nil).Create), ctx, opt)
}

// GetByKey mocks base method.
func (m *MockRepositories) GetByKey(ctx context.Context, key string) (dao.Option, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByKey", ctx, key)
	ret0, _ := ret[0].(dao.Option)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByKey indicates an expected call of GetByKey.
func (mr *MockRepositoriesMockRecorder) GetByKey(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByKey", reflect.TypeOf((*MockRepositories)(nil).GetByKey), ctx, key)
}

// UpdateByKey mocks base method.
func (m *MockRepositories) UpdateByKey(ctx context.Context, key, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateByKey", ctx, key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateByKey indicates an expected call of UpdateByKey.
func (mr *MockRepositoriesMockRecorder) UpdateByKey(ctx, key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateByKey", reflect.TypeOf((*MockRepositories)(nil).UpdateByKey), ctx, key, value)
}
