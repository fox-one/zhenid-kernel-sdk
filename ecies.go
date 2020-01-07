package sdk

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"io"

	"github.com/ethereum/go-ethereum/crypto/ecies"
)

type PrivateKey ecies.PrivateKey
type PublicKey ecies.PublicKey

func NewECIESPrivateKey() (*PrivateKey, error) {
	pri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return (*PrivateKey)(ecies.ImportECDSA(pri)), nil
}

func ECIESPrivateKeyFromBytes(s []byte) (*PrivateKey, error) {
	pri, err := x509.ParseECPrivateKey(s)
	if err != nil {
		return nil, err
	}
	return (*PrivateKey)(ecies.ImportECDSA(pri)), nil
}

func ECIESPublicKeyFromBytes(s []byte) (*PublicKey, error) {
	k, err := x509.ParsePKIXPublicKey(s)
	if err != nil {
		return nil, err
	}
	pub, ok := k.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("invalid ecdsa public key")
	}
	return (*PublicKey)(ecies.ImportECDSAPublic(pub)), nil
}

func (pri *PrivateKey) Marshal() ([]byte, error) {
	return x509.MarshalECPrivateKey((*ecies.PrivateKey)(pri).ExportECDSA())
}

func (pri *PrivateKey) Public() *PublicKey {
	return (*PublicKey)(&pri.PublicKey)
}

func (pri *PrivateKey) ECIESPrivateKey() *ecies.PrivateKey {
	return (*ecies.PrivateKey)(pri)
}

func (pub *PublicKey) Marshal() ([]byte, error) {
	return x509.MarshalPKIXPublicKey((*ecies.PublicKey)(pub).ExportECDSA())
}

func (pri *PrivateKey) Decrypt(c, s1, s2 []byte) (m []byte, err error) {
	return (*ecies.PrivateKey)(pri).Decrypt(c, s1, s2)
}

func (pub *PublicKey) Encrypt(m, s1, s2 []byte) (ct []byte, err error) {
	return pub.EncryptWithSeed(rand.Reader, m, s1, s2)
}

func (pub *PublicKey) EncryptWithSeed(seed io.Reader, m, s1, s2 []byte) (ct []byte, err error) {
	return ecies.Encrypt(seed, (*ecies.PublicKey)(pub), m, s1, s2)
}
