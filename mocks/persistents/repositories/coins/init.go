// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package coins is a generated GoMock package.
package coins

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
func (m *MockRepositories) Create(ctx context.Context, coin dao.Coin) (dao.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, coin)
	ret0, _ := ret[0].(dao.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoriesMockRecorder) Create(ctx, coin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositories)(nil).Create), ctx, coin)
}

// GetByCode mocks base method.
func (m *MockRepositories) GetByCode(ctx context.Context, code string) (dao.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCode", ctx, code)
	ret0, _ := ret[0].(dao.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCode indicates an expected call of GetByCode.
func (mr *MockRepositoriesMockRecorder) GetByCode(ctx, code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCode", reflect.TypeOf((*MockRepositories)(nil).GetByCode), ctx, code)
}

// GetByCoincapIDs mocks base method.
func (m *MockRepositories) GetByCoincapIDs(ctx context.Context, coincapIDs []string) ([]dao.Coin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCoincapIDs", ctx, coincapIDs)
	ret0, _ := ret[0].([]dao.Coin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCoincapIDs indicates an expected call of GetByCoincapIDs.
func (mr *MockRepositoriesMockRecorder) GetByCoincapIDs(ctx, coincapIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCoincapIDs", reflect.TypeOf((*MockRepositories)(nil).GetByCoincapIDs), ctx, coincapIDs)
}

// UpdatePrices mocks base method.
func (m *MockRepositories) UpdatePrices(ctx context.Context, coins []dao.Coin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePrices", ctx, coins)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePrices indicates an expected call of UpdatePrices.
func (mr *MockRepositoriesMockRecorder) UpdatePrices(ctx, coins interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePrices", reflect.TypeOf((*MockRepositories)(nil).UpdatePrices), ctx, coins)
}
