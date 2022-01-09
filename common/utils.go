package common

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetAddressFromPrivateKey(
	privateKey *ecdsa.PrivateKey,
) (address common.Address, err error) {
	if privateKey == nil {
		return common.Address{}, errors.New("error private key can not be nil")
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("error retrieving public key")
		return
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	return
}

func MustGetAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) common.Address {
	address, err := GetAddressFromPrivateKey(privateKey)
	if err != nil {
		panic(err)
	}
	return address
}
