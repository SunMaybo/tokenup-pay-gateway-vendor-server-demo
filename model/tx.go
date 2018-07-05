package model

type Transaction struct {
	ID           uint    `gorm:"primary_key" json:"-"`
	UserId       string  `json:"user_id"`
	OrderID      string  `json:"order_id"`
	Type         int     `json:"type"`
	CurrencyType string  `json:"currency_type"`
	Timestamp    int64   `json:"timestamp"`
	Address      string  `json:"address"`
	Amount       float64 `json:"amount"`
	Status       string  `json:"status"`
}

type TxStatus int

const (
	Pending     TxStatus = iota
	Failed
	UnConfirmed
	Confirms
)

func (s TxStatus) String() string {
	switch s {
	case Pending:
		return "PENDING"
	case Failed:
		return "FAILED"
	case UnConfirmed:
		return "UNCONFIRMED"
	case Confirms:
		return "CONFIRMS"
	default:
		return "PENDING"
	}
}
