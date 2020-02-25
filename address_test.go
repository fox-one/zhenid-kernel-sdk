package sdk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/btcsuite/btcutil/base58"
)

func TestGenerateAddress(t *testing.T) {
	address, _ := NewAddress()
	fmt.Println(*address)
}

func TestGenerateJWT(t *testing.T) {
	assert := assert.New(t)
	address, _ := NewAddress()
	fmt.Println(*address)

	fmt.Println(base58.Encode(address.privateSpendKey[:]))
	fmt.Println(base58.Encode(address.publicSpendKey[:]))
	token := address.ChainJWTToken()
	fmt.Println(token)
	verify := address.Verify(token)
	assert.True(verify, "verify ok")
}


func TestAddressKeyStone(t *testing.T) {
	assert := assert.New(t)
	const keyStone = "HXKS39AvpD8hPnz5orZ6rmbkaK8Bp389CdvZqbcd5NfKcSVQQPfnWrzoW7Zzj9EDNDxm3uPpTxetKRyw2rDcyRDQHJJfhschpRmv2aGVy9gGP8qzNcQFoEhSU4WgFL5tBzn9caa21okhjJ26QRDCwY7yftLU2ekTELdVpHLBCisb8ERFdZauLTJ6oYhD2W391DECP5X5RPxXqv6SSFCHUMrdVdjKzUVTBNqWPnGydjNnFZLf6txY8Fo5HTcNbWDgP1mE3J"
	address, error := AddressFromKeyStone(keyStone)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(address.KeyStone())
		addr := address.KeyStone()
		assert.Equal(addr, keyStone, "address not equal")
	}
}


func TestAddressFromString(t *testing.T) {
	assert := assert.New(t)
	const addressString = "HX2TrZbAJ8oHYrqJmqfJor5YGAhHB1RwzS4w3rC6abdAj3zciyNREFBg8LmdLTtgNZs5hG3Gjrf94L536TXTgoMRRATEhZwT9hdAyz7PWkgTQKRNMxJPFBwufT59DpJHRf26KZYFzf47QeaTLcV83Msh3mfmrNXhxfvPSBBdxHwfgqmZGheg54LaSjtVAit481Rbt7sTfNJaHuec9XSmAkZVW9Aa"
	address, error := AddressFromString(addressString)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(address)
		addr := address.String()
		assert.Equal(addr, addressString, "address not equal")
	}
}

func TestPublicGhostKey(t *testing.T) {
	const desAddressString = "HX2TrZbAJ8oHYrqJmqfJor5YGAhHB1RwzS4w3rC6abdAj3zciyNREFBg8LmdLTtgNZs5hG3Gjrf94L536TXTgoMRRATEhZwT9hdAyz7PWkgTQKRNMxJPFBwufT59DpJHRf26KZYFzf47QeaTLcV83Msh3mfmrNXhxfvPSBBdxHwfgqmZGheg54LaSjtVAit481Rbt7sTfNJaHuec9XSmAkZVW9Aa"
	desAddress, _ := AddressFromString(desAddressString)

	sk, _, _ := NewKey()

	ghostKey := desAddress.GhostPublicKey(sk.Convert(), 0)
	fmt.Println(ghostKey)
}

func TestPrivateGhostKey(t *testing.T) {
	address, _ := NewAddress()
	sk, _ , _ := NewKey()
	ghostKey := address.GhostPublicKey(sk.Convert(), 0)
	fmt.Println(ghostKey)
}

func TestSignAndVerify(t *testing.T) {
	address, _ := NewAddress()
	sk,_, _ := NewKey()
	ghosPrivatetKey := address.GhostPublicKey(sk.Convert(), 0)

	ghosPublicKey := address.GhostPublicKey(sk.Convert(), 0)
	fmt.Println(ghosPrivatetKey)
	fmt.Println(ghosPublicKey)
}
