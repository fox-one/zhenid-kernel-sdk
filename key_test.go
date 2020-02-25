package sdk

import (
	"testing"
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestKey(t *testing.T) {
	assert := assert.New(t)

	const (
		s = "86ff26f552ec6b25d3ab941748c2207166d93c9368b70d32e29c30316e90f404"
		k = "ec2551c5990f6d0c414073da024e517ccfdc29cc0b60cb05515004bd10b4be19"
	)

	sk, _, err := NewKeyWithSeed([]byte(s))
	
	assert.Nil(err, "new key with seed")

	{
		bts, err := sk.MarshalJSON()
		assert.Nil(err, "marshal")
		
		assert.Equal(k, string(bts), "key not matched")

		key1, err := KeyFromString(k)
		assert.Nil(err, "key from string")
		assert.Equal(sk, key1, "key not matched")
	}

	{
		claims := jwt.MapClaims{
			"exp": time.Now().AddDate(0, 0, 1).Unix(),
		}

		token, err := jwt.NewWithClaims(sk, claims).SignedString(sk)
		assert.Nil(err, "signed string")
		fmt.Println(token)

		parseFunc := func(token string, k interface{}) error {
			_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*Key)
				assert.True(ok, "signing method not match")

				return k, nil
			})
			return err
		}

		err = parseFunc(token, sk)
		assert.NotNil(err, "jwt parse")

		err = parseFunc(token, sk.PublicKey())
		assert.Nil(err, "jwt parse")
	}
}
