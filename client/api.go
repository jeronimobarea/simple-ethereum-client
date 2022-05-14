package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate mockgen -source=../client/api.go -destination=../mocks/mock_api.go -package=mocks
type Api interface {
	SendTransaction(quantity *big.Int, fromPk *ecdsa.PrivateKey, to, token common.Address) (*types.Transaction, error)
	CheckBalance(address, token common.Address) (*BalanceResponse, error)
	CheckBalances(addresses []common.Address, token common.Address) (*BalancesResponse, error)
}
