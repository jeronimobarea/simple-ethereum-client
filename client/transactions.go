package client

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/apex/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	simpleAddress "github.com/jeronimobarea/simple-ethereum/address"
)

type TransactionResponse struct {
	Transaction *types.Transaction `json:"transaction"`
}

func (svc *ethereumService) SendTransaction(
	quantity *big.Int,
	fromPk *ecdsa.PrivateKey,
	to,
	token common.Address,
) (resp *TransactionResponse, err error) {
	addresses := []interface{}{to, token}

	if _, invalid, _ := simpleAddress.ValidateAddresses(addresses); invalid != nil {
		err = fmt.Errorf("invalid address/es: %s", invalid)
		log.WithError(err)
		return
	}

	if valid := simpleAddress.SafeBalanceIsValid(quantity); !valid {
		err = errors.New("quantity is not valid")
		log.
			WithField("quantity", quantity).
			WithError(err)
		return
	}
	return svc.API.SendTransaction(quantity, fromPk, to, token)
}
