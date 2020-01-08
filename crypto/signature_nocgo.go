package crypto

import (	
	"crypto/elliptic"
	"github.com/btcsuite/btcd/btcec"
)

// S256 returns an instance of the secp256k1 curve.
func S256() elliptic.Curve {
	return btcec.S256()
}
