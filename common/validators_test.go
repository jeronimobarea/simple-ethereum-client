package common

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/tj/assert"
)

func Test_IsValidAddress(t *testing.T) {
	t.Run("valid common address", func(t *testing.T) {
		validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
		valid := IsValidAddress(validAddress)
		assert.Equal(t, true, valid)
	})

	t.Run("valid string address", func(t *testing.T) {
		validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
		valid := IsValidAddress(validAddress)
		assert.Equal(t, true, valid)
	})

	t.Run("invalid string address", func(t *testing.T) {
		invalidAddress := "0xabc"
		valid := IsValidAddress(invalidAddress)
		assert.Equal(t, false, valid)
	})

	t.Run("invalid nil address", func(t *testing.T) {
		valid := IsValidAddress(nil)
		assert.Equal(t, false, valid)
	})
}

func Test_IsZeroAddress(t *testing.T) {
	t.Run("valid common address", func(t *testing.T) {
		validAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
		valid := IsZeroAddress(validAddress)
		assert.Equal(t, true, valid)
	})

	t.Run("valid string address", func(t *testing.T) {
		validAddress := "0x0000000000000000000000000000000000000000"
		valid := IsZeroAddress(validAddress)
		assert.Equal(t, true, valid)
	})

	t.Run("invalid address", func(t *testing.T) {
		invalidAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
		valid := IsZeroAddress(invalidAddress)
		assert.Equal(t, false, valid)
	})

	t.Run("invalid nil address", func(t *testing.T) {
		valid := IsZeroAddress(nil)
		assert.Equal(t, false, valid)
	})
}

func Test_ValidateAddress(t *testing.T) {
	t.Run("valid common address", func(t *testing.T) {
		validAddress := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
		valid := ValidateAddress(validAddress)
		assert.Equal(t, true, valid)
	})

	t.Run("valid string address", func(t *testing.T) {
		validAddress := "0x323b5d4c32345ced77393b3530b1eed0f346429d"
		valid := ValidateAddress(validAddress)
		assert.Equal(t, true, valid)
	})

	t.Run("invalid zero address", func(t *testing.T) {
		invalidAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
		valid := ValidateAddress(invalidAddress)
		assert.Equal(t, false, valid)
	})

	t.Run("test invalid nil address", func(t *testing.T) {
		valid := ValidateAddress(nil)
		assert.Equal(t, false, valid)
	})
}

func Test_ValidateAddresses(t *testing.T) {
	t.Run("all valid", func(t *testing.T) {
		allValid := []interface{}{
			"0x323b5d4c32345ced77393b3530b1eed0f346429d",
			"0x183b8d4c32345ced78393b3530b1eed1f246499d",
		}
		valid, invalid, success := ValidateAddresses(allValid)
		assert.Equal(t, true, success)
		assert.Equal(t, 0, len(invalid))
		assert.Equal(t, len(allValid), len(valid))
	})

	t.Run("all invalid", func(t *testing.T) {
		allInvalid := []interface{}{
			"0x323b5d4c32345ced77393bfadsfasf3530b1eed0f346429d",
			"0x323b5d4c32345ced77393bfadsfasf3530b1fjkalsfjslk9d",
		}
		valid, invalid, success := ValidateAddresses(allInvalid)
		assert.Equal(t, false, success)
		assert.Equal(t, len(allInvalid), len(invalid))
		assert.Equal(t, 0, len(valid))
	})

	t.Run("partially valid", func(t *testing.T) {
		partiallyValid := []interface{}{
			"0x323b5d4c32345ced77393b3530b1eed0f346429d",
			"0x323b5d4c32345ced77393bfadsfasf3530b1eed0f346429d",
			"jfkld;afjaksl;fa",
		}
		valid, invalid, success := ValidateAddresses(partiallyValid)
		assert.Equal(t, false, success)
		assert.Equal(t, 2, len(invalid))
		assert.Equal(t, 1, len(valid))
	})

	t.Run("empty input", func(t *testing.T) {
		var empty []interface{}
		valid, invalid, success := ValidateAddresses(empty)
		assert.Equal(t, false, success)
		assert.Equal(t, 0, len(invalid))
		assert.Equal(t, 0, len(valid))
	})

	t.Run("nil input", func(t *testing.T) {
		valid, invalid, success := ValidateAddresses(nil)
		assert.Equal(t, false, success)
		assert.Equal(t, 0, len(invalid))
		assert.Equal(t, 0, len(valid))
	})
}
