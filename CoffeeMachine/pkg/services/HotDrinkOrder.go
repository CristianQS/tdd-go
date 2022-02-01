package services

import (
	"CofferMachine/pkg/constants"
	"CofferMachine/pkg/infraestructure"
	"CofferMachine/pkg/model"
	"fmt"
)

type HotDrinkOrder struct {
	drink                   *model.Drink
	sugarQuantity           int
	moneyProvided           float64
	extraHot                bool
	emailNotifier           infraestructure.EmailNotifier
	beverageQuantityChecker infraestructure.BeverageQuantityChecker
}

func (d *HotDrinkOrder) GetDrink() *model.Drink {
	return d.drink
}

func NewHotDrinkOrder(drink *model.Drink, sugarQuantity int, provided float64, extraHot bool,
	emailNotifier infraestructure.EmailNotifier, beverageQuantityChecker infraestructure.BeverageQuantityChecker) *HotDrinkOrder {
	return &HotDrinkOrder{drink: drink, sugarQuantity: sugarQuantity, moneyProvided: provided,
		extraHot: extraHot, emailNotifier: emailNotifier, beverageQuantityChecker: beverageQuantityChecker}
}

func (d *HotDrinkOrder) CreateDrinkMakerCommand() string {
	if d.IsDrinkEmpty(d.drink) {
		return "M:The drink selected is empty, we have already sent an email to refilled your drink :)"
	}
	if d.moneyProvided < d.drink.Cost {
		missingMoney := d.drink.Cost - d.moneyProvided
		return constants.MissingMoneyMessage(missingMoney)
	}
	if d.sugarQuantity > 0 {
		return constants.DrinkingWithSugarMessage(fmt.Sprintf("%s%s", d.drink.Id, d.isExtraHot()), d.sugarQuantity)
	}
	return constants.DrinkingMessage(fmt.Sprintf("%s%s", d.drink.Id, d.isExtraHot()))
}

func (d *HotDrinkOrder) IsDrinkEmpty(drink *model.Drink) bool {
	if d.beverageQuantityChecker.IsEmpty(drink.Name) {
		d.emailNotifier.NotifyMissingDrink(drink.Name)
		return true
	}
	return false
}

func (d *HotDrinkOrder) isExtraHot() string {
	if d.extraHot {
		return "h"
	}
	return ""
}
