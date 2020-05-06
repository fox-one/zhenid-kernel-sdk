package sdk

import (
	"github.com/fox-one/mixin/common"
	"github.com/fox-one/mixin/crypto"
	"github.com/shopspring/decimal"
)

type Input struct {
	Hash  crypto.Hash `json:"hash,omitempty"`
	Index int         `json:"index,omitempty"`
}

type Output struct {
	Type   uint8           `json:"type"`
	Amount decimal.Decimal `json:"amount"`
	Keys   []crypto.Key    `json:"keys,omitempty"`

	// OutputTypeScript fields
	Script common.Script `json:"script,omitempty"`
	Mask   crypto.Key    `json:"mask,omitempty"`
}

type Transaction struct {
	Version uint8       `json:"version"`
	Asset   crypto.Hash `json:"asset"`
	Inputs  []*Input    `json:"inputs"`
	Outputs []*Output   `json:"outputs"`
	Extra   []byte      `json:"extra,omitempty"`
}

// TODO unimplement
func (o Output) Verify(msg, sig []byte) bool {
	panic("unimplement")
	return false
}

// TODO unimplement
// MarshalRaw marshal raw message for signing transaction
func (t Transaction) MarshalRaw() ([]byte, error) {
	panic("unimplement")
	return nil, nil
}
