package sdk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	assert := assert.New(t)

	const (
		s = "86ff26f552ec6b25d3ab941748c2207166d93c9368b70d32e29c30316e90f404"
		k = "9721a9bd1e9fd2a6a999d9442a2d20974c67e8b55746b48d318634c314fa9302"
	)

	key, err := NewKeyWithSeed([]byte(s))
	assert.Nil(err, "new key with seed")

	{
		bts, err := key.MarshalJSON()
		assert.Nil(err, "marshal")
		assert.Equal(k, string(bts), "key not matched")

		key1, err := KeyFromString(k)
		assert.Nil(err, "key from string")
		assert.Equal(key, key1, "key not matched")
	}
}
