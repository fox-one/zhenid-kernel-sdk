package sdk

import (
		// "io"
	"errors"
	"strings"

	// "crypto/rand"

	// "github.com/fox-one/zhenid-kernel-sdk/crypto/ecies"
	"github.com/MixinNetwork/mixin/crypto"
	"github.com/btcsuite/btcutil/base58"
	
)

const HXNetwork = "HX"

type Address struct {
	privateSpendKey   Key 
	privateViewKey    Key
	privateEncryptKey PrivateEncryptyKey //121 bytes

	publicSpendKey   Key
	publicViewKey    Key
	publicEncryptKey PublicEncryptyKey //91bytes
}

/// 生成新的地址， 地址有3对密钥构成，分别是 Spend Key、 View Key 和 Encrypt Key
///	Spend Key 和 View Key 主要是在恒信的区块链网络中使用
///	Ecrypt Key 主要作用是用于对 数据仓库中加密数据的密钥进行密钥交换
/// 三个密钥都不可以泄漏
func NewAddress() (*Address, error) {
	sk, err := NewKey()
	if err != nil {
		return nil, err
	}

	vk, err := NewKey()
	if err != nil {
		return nil, err
	}

	ek, err := NewECIESPrivateKey()
	if err != nil {
		return nil, err
	}

	sek, err := ek.Marshal()
	if err != nil {
		return nil, err
	}

	pek, err := ek.Public().Marshal()
	if err != nil {
		return nil, err
	}

	return &Address{
		privateSpendKey:   sk,
		privateViewKey:    vk,
		privateEncryptKey: sek,

		publicSpendKey:   sk.PublicKey(),
		publicViewKey:    vk.PublicKey(),
		publicEncryptKey: pek,
	}, nil
}


// 根据用户的地址生成对外的交易地址
// 地址的格式规定如下 
// Address： HX + Base58(Private Spend Key Private View Key +Private EncryptKey + CRC)
func (a Address) String() string {
	data := append([]byte(HXNetwork), a.publicSpendKey[:]...)
	data = append(data, a.publicViewKey[:]...)
	data = append(data, a.publicEncryptKey...)
	checksum := crypto.NewHash(data)

	data = append(a.publicSpendKey[:], a.publicViewKey[:]...)
    data = append(data, a.publicEncryptKey...)
	data = append(data, checksum[:4]...)
	
	return HXNetwork + base58.Encode(data)
}

// 通过字符串生成用户地址
// 字符串格式 
// 
// 返回地址或者错误
func AddressFromString(s string) (Address, error) {
	var a Address
	if !strings.HasPrefix(s, HXNetwork) {
		return a, errors.New("invalid address network")
	}
	data := base58.Decode(s[len(HXNetwork):])
	if len(data) != 159 {
		return a, errors.New("invalid address format")
	}
	// 分配 91 bytes 的存储空间
	a.publicEncryptKey = make([]byte, 91)
	
	copy(a.publicSpendKey[:], data[:32])
	copy(a.publicViewKey[:], data[32:64])
	copy(a.publicEncryptKey[:], data[64:])
	return a, nil
}

// 生成地址的JSON格式
//
//
//
func (a Address) MarshalJSON() ([]byte, error) {
	panic("unimplement")
	return nil, nil
}

// 通过JSON恢复地址
//
//
//
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
	pri, _ := ECIESPrivateKeyFromBytes(a.privateEncryptKey)
	return pri
}

func (a Address) PublicSpendKey() *crypto.Key {
	return a.publicSpendKey.Convert()
}

func (a Address) PublicViewKey() *crypto.Key {
	return a.publicViewKey.Convert()
}

func (a Address) PublicEncryptKey() *PublicKey {
	pub, _ := ECIESPublicKeyFromBytes(a.privateEncryptKey)
	return pub
}

// func (a Address) Encrypt(m, s1, s2 []byte) (ct []byte, err error) {
// 	return a.EncryptWithSeed(rand.Reader, m, s1, s2)
// }

// func (a Address) EncryptWithSeed(seed io.Reader, m, s1, s2 []byte) (ct []byte, err error) {
// 	pri := a.PrivateEncryptKey()
// 	return &pri.EncryptWithSeed(seed, m, s1, s2)
// }

// func (a Address) Decrypt(c, s1, s2 []byte) (m []byte, err error) {
// 	pub := a.PublicEncryptKey()
// 	return &pub.Decrypt(c, s1, s2)
// }

func (a Address) GhostPublicKey(r *crypto.Key, outputIndex uint64) *crypto.Key {
	return crypto.DeriveGhostPublicKey(r, a.PublicViewKey(), a.PublicSpendKey(), outputIndex)
}

func (a Address) GhostPrivateKey(mask *crypto.Key, outputIndex uint64) *crypto.Key {
	return crypto.DeriveGhostPrivateKey(mask, a.PrivateViewKey(), a.PrivateSpendKey(), outputIndex)
}

func (a Address) VerifyOutputs(outputs []*Output) []int {
	var verifiedOutputs = make([]int, 0, len(outputs))
	for idx, output := range outputs {
		for _, key := range output.Keys {
			if crypto.ViewGhostOutputKey(&key, a.PrivateViewKey(), &output.Mask, uint64(idx)).String() == a.PublicSpendKey().String() {
				verifiedOutputs = append(verifiedOutputs, idx)
				break
			}
		}
	}
	return verifiedOutputs
}
