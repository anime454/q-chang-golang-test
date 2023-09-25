package cashiersrv

import (
	"cashier-system/internal/core/domain"
	"cashier-system/mocks/mockups"
	"cashier-system/pkg/appmsg"
	"encoding/json"
	"fmt"

	// "cashier-system/mocks/mockups"
	"reflect"
	"testing"
)

type args struct {
	r domain.DoOrderRequest
}

type testTemplate struct {
	name    string
	s       *service
	args    args
	want    domain.OrderResult
	wantErr bool
}

func Test_service_DoOrder(t *testing.T) {

	tests := []testTemplate{}
	cashierService := NewCashierService(mockups.NewLocalConfig())

	successCustomer := domain.Customer{
		Name:  "mockup-user",
		Money: 2000,
	}

	successProduct := domain.Product{
		Name:  "mockup-product",
		Price: 500,
		Stock: 1,
	}

	successOrderRequest := domain.DoOrderRequest{
		Customer: successCustomer,
		Product:  successProduct,
	}

	successOrderResult := domain.OrderResult{
		Order: domain.Order{
			Customer: successCustomer,
			Product:  successProduct,
			OrderChange: domain.OrderChange{
				Change: domain.Change{
					Banknotes: []domain.Banknote{
						{
							Baht:   1000,
							Satang: 100000,
							Amount: 1,
						},
						{
							Baht:   500,
							Satang: 50000,
							Amount: 1,
						},
					},
					Coins: nil,
				},
				Total: 1500,
			},
		},
		Status: appmsg.SUCCESS,
		Reason: appmsg.SUCCESS,
	}

	tests = append(tests, testTemplate{
		name: "Test_service_DoOrder",
		s:    cashierService,
		args: args{
			r: successOrderRequest,
		},
		want:    successOrderResult,
		wantErr: false,
	})

	// Failed: not enough money
	notEnoughMoneyCustomer := domain.Customer{
		Name:  "mockup-user",
		Money: 100,
	}

	notEnoughMoneyProduct := domain.Product{
		Name:  "mockup-product",
		Price: 10000,
		Stock: 1,
	}

	notEnoughMoneyOrderRequest := domain.DoOrderRequest{
		Customer: notEnoughMoneyCustomer,
		Product:  notEnoughMoneyProduct,
	}

	notEnoughMoneyOrderResult := domain.OrderResult{
		Order: domain.Order{
			Customer: notEnoughMoneyCustomer,
			Product:  notEnoughMoneyProduct,
			OrderChange: domain.OrderChange{
				Change: domain.Change{
					Banknotes: nil,
					Coins:     nil,
				},
				Total: 0,
			},
		},
		Status: appmsg.FAILED,
		Reason: appmsg.NOT_ENOUGH_MONEY,
	}

	tests = append(tests, testTemplate{
		name: "Test_service_DoOrder_not_enough_money",
		s:    cashierService,
		args: args{
			r: notEnoughMoneyOrderRequest,
		},
		want:    notEnoughMoneyOrderResult,
		wantErr: false,
	})

	// Failed: no change
	notEnoughChangeCustomer := domain.Customer{
		Name:  "mockup-user",
		Money: 1000000,
	}

	notEnoughChangeProduct := domain.Product{
		Name:  "mockup-product",
		Price: 1000000,
		Stock: 1,
	}

	notEnoughChangeOrderRequest := domain.DoOrderRequest{
		Customer: notEnoughChangeCustomer,
		Product:  notEnoughChangeProduct,
	}

	notEnoughChangeOrderResult := domain.OrderResult{
		Order: domain.Order{
			Customer: notEnoughChangeCustomer,
			Product:  notEnoughChangeProduct,
			OrderChange: domain.OrderChange{
				Change: domain.Change{
					Banknotes: nil,
					Coins:     nil,
				},
				Total: 0,
			},
		},
		Status: appmsg.FAILED,
		Reason: appmsg.NOT_ENOUGH_CHANGE,
	}

	tests = append(tests, testTemplate{
		name: "Test_service_DoOrder_no_change",
		s:    cashierService,
		args: args{
			r: notEnoughChangeOrderRequest,
		},
		want:    notEnoughChangeOrderResult,
		wantErr: false,
	})

	for _, tt := range tests {
		fmt.Println(tt.name)
		t.Run(tt.name, func(t *testing.T) {

			got, err := tt.s.DoOrder(tt.args.r)

			if (err != nil) != tt.wantErr {
				t.Errorf("service.DoOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				gotJson, _ := json.Marshal(got)
				wantJson, _ := json.Marshal(tt.want)
				t.Errorf("service.DoOrder() = %s, want %s", gotJson, wantJson)
			}

		})

	}

}
