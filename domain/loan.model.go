package domain

type Loan struct {
	ID     int          `json:"id"`
	Amount float64      `json:"amount"`
	Status string       `json:"status"`
	OwedBy ResponseUser `json:"owedBy"`
}
