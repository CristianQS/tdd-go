package repository

import (
	"CofferMachine/pkg/model"
)

type DrinkRepository interface {
	Add(order *model.Drink)
	GetOrders()map[model.Drink]int
}
