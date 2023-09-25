package mockups

import (
	"cashier-system/internal/core/domain"
	"cashier-system/pkg/appmsg"
	"errors"
)

type localConfig struct{}

func NewLocalConfig() *localConfig {
	return &localConfig{}
}

var db = map[string]domain.Change{}

func (l *localConfig) NewMaximumChange(pcId string) (domain.Change, error) {

	db[pcId] = domain.Change{
		Banknotes: []domain.Banknote{
			{Baht: 1000, Satang: 100000, Amount: 10},
			{Baht: 500, Satang: 50000, Amount: 20},
			{Baht: 100, Satang: 10000, Amount: 15},
			{Baht: 50, Satang: 5000, Amount: 20},
			{Baht: 20, Satang: 2000, Amount: 30},
		},
		Coins: []domain.Coin{
			{Baht: 10, Satang: 1000, Amount: 20},
			{Baht: 5, Satang: 500, Amount: 20},
			{Baht: 1, Satang: 100, Amount: 20},
			{Baht: 0.25, Satang: 25, Amount: 50},
		},
	}

	return db[pcId], nil

}

func (l *localConfig) GetMaximumChange(pcId string) (domain.Change, error) {

	if _, ok := db[pcId]; !ok {
		return domain.Change{}, errors.New(appmsg.DB_NOT_FOUND)
	}

	return db[pcId], nil

}

func (l *localConfig) UpdateMaximumChange(pcId string, change domain.Change) error {

	if _, ok := db[pcId]; !ok {
		return errors.New(appmsg.DB_NOT_FOUND)
	}

	db[pcId] = change

	return nil
}
