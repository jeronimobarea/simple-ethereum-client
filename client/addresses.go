package client

import (
	"fmt"
	"math/big"

	"github.com/apex/log"
	"github.com/ethereum/go-ethereum/common"
	validator "github.com/jeronimobarea/simple-ethereum/address"
)

type (
	BalanceResponse struct {
		Balance *big.Int `json:"balance"`
		Error   error    `json:"error"`
	}

	BalancesResponse struct {
		Processed map[common.Address]*BalanceResponse `json:"processed"`
	}
)

func (svc *ethereumService) CheckBalance(
	address,
	token common.Address,
) (resp *BalanceResponse, err error) {
	if valid := validator.ValidateAddress(address); !valid {
		err = fmt.Errorf("invalid addresss: %s", address)
		log.
			WithField("address", address).
			WithError(err)
		return
	}
	if valid := validator.ValidateAddress(token); !valid {
		err = fmt.Errorf("invalid token addresss: %s", address)
		log.
			WithField("token_address", token).
			WithError(err)
		return
	}
	return svc.API.CheckBalance(address, token)
}

func (svc *ethereumService) CheckBalances(
	addresses []common.Address,
	token common.Address,
) (resp *BalancesResponse, err error) {
	return svc.API.CheckBalances(addresses, token)
}
