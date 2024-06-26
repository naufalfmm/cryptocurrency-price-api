// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package userCoins is a generated GoMock package.
package userCoins

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	dto "github.com/naufalfmm/cryptocurrency-price-api/model/dto"
	sqliteOrm "github.com/naufalfmm/cryptocurrency-price-api/utils/orm/driver/sqliteOrm"
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
func (m *MockRepositories) Create(ctx context.Context, userCoin dao.UserCoin) (dao.UserCoin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userCoin)
	ret0, _ := ret[0].(dao.UserCoin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoriesMockRecorder) Create(ctx, userCoin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositories)(nil).Create), ctx, userCoin)
}

// DeleteByID mocks base method.
func (m *MockRepositories) DeleteByID(ctx context.Context, id uint64, deletedBy string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByID", ctx, id, deletedBy)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByID indicates an expected call of DeleteByID.
func (mr *MockRepositoriesMockRecorder) DeleteByID(ctx, id, deletedBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByID", reflect.TypeOf((*MockRepositories)(nil).DeleteByID), ctx, id, deletedBy)
}

// Get mocks base method.
func (m *MockRepositories) Get(ctx context.Context, userID, coinID uint64) (dao.UserCoin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userID, coinID)
	ret0, _ := ret[0].(dao.UserCoin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRepositoriesMockRecorder) Get(ctx, userID, coinID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepositories)(nil).Get), ctx, userID, coinID)
}

// GetAll mocks base method.
func (m *MockRepositories) GetAll(ctx context.Context, req dto.GetAllRequest, queryModifier sqliteOrm.QueryModifier) ([]dao.UserCoin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, req, queryModifier)
	ret0, _ := ret[0].([]dao.UserCoin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockRepositoriesMockRecorder) GetAll(ctx, req, queryModifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockRepositories)(nil).GetAll), ctx, req, queryModifier)
}
