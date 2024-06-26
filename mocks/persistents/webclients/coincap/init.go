// Code generated by MockGen. DO NOT EDIT.
// Source: init.go

// Package coincap is a generated GoMock package.
package coincap

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dao "github.com/naufalfmm/cryptocurrency-price-api/model/dao"
	dto "github.com/naufalfmm/cryptocurrency-price-api/model/dto"
)

// MockCoincap is a mock of Coincap interface.
type MockCoincap struct {
	ctrl     *gomock.Controller
	recorder *MockCoincapMockRecorder
}

// MockCoincapMockRecorder is the mock recorder for MockCoincap.
type MockCoincapMockRecorder struct {
	mock *MockCoincap
}

// NewMockCoincap creates a new mock instance.
func NewMockCoincap(ctrl *gomock.Controller) *MockCoincap {
	mock := &MockCoincap{ctrl: ctrl}
	mock.recorder = &MockCoincapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCoincap) EXPECT() *MockCoincapMockRecorder {
	return m.recorder
}

// GetAllAssets mocks base method.
func (m *MockCoincap) GetAllAssets(ctx context.Context, req dto.AllAssetsCoincapRequest) (dao.AllAsset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAssets", ctx, req)
	ret0, _ := ret[0].(dao.AllAsset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAssets indicates an expected call of GetAllAssets.
func (mr *MockCoincapMockRecorder) GetAllAssets(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAssets", reflect.TypeOf((*MockCoincap)(nil).GetAllAssets), ctx, req)
}

// GetAllRates mocks base method.
func (m *MockCoincap) GetAllRates(ctx context.Context) (dao.GetAllRates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRates", ctx)
	ret0, _ := ret[0].(dao.GetAllRates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRates indicates an expected call of GetAllRates.
func (mr *MockCoincapMockRecorder) GetAllRates(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRates", reflect.TypeOf((*MockCoincap)(nil).GetAllRates), ctx)
}
