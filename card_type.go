package sdk

const (
	_              = iota
	CardTypeID     = 1
	CardTypePolice = 2

	CardNameID     = "身份证"
	CardNamePolice = "警官证"
)

type Card struct {
	ID   int    `json:"card_id"`
	Name string `json:"name"`
}
