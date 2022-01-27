package model

import (
	"CofferMachine/pkg/constants"
	"fmt"
)

type HotDrinkOrder struct {
	drink         *Drink
	sugarQuantity int
	moneyProvided float64
	extraHot      bool
}

func (d *HotDrinkOrder) GetDrink() *Drink {
	return d.drink
}

func NewHotDrinkOrder(drink *Drink, sugarQuantity int, provided float64, extraHot bool) *HotDrinkOrder {
	return &HotDrinkOrder{drink: drink, sugarQuantity: sugarQuantity, moneyProvided: provided,
		extraHot: extraHot}
}

func (d *HotDrinkOrder) CreateDrinkMakerCommand() string {
	if d.moneyProvided < d.drink.Cost {
		missingMoney := d.drink.Cost - d.moneyProvided
		return constants.MissingMoneyMessage(missingMoney)
	}
	if d.sugarQuantity > 0 {
		return constants.DrinkingWithSugarMessage(fmt.Sprintf("%s%s", d.drink.Id, d.isExtraHot()), d.sugarQuantity)
	}
	return constants.DrinkingMessage(fmt.Sprintf("%s%s", d.drink.Id, d.isExtraHot()))
}

func (d *HotDrinkOrder) isExtraHot() string {
	if d.extraHot {
		return "h"
	}
	return ""
}
