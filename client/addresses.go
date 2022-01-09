package client

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
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

func (svc *ethereumService) SimpleCheckBalance(
	address,
	token common.Address,
) (resp *BalanceResponse, err error) {
	addresses := []interface{}{address, token}
	if _, invalid, _ := simpleCommon.ValidateAddresses(addresses); invalid != nil {
		err = fmt.Errorf("invalid address/es: %s", invalid)
		return
	}
	return svc.API.SimpleCheckBalance(address, token)
}

func (svc *ethereumService) SimpleCheckBalances(
	addresses []common.Address,
	token common.Address,
) (resp *BalancesResponse, err error) {
	return svc.API.SimpleCheckBalances(addresses, token)
}
