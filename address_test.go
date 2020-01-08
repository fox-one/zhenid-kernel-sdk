package sdk

import (
	"testing"
	"fmt"
)

func TestGenerateAddress(t *testing.T) {
	address, _ := NewAddress()
	fmt.Println(*address)
}

func TestAddressFromString(t *testing.T) {
	const addressString = "HX2TrZbAJ8oHYrqJmqfJor5YGAhHB1RwzS4w3rC6abdAj3zciyNREFBg8LmdLTtgNZs5hG3Gjrf94L536TXTgoMRRATEhZwT9hdAyz7PWkgTQKRNMxJPFBwufT59DpJHRf26KZYFzf47QeaTLcV83Msh3mfmrNXhxfvPSBBdxHwfgqmZGheg54LaSjtVAit481Rbt7sTfNJaHuec9XSmAkZVW9Aa"
	address, error := AddressFromString(addressString)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(address)	
	}
}