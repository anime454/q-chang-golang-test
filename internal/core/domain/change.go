package domain

import "cashier-system/pkg/currency"

type Change struct {
	Banknotes []Banknote `json:"banknotes"`
	Coins     []Coin     `json:"coins"`
}

func (r Change) GetTotal() int {
	total := 0

	for _, banknote := range r.Banknotes {
		total += currency.BahtToSatang(banknote.Baht) * banknote.Amount
	}

	for _, coin := range r.Coins {
		total += currency.BahtToSatang(coin.Baht) * coin.Amount
	}

	return total
}
