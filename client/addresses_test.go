package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
)

/*
 * SimpleCheckBalance
 */
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
		"SimpleCheckBalance",
		mock.Anything, // address
		mock.Anything, // token
	).Return(mockResponse, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleCheckBalance(address, token)
	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, res.Balance)
	assert.NoError(t, res.Error)
}

func Test_CheckBalance_InvalidInput(t *testing.T) {
	address := common.HexToAddress("")
	token := common.HexToAddress("")

	var expectedRes *BalanceResponse

	mockApi := &MockApi{}

	mockApi.On(
		"SimpleCheckBalance",
		mock.Anything, // address
		mock.Anything, // token
	).Return(nil, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleCheckBalance(address, token)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("invalid address/es: %s", []interface{}{address, token}))
	assert.Equal(t, expectedRes, res)
}
