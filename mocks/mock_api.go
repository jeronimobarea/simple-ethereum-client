// Code generated by MockGen. DO NOT EDIT.
// Source: ../client/api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	ecdsa "crypto/ecdsa"
	big "math/big"
	reflect "reflect"

	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	gomock "github.com/golang/mock/gomock"
	client "github.com/jeronimobarea/simple-ethereum-client/client"
)

// MockApi is a mock of Api interface.
type MockApi struct {
	ctrl     *gomock.Controller
	recorder *MockApiMockRecorder
}

// MockApiMockRecorder is the mock recorder for MockApi.
type MockApiMockRecorder struct {
	mock *MockApi
}

// NewMockApi creates a new mock instance.
func NewMockApi(ctrl *gomock.Controller) *MockApi {
	mock := &MockApi{ctrl: ctrl}
	mock.recorder = &MockApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApi) EXPECT() *MockApiMockRecorder {
	return m.recorder
}

// CheckBalance mocks base method.
func (m *MockApi) CheckBalance(address, token common.Address) (*client.BalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBalance", address, token)
	ret0, _ := ret[0].(*client.BalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBalance indicates an expected call of CheckBalance.
func (mr *MockApiMockRecorder) CheckBalance(address, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBalance", reflect.TypeOf((*MockApi)(nil).CheckBalance), address, token)
}

// CheckBalances mocks base method.
func (m *MockApi) CheckBalances(addresses []common.Address, token common.Address) (*client.BalancesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckBalances", addresses, token)
	ret0, _ := ret[0].(*client.BalancesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckBalances indicates an expected call of CheckBalances.
func (mr *MockApiMockRecorder) CheckBalances(addresses, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckBalances", reflect.TypeOf((*MockApi)(nil).CheckBalances), addresses, token)
}

// SendTransaction mocks base method.
func (m *MockApi) SendTransaction(quantity *big.Int, fromPk *ecdsa.PrivateKey, to, token common.Address) (*types.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransaction", quantity, fromPk, to, token)
	ret0, _ := ret[0].(*types.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTransaction indicates an expected call of SendTransaction.
func (mr *MockApiMockRecorder) SendTransaction(quantity, fromPk, to, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransaction", reflect.TypeOf((*MockApi)(nil).SendTransaction), quantity, fromPk, to, token)
}
