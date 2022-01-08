package client

import (
	"fmt"
	"math/big"

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
	addresses := []interface{}{address, token}
	if _, invalid, _ := validator.ValidateAddresses(addresses); invalid != nil {
		err = fmt.Errorf("invalid address/es: %s", invalid)
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
