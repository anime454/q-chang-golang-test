package domain

import (
	"cashier-system/pkg/appmsg"
	curr "cashier-system/pkg/currency"
)

type Order struct {
	Customer    Customer    `json:"customer"`
	Product     Product     `json:"product"`
	OrderChange OrderChange `json:"order_change"`
}

type DoOrderRequest struct {
	Customer Customer `json:"customer"`
	Product  Product  `json:"product"`
}

type OrderChange struct {
	Change Change  `json:"change"`
	Total  float64 `json:"total"`
}

type OrderResult struct {
	Order  Order  `json:"order"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func (r Order) CalculateChange(maximumChange *Change) (OrderChange, error) {
	productPrice := curr.BahtToSatang(r.Product.Price)
	customerPay := curr.BahtToSatang(r.Customer.Money)
	counter := customerPay - productPrice

	totalChange := curr.SatangToBaht(counter)
	result := OrderChange{
		Total: totalChange,
	}

	for i, v := range maximumChange.Banknotes {

		banknoteRequest := counter / v.Satang
		banknoteAvailable := maximumChange.Banknotes[i].Amount

		if counter <= 0 {
			break
		}

		if banknoteRequest <= 0 {
			continue
		}

		if banknoteRequest >= banknoteAvailable {
			banknoteRequest = banknoteAvailable
		}

		counter -= banknoteRequest * v.Satang
		banknoteAvailable -= banknoteRequest

		result.Change.Banknotes = append(result.Change.Banknotes, Banknote{
			Satang: v.Satang,
			Baht:   v.Baht,
			Amount: banknoteRequest,
		})

		maximumChange.Banknotes[i].Amount = banknoteAvailable
	}

	for i, v := range maximumChange.Coins {
		coinRequest := counter / v.Satang
		coinAvailable := maximumChange.Coins[i].Amount

		if counter <= 0 {
			break
		}

		if coinRequest <= 0 {
			continue
		}

		if coinRequest >= coinAvailable {
			coinRequest = coinAvailable
		}

		counter -= coinRequest * v.Satang
		coinAvailable -= coinRequest

		result.Change.Coins = append(result.Change.Coins, Coin{
			Satang: v.Satang,
			Baht:   v.Baht,
			Amount: coinRequest,
		})

		maximumChange.Coins[i].Amount = coinAvailable
	}

	return result, nil

}

func (r Order) OrderSuccess() OrderResult {
	return OrderResult{
		Order:  r,
		Status: appmsg.SUCCESS,
		Reason: appmsg.SUCCESS,
	}
}

func (r Order) OrderFailed(reason string) OrderResult {
	return OrderResult{
		Order:  r,
		Status: appmsg.FAILED,
		Reason: reason,
	}
}
