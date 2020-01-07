package sdk

import (
	"crypto/rand"

	"github.com/MixinNetwork/mixin/crypto"
)

type Address struct {
	privateSpendKey   Key
	privateViewKey    Key
	privateEncryptKey *PrivateKey

	publicSpendKey   Key
	publicViewKey    Key
	publicEncryptKey *PublicKey
}

func NewAddress() (*Address, error) {
	ek, err := NewECIESPrivateKey()
	if err != nil {
		return nil, err
	}

	sk, err := NewKey()
	if err != nil {
		return nil, err
	}

	vk, err := NewKey()
	if err != nil {
		return nil, err
	}

	return &Address{
		privateSpendKey:   sk,
		privateViewKey:    vk,
		privateEncryptKey: ek,

		publicSpendKey:   sk.PublicKey(),
		publicViewKey:    vk.PublicKey(),
		publicEncryptKey: ek.Public(),
	}, nil
}

// TODO unimplement
func AddressFromString(s string) (*Address, error) {
	panic("unimplement")
	return nil, nil
}

// TODO unimplement
func (a Address) MarshalJSON() ([]byte, error) {
	panic("unimplement")
	return nil, nil
}

// TODO unimplement
func (a *Address) UnmarshalJSON(b []byte) error {
	panic("unimplement")
	return nil
}

func (a Address) PrivateSpendKey() *crypto.Key {
	return a.privateSpendKey.Convert()
}

func (a Address) PrivateViewKey() *crypto.Key {
	return a.privateViewKey.Convert()
}

func (a Address) PrivateEncryptKey() *PrivateKey {
	return a.privateEncryptKey
}

func (a Address) PublicSpendKey() *crypto.Key {
	return a.publicSpendKey.Convert()
}

func (a Address) PublicViewKey() *crypto.Key {
	return a.publicViewKey.Convert()
}

func (a Address) PublicEncryptKey() *PublicKey {
	return a.publicEncryptKey
}

func (a Address) Encrypt(m, s1, s2 []byte) (ct []byte, err error) {
	return a.PublicEncryptKey().EncryptWithRand(rand.Reader, m, s1, s2)
}

func (a Address) Decrypt(c, s1, s2 []byte) (m []byte, err error) {
	return a.PrivateEncryptKey().Decrypt(c, s1, s2)
}

// TODO unimplement
func (a Address) Sign(raw []byte) (crypto.Signature, error) {
	panic("unimplement")
	var sig crypto.Signature
	return sig, nil
}

// TODO unimplement
func (a Address) Verify(raw, s []byte) error {
	panic("unimplement")
	return nil
}

// TODO unimplement
func (a Address) GhostPublicKey(r crypto.Key) (crypto.Key, error) {
	panic("unimplement")
	var key crypto.Key
	return key, nil
}

func (a Address) GhostPrivateKey(mask, key crypto.Hash) error {
	panic("unimplement")
	return nil
}
