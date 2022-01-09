package common

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tj/assert"
)

/*
 * GetAddressFromPrivateKey
 */
func Test_GetAddressFromPrivateKey_ValidKey(t *testing.T) {
	validKey, err := crypto.HexToECDSA("09e910621c2e988e9f7f6ffcd7024f54ec1461fa6e86a4b545e9e1fe21c28866")
	assert.NoError(t, err)
	address, err := GetAddressFromPrivateKey(validKey)
	assert.NoError(t, err)
	assert.Equal(t, common.HexToAddress("0x00B54E93EE2EBA3086A55F4249873E291D1AB06C"), address)
}

func Test_GetAddressFromPrivateKey_NilKey(t *testing.T) {
	address, err := GetAddressFromPrivateKey(nil)
	assert.Error(t, err)
	assert.Equal(t, common.Address{}, address)
}
