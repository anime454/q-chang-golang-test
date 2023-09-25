package port

import "cashier-system/internal/core/domain"

type CashierRepository interface {
	GetMaximumChange(pcId string) (domain.Change, error)
	NewMaximumChange(pcId string) (domain.Change, error)
	UpdateMaximumChange(pcId string, change domain.Change) error
}
