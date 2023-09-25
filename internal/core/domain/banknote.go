package domain

type Banknote struct {
	Baht   float64 `json:"baht"`
	Satang int     `json:"satang"`
	Amount int     `json:"amount"`
}
