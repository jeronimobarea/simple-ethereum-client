package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Api interface {
	SendTransaction(quantity *big.Int, fromPk *ecdsa.PrivateKey, to, token common.Address) (*types.Transaction, error)
	CheckBalance(address, token common.Address) (*BalanceResponse, error)
	CheckBalances(addresses []common.Address, token common.Address) (*BalancesResponse, error)
}
