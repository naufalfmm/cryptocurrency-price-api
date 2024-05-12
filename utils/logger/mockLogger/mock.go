// Code generated by MockGen. DO NOT EDIT.
// Source: logger.go

// Package mockLogger is a generated GoMock package.
package mockLogger

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	logger "github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
)

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockLogger) Any(key string, val any) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Any", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockLoggerMockRecorder) Any(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockLogger)(nil).Any), key, val)
}

// Bool mocks base method.
func (m *MockLogger) Bool(key string, val bool) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bool", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Bool indicates an expected call of Bool.
func (mr *MockLoggerMockRecorder) Bool(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bool", reflect.TypeOf((*MockLogger)(nil).Bool), key, val)
}

// Debug mocks base method.
func (m *MockLogger) Debug(ctx context.Context, msg string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug", ctx, msg)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Debug indicates an expected call of Debug.
func (mr *MockLoggerMockRecorder) Debug(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), ctx, msg)
}

// Dur mocks base method.
func (m *MockLogger) Dur(key string, val time.Duration) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dur", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Dur indicates an expected call of Dur.
func (mr *MockLoggerMockRecorder) Dur(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dur", reflect.TypeOf((*MockLogger)(nil).Dur), key, val)
}

// Err mocks base method.
func (m *MockLogger) Err(err error) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err", err)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockLoggerMockRecorder) Err(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockLogger)(nil).Err), err)
}

// Error mocks base method.
func (m *MockLogger) Error(ctx context.Context, msg string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error", ctx, msg)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockLoggerMockRecorder) Error(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockLogger)(nil).Error), ctx, msg)
}

// Float32 mocks base method.
func (m *MockLogger) Float32(key string, val float32) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float32", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Float32 indicates an expected call of Float32.
func (mr *MockLoggerMockRecorder) Float32(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float32", reflect.TypeOf((*MockLogger)(nil).Float32), key, val)
}

// Float64 mocks base method.
func (m *MockLogger) Float64(key string, val float64) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Float64", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Float64 indicates an expected call of Float64.
func (mr *MockLoggerMockRecorder) Float64(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Float64", reflect.TypeOf((*MockLogger)(nil).Float64), key, val)
}

// Info mocks base method.
func (m *MockLogger) Info(ctx context.Context, msg string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", ctx, msg)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Info indicates an expected call of Info.
func (mr *MockLoggerMockRecorder) Info(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockLogger)(nil).Info), ctx, msg)
}

// Int mocks base method.
func (m *MockLogger) Int(key string, val int) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Int indicates an expected call of Int.
func (mr *MockLoggerMockRecorder) Int(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int", reflect.TypeOf((*MockLogger)(nil).Int), key, val)
}

// Int16 mocks base method.
func (m *MockLogger) Int16(key string, val int16) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int16", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Int16 indicates an expected call of Int16.
func (mr *MockLoggerMockRecorder) Int16(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int16", reflect.TypeOf((*MockLogger)(nil).Int16), key, val)
}

// Int32 mocks base method.
func (m *MockLogger) Int32(key string, val int32) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int32", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Int32 indicates an expected call of Int32.
func (mr *MockLoggerMockRecorder) Int32(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int32", reflect.TypeOf((*MockLogger)(nil).Int32), key, val)
}

// Int64 mocks base method.
func (m *MockLogger) Int64(key string, val int64) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int64", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Int64 indicates an expected call of Int64.
func (mr *MockLoggerMockRecorder) Int64(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int64", reflect.TypeOf((*MockLogger)(nil).Int64), key, val)
}

// Int8 mocks base method.
func (m *MockLogger) Int8(key string, val int8) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Int8", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Int8 indicates an expected call of Int8.
func (mr *MockLoggerMockRecorder) Int8(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Int8", reflect.TypeOf((*MockLogger)(nil).Int8), key, val)
}

// Printf mocks base method.
func (m *MockLogger) Printf(arg0 string, arg1 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Printf", varargs...)
}

// Printf indicates an expected call of Printf.
func (mr *MockLoggerMockRecorder) Printf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Printf", reflect.TypeOf((*MockLogger)(nil).Printf), varargs...)
}

// Send mocks base method.
func (m *MockLogger) Send() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send")
}

// Send indicates an expected call of Send.
func (mr *MockLoggerMockRecorder) Send() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockLogger)(nil).Send))
}

// Str mocks base method.
func (m *MockLogger) Str(key, val string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Str", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Str indicates an expected call of Str.
func (mr *MockLoggerMockRecorder) Str(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Str", reflect.TypeOf((*MockLogger)(nil).Str), key, val)
}

// Time mocks base method.
func (m *MockLogger) Time(key string, t time.Time) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Time", key, t)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Time indicates an expected call of Time.
func (mr *MockLoggerMockRecorder) Time(key, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Time", reflect.TypeOf((*MockLogger)(nil).Time), key, t)
}

// Uint mocks base method.
func (m *MockLogger) Uint(key string, val uint) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Uint indicates an expected call of Uint.
func (mr *MockLoggerMockRecorder) Uint(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint", reflect.TypeOf((*MockLogger)(nil).Uint), key, val)
}

// Uint16 mocks base method.
func (m *MockLogger) Uint16(key string, val uint16) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint16", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Uint16 indicates an expected call of Uint16.
func (mr *MockLoggerMockRecorder) Uint16(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint16", reflect.TypeOf((*MockLogger)(nil).Uint16), key, val)
}

// Uint32 mocks base method.
func (m *MockLogger) Uint32(key string, val uint32) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint32", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Uint32 indicates an expected call of Uint32.
func (mr *MockLoggerMockRecorder) Uint32(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint32", reflect.TypeOf((*MockLogger)(nil).Uint32), key, val)
}

// Uint64 mocks base method.
func (m *MockLogger) Uint64(key string, val uint64) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint64", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Uint64 indicates an expected call of Uint64.
func (mr *MockLoggerMockRecorder) Uint64(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint64", reflect.TypeOf((*MockLogger)(nil).Uint64), key, val)
}

// Uint8 mocks base method.
func (m *MockLogger) Uint8(key string, val uint8) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Uint8", key, val)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Uint8 indicates an expected call of Uint8.
func (mr *MockLoggerMockRecorder) Uint8(key, val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Uint8", reflect.TypeOf((*MockLogger)(nil).Uint8), key, val)
}

// Warn mocks base method.
func (m *MockLogger) Warn(ctx context.Context, msg string) logger.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Warn", ctx, msg)
	ret0, _ := ret[0].(logger.Logger)
	return ret0
}

// Warn indicates an expected call of Warn.
func (mr *MockLoggerMockRecorder) Warn(ctx, msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warn", reflect.TypeOf((*MockLogger)(nil).Warn), ctx, msg)
}
