package client

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum"
	"math/big"

	"github.com/apex/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	simpleAddress "github.com/jeronimobarea/simple-ethereum/address"
	"github.com/jeronimobarea/simple-ethereum/transaction"
)

type apiImplementation struct {
	client *ethclient.Client
}

func NewApiImplementation(provider string) Api {
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.
			WithError(err).
			Fatal("error setting up ethereum service")
	}
	return &apiImplementation{client: client}
}

func (api *apiImplementation) SendTransaction(
	quantity *big.Int,
	fromPk *ecdsa.PrivateKey,
	to,
	token common.Address,
) (resp *TransactionResponse, err error) {
	from, err := simpleAddress.GetAddressFromPrivateKey(fromPk)
	if err != nil {
		log.
			WithError(err).
			Error("error retrieving the address")
		return
	}

	gasPrice, err := api.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.
			WithError(err).
			Error("error retrieving the gas price")
		return
	}

	data := transaction.BuildTransactionData(to, quantity)

	gasLimit, err := api.client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &token,
		Data: data,
	})
	if err != nil {
		log.
			WithError(err).
			Error("error retrieving the gas limit")
		return
	}

	nonce, err := api.client.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.
			WithField("from", from).
			WithError(err).
			Error("error retrieving the nonce")
		return
	}

	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		Value:    quantity,
		Data:     data,
	})

	chainID, err := api.client.NetworkID(context.Background())
	if err != nil {
		log.
			WithError(err).
			Error("error retrieving the chain Id")
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), fromPk)
	if err != nil {
		log.
			WithError(err).
			Error("error signing transaction")
		return
	}

	err = api.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.
			WithError(err).
			Error("error sending transaction")
		return
	}

	resp = &TransactionResponse{Transaction: signedTx}
	return
}

func (api *apiImplementation) CheckBalance(
	address,
	token common.Address,
) (resp *BalanceResponse, err error) {
	instance, err := NewToken(token, api.client)
	if err != nil {
		log.
			WithField("token_address", token).
			WithError(err).
			Error("error getting the token instance")
		return
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.
			WithField("address", address).
			WithField("token_address", token).
			WithError(err).
			Error("error retrieving the address balance")
	}

	resp = &BalanceResponse{
		Balance: balance,
		Error:   err,
	}
	return
}

func (api *apiImplementation) CheckBalances(
	addresses []common.Address,
	token common.Address,
) (resp *BalancesResponse, err error) {
	tmpBalances := make(map[common.Address]*BalanceResponse)

	for _, address := range addresses {
		var balance *BalanceResponse

		balance, err = api.CheckBalance(address, token)
		if err != nil {
			tmpBalances[address] = &BalanceResponse{
				Error: err,
			}
			continue
		}
		tmpBalances[address] = balance
	}

	resp = &BalancesResponse{Processed: tmpBalances}
	return
}
