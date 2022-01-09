package common

import (
	"crypto/ecdsa"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// RawAddress Used for declaring ethereum addresses that
// is not common.Address type
type RawAddress string

func (a RawAddress) IsValid() bool {
	return IsValidAddress(a)
}

func (a RawAddress) ToString() string {
	return string(a)
}

func (a RawAddress) ToCommonAddress() common.Address {
	return common.HexToAddress(a.ToString())
}

type PrivateAddress string

func (a PrivateAddress) IsValid() bool {
	return IsValidAddress(a)
}

func (a PrivateAddress) ToString() string {
	return string(a)
}

func (a *PrivateAddress) ToECDSA() (pk *ecdsa.PrivateKey, err error) {
	address := strings.Replace(a.ToString(), "0x", "", 1)
	pk, err = crypto.HexToECDSA(address)
	if err != nil {
		return
	}
	return
}
