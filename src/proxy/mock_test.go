// Code generated by MockGen. DO NOT EDIT.
// Source: ./ctrl_server.go

// Package proxy is a generated GoMock package.
package proxy

import (
	hash "github.com/Fantom-foundation/go-lachesis/src/hash"
	inter "github.com/Fantom-foundation/go-lachesis/src/inter"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNode is a mock of Node interface
type MockNode struct {
	ctrl     *gomock.Controller
	recorder *MockNodeMockRecorder
}

// MockNodeMockRecorder is the mock recorder for MockNode
type MockNodeMockRecorder struct {
	mock *MockNode
}

// NewMockNode creates a new mock instance
func NewMockNode(ctrl *gomock.Controller) *MockNode {
	mock := &MockNode{ctrl: ctrl}
	mock.recorder = &MockNodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNode) EXPECT() *MockNodeMockRecorder {
	return m.recorder
}

// GetID mocks base method
func (m *MockNode) GetID() hash.Peer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(hash.Peer)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockNodeMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockNode)(nil).GetID))
}

// AddInternalTxn mocks base method
func (m *MockNode) AddInternalTxn(arg0 inter.InternalTransaction) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddInternalTxn", arg0)
}

// AddInternalTxn indicates an expected call of AddInternalTxn
func (mr *MockNodeMockRecorder) AddInternalTxn(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInternalTxn", reflect.TypeOf((*MockNode)(nil).AddInternalTxn), arg0)
}

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

// GetStakeOf mocks base method
func (m *MockConsensus) GetStakeOf(peer hash.Peer) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStakeOf", peer)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetStakeOf indicates an expected call of GetStakeOf
func (mr *MockConsensusMockRecorder) GetStakeOf(peer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStakeOf", reflect.TypeOf((*MockConsensus)(nil).GetStakeOf), peer)
}
