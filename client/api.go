package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Api interface {
	SendTransaction(quantity *big.Int, fromPk *ecdsa.PrivateKey, to, token common.Address) (*TransactionResponse, error)
	CheckBalance(address, token common.Address) (*BalanceResponse, error)
	CheckBalances(addresses []common.Address, token common.Address) (*BalancesResponse, error)
}
