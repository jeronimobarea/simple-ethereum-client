package common

import (
	"errors"
	"math/big"
)

// BalanceIsValid will return an error and a boolean
func BalanceIsValid(balance *big.Int) (bool, error) {
	if balance == nil {
		return false, errors.New("balance cannot be nil")
	}

	if !balance.IsUint64() {
		return false, errors.New("balance cannot be negative")
	}
	return balance.BitLen() > 0, nil
}

// BalanceMustBeValid will panic in case there is an error when processing the balance
func BalanceMustBeValid(balance *big.Int) bool {
	valid, err := BalanceIsValid(balance)
	if err != nil {
		panic(err)
	}
	return valid
}

// SafeBalanceIsValid Use this function when you want to behave with a normal bool check
func SafeBalanceIsValid(balance *big.Int) bool {
	valid, err := BalanceIsValid(balance)
	if err != nil {
		return false
	}
	return valid
}
