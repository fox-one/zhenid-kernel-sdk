package sdk

import (
	"bytes"
	"errors"
	"strconv"
	"strings"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/btcsuite/btcutil/base58"
)

const (
	AddressPrefix = "ZHENID-PRI:"
)

type Address struct {
	PrivateSpendKey   crypto.Key
	PrivateViewKey    crypto.Key
	PrivateEncryptKey *PrivateKey

	PublicSpendKey   crypto.Key
	PublicViewKey    crypto.Key
	PublicEncryptKey *PublicKey
}

func AddressFromString(s string) (Address, error) {
	var a Address
	if !strings.HasPrefix(s, AddressPrefix) {
		return a, errors.New("invalid address network")
	}
	data := base58.Decode(s[len(AddressPrefix):])
	if len(data) != 68 {
		return a, errors.New("invalid address format")
	}
	checksum := crypto.NewHash(append([]byte(AddressPrefix), data[:64+91]...))
	if !bytes.Equal(checksum[:4], data[64+91:]) {
		return a, errors.New("invalid address checksum")
	}
	pub, err := ECIESPublicKeyFromBytes(data[64 : 64+91])
	if err != nil {
		return a, err
	}
	a.PublicEncryptKey = pub
	copy(a.PublicSpendKey[:], data[:32])
	copy(a.PublicViewKey[:], data[32:])
	return a, nil
}

func (a Address) String() string {
	data := make([]byte, 64+91)
	copy(data[:32], a.PublicSpendKey[:])
	copy(data[32:64], a.PublicViewKey[:])
	pub, _ := a.PublicEncryptKey.Marshal()
	copy(data[64:], pub)
	checksum := crypto.NewHash(data)
	return AddressPrefix + base58.Encode(append(data, checksum[:4]...))
}

func (a Address) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(a.String())), nil
}

func (a *Address) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}
	m, err := AddressFromString(unquoted)
	if err != nil {
		return err
	}
	a.PrivateSpendKey = m.PrivateSpendKey
	a.PrivateViewKey = m.PrivateViewKey
	a.PrivateEncryptKey = m.PrivateEncryptKey
	a.PublicSpendKey = m.PublicSpendKey
	a.PublicViewKey = m.PublicViewKey
	a.PublicEncryptKey = m.PublicEncryptKey
	return nil
}
