package message

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tj/assert"
)

/*
 * SignMessage
 */
func Test_SignMessage_ValidInput(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	message := "Some random message"
	_, err = SignMessage(privateKey, message)
	assert.NoError(t, err)
}

func Test_SignMessage_InvalidInput(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	assert.NoError(t, err)
	_, err = SignMessage(privateKey, "")
	assert.Error(t, err)
}

func Test_SignMessage_InvalidKey(t *testing.T) {
	message := "Some random message"
	_, err := SignMessage(nil, message)
	assert.Error(t, err)
}

func Test_SignMessage_InvalidParams(t *testing.T) {
	_, err := SignMessage(nil, "")
	assert.Error(t, err)
}
