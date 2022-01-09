package client

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/mock"
	"github.com/tj/assert"
)

/*
 * SimpleSendTransaction
 */
func Test_SendTransaction_ValidInput(t *testing.T) {
	quantity := big.NewInt(10)
	fromPk, err := crypto.GenerateKey()
	assert.NoError(t, err)
	to := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	token := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")

	mockApi := &MockApi{}

	mockResponse := &TransactionResponse{
		Transaction: &types.Transaction{},
	}

	mockApi.On(
		"SimpleSendTransaction",
		mock.Anything, // quantity
		mock.Anything, // fromPk
		mock.Anything, // to
		mock.Anything, // token
	).Return(mockResponse, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleSendTransaction(quantity, fromPk, to, token)
	assert.NoError(t, err)
	assert.Equal(t, mockResponse, res)
}

func Test_SendTransaction_InvalidPrivateKey(t *testing.T) {
	quantity := big.NewInt(10)
	to := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	token := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")

	var expectedRes *TransactionResponse

	mockApi := &MockApi{}

	mockApi.On(
		"SimpleSendTransaction",
		mock.Anything, // quantity
		mock.Anything, // fromPk
		mock.Anything, // to
		mock.Anything, // token
	).Return(nil, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleSendTransaction(quantity, nil, to, token)
	assert.Error(t, err)
	assert.EqualError(t, err, "private key cannot be nil")
	assert.Equal(t, expectedRes, res)
}

func Test_SendTransaction_InvalidQuantity(t *testing.T) {
	quantity := big.NewInt(-10)
	fromPk, err := crypto.GenerateKey()
	assert.NoError(t, err)
	to := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	token := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")

	var expectedRes *TransactionResponse

	mockApi := &MockApi{}

	mockApi.On(
		"SimpleSendTransaction",
		mock.Anything, // quantity
		mock.Anything, // fromPk
		mock.Anything, // to
		mock.Anything, // token
	).Return(nil, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleSendTransaction(quantity, fromPk, to, token)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("quantity is not valid %d", quantity))
	assert.Equal(t, expectedRes, res)
}

func Test_SendTransaction_NilQuantity(t *testing.T) {
	fromPk, err := crypto.GenerateKey()
	assert.NoError(t, err)
	to := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	token := common.HexToAddress("0x323b5d4c32345ced77393b3530b1eed0f346429d")
	var quantity *big.Int

	var expectedRes *TransactionResponse

	mockApi := &MockApi{}

	mockApi.On(
		"SimpleSendTransaction",
		mock.Anything, // quantity
		mock.Anything, // fromPk
		mock.Anything, // to
		mock.Anything, // token
	).Return(nil, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleSendTransaction(quantity, fromPk, to, token)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("quantity is not valid %v", quantity))
	assert.Equal(t, expectedRes, res)
}

func Test_SendTransaction_InvalidAddress(t *testing.T) {
	quantity := big.NewInt(10)
	fromPk, err := crypto.GenerateKey()
	assert.NoError(t, err)
	to := common.HexToAddress("")
	token := common.HexToAddress("")

	var expectedRes *TransactionResponse

	mockApi := &MockApi{}

	mockApi.On(
		"SimpleSendTransaction",
		mock.Anything, // quantity
		mock.Anything, // fromPk
		mock.Anything, // to
		mock.Anything, // token
	).Return(nil, nil)

	service := NewService(&Resources{API: mockApi})

	res, err := service.SimpleSendTransaction(quantity, fromPk, to, token)
	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("invalid address/es: %s", []interface{}{to, token}))
	assert.Equal(t, expectedRes, res)
}
