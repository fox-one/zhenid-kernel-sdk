
package sdk

import (
	"github.com/MixinNetwork/mixin/crypto"
)

type Input struct {
	Hash  crypto.Hash
	Index int    
}


type Output struct {
	Type   []byte          
	// Amount decimal.Decimal 
	Keys   []crypto.Key   

	// OutputTypeScript fields

	Mask   crypto.Key    
}


type Transaction struct {
	Version []byte
	Asset   crypto.Hash 
	Inputs  []*Input    
	Outputs []*Output  
	Extra   []byte      
}
