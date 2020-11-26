// Code generated by MockGen. DO NOT EDIT.
// Source: structure.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockStacker is a mock of Stacker interface
type MockStacker struct {
	ctrl     *gomock.Controller
	recorder *MockStackerMockRecorder
}

// MockStackerMockRecorder is the mock recorder for MockStacker
type MockStackerMockRecorder struct {
	mock *MockStacker
}

// NewMockStacker creates a new mock instance
func NewMockStacker(ctrl *gomock.Controller) *MockStacker {
	mock := &MockStacker{ctrl: ctrl}
	mock.recorder = &MockStackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStacker) EXPECT() *MockStackerMockRecorder {
	return m.recorder
}

// Append mocks base method
func (m *MockStacker) Append(arg0 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Append", arg0)
}

// Append indicates an expected call of Append
func (mr *MockStackerMockRecorder) Append(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockStacker)(nil).Append), arg0)
}

// Pop mocks base method
func (m *MockStacker) Pop() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pop")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Pop indicates an expected call of Pop
func (mr *MockStackerMockRecorder) Pop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pop", reflect.TypeOf((*MockStacker)(nil).Pop))
}

// MockQueuer is a mock of Queuer interface
type MockQueuer struct {
	ctrl     *gomock.Controller
	recorder *MockQueuerMockRecorder
}

// MockQueuerMockRecorder is the mock recorder for MockQueuer
type MockQueuerMockRecorder struct {
	mock *MockQueuer
}

// NewMockQueuer creates a new mock instance
func NewMockQueuer(ctrl *gomock.Controller) *MockQueuer {
	mock := &MockQueuer{ctrl: ctrl}
	mock.recorder = &MockQueuerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueuer) EXPECT() *MockQueuerMockRecorder {
	return m.recorder
}

// Enque mocks base method
func (m *MockQueuer) Enque(arg0 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Enque", arg0)
}

// Enque indicates an expected call of Enque
func (mr *MockQueuerMockRecorder) Enque(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enque", reflect.TypeOf((*MockQueuer)(nil).Enque), arg0)
}

// Dequeue mocks base method
func (m *MockQueuer) Dequeue() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dequeue")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Dequeue indicates an expected call of Dequeue
func (mr *MockQueuerMockRecorder) Dequeue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dequeue", reflect.TypeOf((*MockQueuer)(nil).Dequeue))
}