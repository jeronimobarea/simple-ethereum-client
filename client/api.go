package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Api interface {
	SimpleSendTransaction(quantity *big.Int, fromPk *ecdsa.PrivateKey, to, token common.Address) (*TransactionResponse, error)
	SimpleCheckBalance(address, token common.Address) (*BalanceResponse, error)
	SimpleCheckBalances(addresses []common.Address, token common.Address) (*BalancesResponse, error)
}
