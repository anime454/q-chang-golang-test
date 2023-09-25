package port

import "cashier-system/internal/core/domain"

type CashierService interface {
	DoOrder(o domain.DoOrderRequest) (domain.OrderResult, error)
}
