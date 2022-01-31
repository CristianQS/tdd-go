package repository

import (
	"CofferMachine/pkg/model"
)

type DrinkRepository interface {
	Add(order *model.Drink)
	GetDrinks() map[*model.Drink]int
}
