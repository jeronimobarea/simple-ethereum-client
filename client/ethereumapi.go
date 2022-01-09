package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"

	simpleCommon "github.com/jeronimobarea/simple-ethereum-client/common"
	simpleToken "github.com/jeronimobarea/simple-ethereum-client/token"
)

type apiImplementation struct {
	client *ethclient.Client
}

func NewApiImplementation(client *ethclient.Client) Api {
	if client == nil {
		panic("ethereum client cannot be nil")
	}
	return &apiImplementation{client: client}
}

func SafeNewApiImplementation(client *ethclient.Client) (Api, error) {
	if client == nil {
		return nil, errors.New("ethereum client cannot be nil")
	}
	return &apiImplementation{client: client}, nil
}

func (api *apiImplementation) SimpleSendTransaction(
	quantity *big.Int,
	fromPk *ecdsa.PrivateKey,
	to,
	token common.Address,
) (resp *TransactionResponse, err error) {
	from, err := simpleCommon.GetAddressFromPrivateKey(fromPk)
	if err != nil {
		return
	}

	gasPrice, err := api.client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	transferFnSignature := []byte("transfer(address,uint256)")

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(to.Bytes(), common.HashLength)
	paddedAmount := common.LeftPadBytes(quantity.Bytes(), common.HashLength)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	if err != nil {
		return
	}

	gasLimit, err := api.client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &token,
		Data: data,
	})
	if err != nil {
		return
	}

	nonce, err := api.client.PendingNonceAt(context.Background(), from)
	if err != nil {
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
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), fromPk)
	if err != nil {
		return
	}

	err = api.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}

	resp = &TransactionResponse{Transaction: signedTx}
	return
}

func (api *apiImplementation) SimpleCheckBalance(
	address,
	token common.Address,
) (resp *BalanceResponse, err error) {
	instance, err := simpleToken.NewToken(token, api.client)
	if err != nil {
		return
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, address)

	resp = &BalanceResponse{
		Balance: balance,
		Error:   err,
	}
	return
}

func (api *apiImplementation) SimpleCheckBalances(
	addresses []common.Address,
	token common.Address,
) (resp *BalancesResponse, err error) {
	tmpBalances := make(map[common.Address]*BalanceResponse)

	for _, address := range addresses {
		var balance *BalanceResponse

		balance, err = api.SimpleCheckBalance(address, token)
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
