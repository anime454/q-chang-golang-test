package cashiersrv

import (
	"cashier-system/internal/core/domain"
	"cashier-system/internal/core/port"
	"cashier-system/pkg/appmsg"
	curr "cashier-system/pkg/currency"
)

type service struct {
	repo port.CashierRepository
}

func NewCashierService(repo port.CashierRepository) *service {
	return &service{repo: repo}
}

func (s *service) DoOrder(r domain.DoOrderRequest) (domain.OrderResult, error) {

	o := domain.Order{
		Customer: r.Customer,
		Product:  r.Product,
	}

	productPriceBht := curr.BahtToSatang(o.Product.Price)
	customerMoneyBht := curr.BahtToSatang(o.Customer.Money)

	if customerMoneyBht-productPriceBht < 0 {
		return o.OrderFailed(appmsg.NOT_ENOUGH_MONEY), nil
	}

	pcName := "pc-1"
	maximumChange, err := s.repo.NewMaximumChange(pcName)
	if err != nil {
		return o.OrderFailed(appmsg.GET_MAXIMUM_CHANGE_FAILED), err
	}

	// maximumChange, err := s.repo.GetMaximumChange(pcName)
	// if err != nil {
	// 	if err.Error() == appmsg.DB_NOT_FOUND {
	// 	} else {
	// 		return o.OrderFailed(appmsg.GET_MAXIMUM_CHANGE_FAILED), err
	// 	}
	// }

	if maximumChange.GetTotal() <= 0 {
		return o.OrderFailed(appmsg.NO_CHANGE), nil
	}

	if productPriceBht > maximumChange.GetTotal() {
		return o.OrderFailed(appmsg.NOT_ENOUGH_CHANGE), nil
	}

	o.OrderChange, err = o.CalculateChange(&maximumChange)
	if err != nil {
		return o.OrderFailed(appmsg.CALCULATE_CHANGE_FAILED), err
	}

	err = s.repo.UpdateMaximumChange("pc-1", maximumChange)
	if err != nil {
		return o.OrderFailed(appmsg.UPDATE_MAXIMUM_CHANGE_FAILED), err
	}

	return o.OrderSuccess(), nil

}
