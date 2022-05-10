package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
)

type MockApi struct {
	mock.Mock
}

func (api *MockApi) SendTransaction(
	quantity *big.Int,
	fromPk *ecdsa.PrivateKey,
	to,
	token common.Address,
) (
	resp *types.Transaction,
	err error,
) {
	args := api.Called(quantity, fromPk, to, token)
	if args.Get(0) != nil {
		resp = args.Get(0).(*types.Transaction)
	}
	return resp, args.Error(1)
}

func (api *MockApi) CheckBalance(
	address, token common.Address,
) (
	resp *BalanceResponse,
	err error,
) {
	args := api.Called(address, token)
	if args.Get(0) != nil {
		resp = args.Get(0).(*BalanceResponse)
	}
	return resp, args.Error(1)
}

func (api *MockApi) CheckBalances(
	address []common.Address, token common.Address,
) (
	resp *BalancesResponse,
	err error,
) {
	args := api.Called(address, token)
	if args.Get(0) != nil {
		resp = args.Get(0).(*BalancesResponse)
	}
	return resp, args.Error(1)
}
