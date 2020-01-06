package sdk

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestECIES(t *testing.T) {
	const (
		p = "3077020101042071e441a0fa1d8d822e9fa8abeba0e0504179098205d3e970c72c8fb59b378defa00a06082a8648ce3d030107a14403420004c5bfd3b66fc8d41efa9e946cf0bae040f059619847ff7b2e1f1f815bcbc0a352ad78cc600492ac528a4282366eb0ffcd4a79c8f5e36c9cbbfc6475e75135531b"
		P = "3059301306072a8648ce3d020106082a8648ce3d03010703420004c5bfd3b66fc8d41efa9e946cf0bae040f059619847ff7b2e1f1f815bcbc0a352ad78cc600492ac528a4282366eb0ffcd4a79c8f5e36c9cbbfc6475e75135531b"
		m = "just a test"
	)

	assert := assert.New(t)
	b, _ := hex.DecodeString(p)
	pri, err := ECIESPrivateKeyFromBytes(b)
	assert.Nil(err, "new ecies private key")
	pub := pri.Public()

	// private key marshal
	{
		bts, err := pri.Marshal()
		assert.Nil(err, "new ecies private key")
		assert.Equal(bts, b, "ecies private key marshal not matched")
	}

	// public key marshal
	{
		b, _ := hex.DecodeString(P)
		pub1, err := ECIESPublicKeyFromBytes(b)
		assert.Nil(err, "new ecies public key")
		assert.Equal(pub, pub1, "ecise public key not matched")
	}

	// encrypt & decrypt
	{
		bts, err := pub.Encrypt([]byte(m), nil, nil)
		assert.Nil(err, "ecies encrypt")

		plain, err := pri.Decrypt(bts, nil, nil)
		assert.Nil(err, "ecies decrypt")
		assert.Equal(m, string(plain), "decrypted message not matched")
	}
}
