package sdk

import (
	"crypto/rand"

	"github.com/MixinNetwork/mixin/crypto"
)

type Address struct {
	PrivateSpendKey   crypto.Key  `json:"spend,omitempty"`
	PrivateViewKey    crypto.Key  `json:"view,omitempty"`
	PrivateEncryptKey *PrivateKey `json:"encrypt_key,omitempty"`

	PublicSpendKey   crypto.Key `json:"public_spend,omitempty"`
	PublicViewKey    crypto.Key `json:"public_view,omitempty"`
	PublicEncryptKey *PublicKey `json:"public_encrypt_key,omitempty"`
}

// TODO unimplement
func NewAddress() (*Address, error) {
	panic("unimplement")
	return nil, nil
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

func (a Address) Encrypt(m, s1, s2 []byte) (ct []byte, err error) {
	return a.PublicEncryptKey.EncryptWithRand(rand.Reader, m, s1, s2)
}

func (a Address) Decrypt(c, s1, s2 []byte) (m []byte, err error) {
	return a.PrivateEncryptKey.Decrypt(c, s1, s2)
}

// TODO unimplement
func (a Address) Sign(raw []byte) (string, error) {
	panic("unimplement")
	return "", nil
}

// TODO unimplement
func (a Address) Verify(raw, s []byte) error {
	panic("unimplement")
	return nil
}

// TODO unimplement
func (a Address) GhostKey(r crypto.Hash) (*crypto.Hash, error) {
	panic("unimplement")
	return nil, nil
}

func (a Address) VerifyGhostKey(mask, key crypto.Hash) error {
	panic("unimplement")
	return nil
}
