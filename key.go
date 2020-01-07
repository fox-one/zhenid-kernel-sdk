package sdk

import (
	"crypto/rand"
	"encoding/hex"
	"errors"

	"github.com/MixinNetwork/mixin/crypto"
)

type Key crypto.Key

func NewKey() (Key, error) {
	seed := make([]byte, 64)
	_, err := rand.Read(seed)
	if err != nil {
		return Key{}, err
	}
	return NewKeyWithSeed(seed)
}

func NewKeyWithSeed(seed []byte) (Key, error) {
	if len(seed) != 64 {
		return Key{}, errors.New("seed length must be 64")
	}
	return (Key)(crypto.NewKeyFromSeed(seed)), nil
}

func KeyFromString(s string) (Key, error) {
	var k Key
	err := k.UnmarshalJSON([]byte(s))
	return k, err
}

func (k Key) Convert() *crypto.Key {
	key := crypto.Key(k)
	if !key.HasValue() {
		return nil
	}
	return &key
}

func (k Key) PublicKey() Key {
	return Key(crypto.Key(k).Public())
}

func (k Key) MarshalJSON() ([]byte, error) {
	return []byte(hex.EncodeToString(k[:])), nil
}

func (k *Key) UnmarshalJSON(b []byte) error {
	data, err := hex.DecodeString(string(b))
	if err != nil {
		return err
	}

	if len(data) != 32 {
		return errors.New("key length must be 64")
	}

	copy(k[:], data)
	return nil
}
