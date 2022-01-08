package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jeronimobarea/simple-ethereum/constants"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
)

func Test_CheckBalance_ValidInput(t *testing.T) {
	address := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	token := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")

	expectedBalance := big.NewInt(10)

	mockApi := &MockApi{}

	mockResponse := &BalanceResponse{
		Balance: expectedBalance,
		Error:   nil,
	}
	mockApi.On(
		"CheckBalance",
		mock.Anything, // address
		mock.Anything, // token
	).Return(mockResponse, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.CheckBalance(address, token)
	assert.NoError(t, err, constants.ShouldNotFail)
	assert.Equal(t, expectedBalance, res.Balance)
	assert.NoError(t, res.Error, constants.ShouldNotFail)
}

func Test_CheckBalance_InvalidInput(t *testing.T) {
	address := common.HexToAddress("")
	token := common.HexToAddress("")

	var expectedRes *BalanceResponse

	mockApi := &MockApi{}

	mockApi.On(
		"CheckBalance",
		mock.Anything, // address
		mock.Anything, // token
	).Return(nil, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.CheckBalance(address, token)
	assert.Error(t, err, constants.ShouldFail)
	assert.EqualError(t, err, fmt.Sprintf("invalid address/es: %s", []interface{}{address, token}))
	assert.Equal(t, expectedRes, res)
}
