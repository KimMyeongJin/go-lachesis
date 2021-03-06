// Code generated by MockGen. DO NOT EDIT.
// Source: consensus.go

// Package posnode is a generated GoMock package.
package posnode

import (
	hash "github.com/Fantom-foundation/go-lachesis/src/hash"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockConsensus is a mock of Consensus interface
type MockConsensus struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusMockRecorder
}

// MockConsensusMockRecorder is the mock recorder for MockConsensus
type MockConsensusMockRecorder struct {
	mock *MockConsensus
}

// NewMockConsensus creates a new mock instance
func NewMockConsensus(ctrl *gomock.Controller) *MockConsensus {
	mock := &MockConsensus{ctrl: ctrl}
	mock.recorder = &MockConsensusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConsensus) EXPECT() *MockConsensusMockRecorder {
	return m.recorder
}

// PushEvent mocks base method
func (m *MockConsensus) PushEvent(arg0 hash.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushEvent", arg0)
}

// PushEvent indicates an expected call of PushEvent
func (mr *MockConsensusMockRecorder) PushEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushEvent", reflect.TypeOf((*MockConsensus)(nil).PushEvent), arg0)
}

// GetStakeOf mocks base method
func (m *MockConsensus) GetStakeOf(arg0 hash.Peer) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStakeOf", arg0)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetStakeOf indicates an expected call of GetStakeOf
func (mr *MockConsensusMockRecorder) GetStakeOf(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStakeOf", reflect.TypeOf((*MockConsensus)(nil).GetStakeOf), arg0)
}
