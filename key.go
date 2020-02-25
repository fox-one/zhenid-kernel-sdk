package sdk

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/dgrijalva/jwt-go"
)

type Key crypto.Key
type PrivateEncryptyKey []byte
type PublicEncryptyKey []byte

type HXAddress struct {
	PrivateSpendKey crypto.Key
	PrivateViewKey  crypto.Key
	PublicSpendKey  crypto.Key
	PublicViewKey   crypto.Key
}

var (
	ErrED25519Verification = errors.New("ed25519: verification error")
)

func init() {
	key := &Key{}
	jwt.RegisterSigningMethod(key.Alg(), func() jwt.SigningMethod {
		return key
	})
}

func NewKey() (Key, Key, error) {
	seed := make([]byte, 64)
	_, err := rand.Read(seed)
	if err != nil {
		return Key{}, Key{}, err
	}
	return NewKeyWithSeed(seed)
}

func NewKeyWithSeed(seed []byte) (Key, Key, error) {
	if len(seed) != 64 {
		return Key{}, Key{}, errors.New("seed length must be 64")
	}
	address := NewAddressFromSeed(seed)

	return (Key)(address.PublicSpendKey), (Key)(address.PublicViewKey) , nil
}

func NewAddressFromSeed(seed []byte) HXAddress {
	hash1 := crypto.NewHash(seed)
	hash2 := crypto.NewHash(hash1[:])
	src := append(hash1[:], hash2[:]...)
	spend := crypto.NewKeyFromSeed(seed)
	view := crypto.NewKeyFromSeed(src)

	return HXAddress{
		PrivateSpendKey: spend,
		PrivateViewKey:  view,
		PublicSpendKey:  spend.Public(),
		PublicViewKey:   view.Public(),
	}
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

func (k Key) Alg() string {
	return "ED25519"
}

func (k Key) Verify(signingString, signature string, key interface{}) error {
	// Get the key
	var edKey *crypto.Key
	switch k := key.(type) {
	case *Key:
		edKey = (*crypto.Key)(k)
	case Key:
		edKey = (*crypto.Key)(&k)
	default:
		return jwt.ErrInvalidKeyType
	}

	var sig crypto.Signature
	{
		// Decode the signature
		var err error
		var s []byte

		fmt.Println("===Sign===")
		fmt.Println(signature)
		if s, err = jwt.DecodeSegment(signature); err != nil {
			return err
		}
		fmt.Println("======")
		copy(sig[:], s)
	}

	hasher := sha256.New()
	hasher.Write([]byte(signingString))

	if (*crypto.Key)(edKey).Verify(hasher.Sum(nil), sig) {
		return nil
	}
	return ErrED25519Verification
}

func (k Key) Sign(signingString string, key interface{}) (string, error) {
	// Get the key
	var edKey *crypto.Key
	switch k := key.(type) {
	case *Key:
		edKey = (*crypto.Key)(k)
	case Key:
		edKey = (*crypto.Key)(&k)
	default:
		return "", jwt.ErrInvalidKeyType
	}

	fmt.Println("===SignString===")
	fmt.Println(signingString)
	fmt.Println("===SignString===")

	hasher := sha256.New()
	hasher.Write([]byte(signingString))
	
	sig := edKey.Sign(hasher.Sum(nil))
	
	var s = make([]byte, 64)
	copy(s, sig[:])
	return jwt.EncodeSegment(s), nil
}
