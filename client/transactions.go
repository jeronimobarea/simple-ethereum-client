package client

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
)

func (svc *ethereumService) SendTransaction(
	quantity *big.Int,
	fromPk *ecdsa.PrivateKey,
	to,
	token common.Address,
) (resp *types.Transaction, err error) {
	addresses := []interface{}{to, token}

	if _, invalid, _ := simpleCommon.ValidateAddresses(addresses); invalid != nil {
		err = fmt.Errorf("invalid address/es: %s", invalid)
		return
	}

	if valid := simpleCommon.SafeBalanceIsValid(quantity); !valid {
		err = fmt.Errorf("quantity is not valid %d", quantity)
		return
	}

	if fromPk == nil {
		err = errors.New("private key cannot be nil")
		return
	}
	return svc.API.SendTransaction(quantity, fromPk, to, token)
}
