package common

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tj/assert"
)

/*
 * IsValidAddress
 */
func Test_IsValidAddress_ValidCommonAddress(t *testing.T) {
	validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	valid := IsValidAddress(validAddress)
	assert.Equal(t, true, valid)
}

func Test_IsValidAddress_ValidStringAddress(t *testing.T) {
	validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	valid := IsValidAddress(validAddress)
	assert.Equal(t, true, valid)
}

func Test_IsValidAddress_InvalidStringAddress(t *testing.T) {
	invalidAddress := "0xabc"
	valid := IsValidAddress(invalidAddress)
	assert.Equal(t, false, valid)
}

func Test_IsValidAddress_InvalidNilAddress(t *testing.T) {
	valid := IsValidAddress(nil)
	assert.Equal(t, false, valid)
}

/*
 * IsZeroAddress
 */
func Test_IsZeroAddress_ValidCommonAddress(t *testing.T) {
	validAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	valid := IsZeroAddress(validAddress)
	assert.Equal(t, true, valid)
}

func Test_IsZeroAddress_ValidStringAddress(t *testing.T) {
	validAddress := "0x0000000000000000000000000000000000000000"
	valid := IsZeroAddress(validAddress)
	assert.Equal(t, true, valid)
}

func Test_IsZeroAddress_InvalidAddress(t *testing.T) {
	invalidAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	valid := IsZeroAddress(invalidAddress)
	assert.Equal(t, false, valid)
}

func Test_IsZeroAddress_InvalidNilAddress(t *testing.T) {
	valid := IsZeroAddress(nil)
	assert.Equal(t, false, valid)
}

/*
 * ValidateAddress
 */
func Test_ValidateAddress_ValidCommonAddress(t *testing.T) {
	validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	valid := ValidateAddress(validAddress)
	assert.Equal(t, true, valid)
}

func Test_ValidateAddress_ValidStringAddress(t *testing.T) {
	validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
	valid := ValidateAddress(validAddress)
	assert.Equal(t, true, valid)
}

func Test_ValidateAddress_InvalidZeroAddress(t *testing.T) {
	invalidAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	valid := ValidateAddress(invalidAddress)
	assert.Equal(t, false, valid)
}

func Test_ValidateAddress_InvalidNilAddress(t *testing.T) {
	valid := ValidateAddress(nil)
	assert.Equal(t, false, valid)
}

/*
 * ValidateAddresses
 */
func Test_ValidateAddresses_AllValid(t *testing.T) {
	allValid := []interface{}{
		"0x323b5d4c32345ced77393b3530b1eed0f346429d",
		"0x183b8d4c32345ced78393b3530b1eed1f246499d",
	}
	valid, invalid, success := ValidateAddresses(allValid)
	assert.Equal(t, true, success)
	assert.Equal(t, 0, len(invalid))
	assert.Equal(t, len(allValid), len(valid))
}

func Test_ValidateAddresses_AllInvalid(t *testing.T) {
	allInvalid := []interface{}{
		"0x323b5d4c32345ced77393bfadsfasf3530b1eed0f346429d",
		"0x323b5d4c32345ced77393bfadsfasf3530b1fjkalsfjslk9d",
	}
	valid, invalid, success := ValidateAddresses(allInvalid)
	assert.Equal(t, false, success)
	assert.Equal(t, len(allInvalid), len(invalid))
	assert.Equal(t, 0, len(valid))
}

func Test_ValidateAddresses_PartialValid(t *testing.T) {
	partiallyValid := []interface{}{
		"0x323b5d4c32345ced77393b3530b1eed0f346429d",
		"0x323b5d4c32345ced77393bfadsfasf3530b1eed0f346429d",
		"jfkld;afjaksl;fa",
	}
	valid, invalid, success := ValidateAddresses(partiallyValid)
	assert.Equal(t, false, success)
	assert.Equal(t, 2, len(invalid))
	assert.Equal(t, 1, len(valid))
}

func Test_ValidateAddresses_EmptyInput(t *testing.T) {
	var empty []interface{}
	valid, invalid, success := ValidateAddresses(empty)
	assert.Equal(t, false, success)
	assert.Equal(t, 0, len(invalid))
	assert.Equal(t, 0, len(valid))
}

func Test_ValidateAddresses_NilInput(t *testing.T) {
	valid, invalid, success := ValidateAddresses(nil)
	assert.Equal(t, false, success)
	assert.Equal(t, 0, len(invalid))
	assert.Equal(t, 0, len(valid))
}
