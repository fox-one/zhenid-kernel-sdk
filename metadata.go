package sdk

type Metadata struct {
	Version    int   `json:"v"`
	CardType   int   `json:"t,omitempty"`
	StartDate  int64 `json:"s,omitempty"`
	ExpireDate int64 `json:"exp,omitempty"`
}
