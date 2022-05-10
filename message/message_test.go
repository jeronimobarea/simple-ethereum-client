package message

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tj/assert"
)

func Test_SignMessage(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		privateKey, err := crypto.GenerateKey()
		assert.NoError(t, err)
		message := "Some random message"
		_, err = SignMessage(privateKey, message)
		assert.NoError(t, err)
	})

	t.Run("invalid input", func(t *testing.T) {
		privateKey, err := crypto.GenerateKey()
		assert.NoError(t, err)
		_, err = SignMessage(privateKey, "")
		assert.Error(t, err)
	})

	t.Run("invalid key", func(t *testing.T) {
		message := "Some random message"
		_, err := SignMessage(nil, message)
		assert.Error(t, err)
	})

	t.Run("invalid params", func(t *testing.T) {
		_, err := SignMessage(nil, "")
		assert.Error(t, err)
	})
}
