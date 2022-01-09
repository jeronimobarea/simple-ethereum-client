package common

import (
	"math/big"
	"testing"

	"github.com/tj/assert"
)

/*
 * SafeBalanceIsValid
 */
func Test_SafeBalanceIsValid_ValidBalance(t *testing.T) {
	validBalance := big.NewInt(100)
	valid := SafeBalanceIsValid(validBalance)
	assert.Equal(t, true, valid)
}

func Test_SafeBalanceIsValid_ZeroBalance(t *testing.T) {
	invalidBalance := big.NewInt(0)
	valid := SafeBalanceIsValid(invalidBalance)
	assert.Equal(t, false, valid)
}

func Test_SafeBalanceIsValid_NegativeBalance(t *testing.T) {
	invalidBalance := big.NewInt(-10)
	valid := SafeBalanceIsValid(invalidBalance)
	assert.Equal(t, false, valid)
}

func Test_SafeBalanceIsValid_NilBalance(t *testing.T) {
	valid := SafeBalanceIsValid(nil)
	assert.Equal(t, false, valid)
}

/*
 * BalanceIsValid
 */
func Test_BalanceIsValid_ValidBalance(t *testing.T) {
	validBalance := big.NewInt(100)
	valid, err := BalanceIsValid(validBalance)
	assert.NoError(t, err)
	assert.Equal(t, true, valid)
}

func Test_BalanceIsValid_ZeroBalance(t *testing.T) {
	invalidBalance := big.NewInt(0)
	valid, err := BalanceIsValid(invalidBalance)
	assert.NoError(t, err)
	assert.Equal(t, false, valid)
}

func Test_BalanceIsValid_NegativeBalance(t *testing.T) {
	invalidBalance := big.NewInt(-10)
	valid, err := BalanceIsValid(invalidBalance)
	assert.Error(t, err)
	assert.EqualError(t, err, "balance cannot be negative")
	assert.Equal(t, false, valid)
}

func Test_BalanceIsValid_NilBalance(t *testing.T) {
	valid, err := BalanceIsValid(nil)
	assert.Error(t, err)
	assert.EqualError(t, err, "balance cannot be nil")
	assert.Equal(t, false, valid)
}
