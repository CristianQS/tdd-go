package repository

import (
	"CofferMachine/pkg/model"
)

type DrinkRepositoryInMemory struct{
	drinks map[*model.Drink]int
}

func NewOrderRepositoryInMemory() *DrinkRepositoryInMemory {
	return &DrinkRepositoryInMemory{drinks: make(map[*model.Drink]int) }
}

func (r *DrinkRepositoryInMemory) Add(order *model.Drink) {
	if r.drinks[order] != 0 {
		r.drinks[order] += 1
	} else {
		r.drinks[order] = 1
	}
}

func (DrinkRepositoryInMemory) GetOrders() map[model.Drink]int {
	panic("implement me")
}
