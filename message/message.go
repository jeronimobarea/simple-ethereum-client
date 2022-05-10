package message

import (
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

func SignMessage(
	privateKey *ecdsa.PrivateKey, message string,
) ([]byte, error) {
	if privateKey == nil {
		return nil, errors.New("error private key cannot be nil")
	}

	if message == "" {
		return nil, errors.New("error message cannot be empty")
	}

	data := []byte(message)
	hash := crypto.Keccak256Hash(data)

	return crypto.Sign(hash.Bytes(), privateKey)
}
