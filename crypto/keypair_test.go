package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestGeneratePrivateKey(t *testing.T) {
// 	privKey := GeneratePrivateKey()
// 	pubKey := privKey.PublicKey()
// 	//address := pubKey.Address()

// 	msg := []byte("Hello world")
// 	sig, err := privKey.Sign(msg)
// 	assert.Nil(t, err)

// 	b := sig.Verify(pubKey, msg)

// 	assert.True(t, b)

// }

func TestKeypairSignVerifyValid(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.PublicKey()
	msg := []byte("Hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	b := sig.Verify(publicKey, msg)

	assert.True(t, b)

}

func TestKeypairSignVerifyFail(t *testing.T) {
	privKey := GeneratePrivateKey()
	PublicKey := privKey.PublicKey()
	msg := []byte("Hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(PublicKey, []byte("xxxxxx")))

}
