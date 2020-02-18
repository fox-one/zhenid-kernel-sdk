package sdk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAddress(t *testing.T) {
	address, _ := NewAddress()
	fmt.Println(*address)
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
