package services

import (
	"CofferMachine/pkg/constants"
	"CofferMachine/pkg/model"
)

type JuiceOrder struct {
	drink         *model.Drink
	moneyProvided float64
}

func (d *JuiceOrder) GetDrink() *model.Drink {
	return d.drink
}

func NewJuiceOrder(drink *model.Drink, moneyProvided float64) *JuiceOrder {
	return &JuiceOrder{drink: drink, moneyProvided: moneyProvided}
}

func (d *JuiceOrder) CreateDrinkMakerCommand() string {
	if d.moneyProvided < d.drink.Cost {
		missingMoney := d.drink.Cost - d.moneyProvided
		return constants.MissingMoneyMessage(missingMoney)
	}
	return constants.DrinkingMessage(d.drink.Id)
}
