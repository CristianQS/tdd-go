package services

import (
	"CofferMachine/pkg/constants"
	"CofferMachine/pkg/infraestructure"
	"CofferMachine/pkg/model"
)

type JuiceOrder struct {
	drink                   *model.Drink
	moneyProvided           float64
	emailNotifier           infraestructure.EmailNotifier
	beverageQuantityChecker infraestructure.BeverageQuantityChecker
}

func (j *JuiceOrder) GetDrink() *model.Drink {
	return j.drink
}

func NewJuiceOrder(drink *model.Drink, moneyProvided float64, emailNotifier infraestructure.EmailNotifier,
	beverageQuantityChecker infraestructure.BeverageQuantityChecker) *JuiceOrder {
	return &JuiceOrder{drink: drink, moneyProvided: moneyProvided, emailNotifier: emailNotifier,
		beverageQuantityChecker: beverageQuantityChecker}
}

func (j *JuiceOrder) CreateDrinkMakerCommand() string {
	if j.IsDrinkEmpty(j.drink) {
		return "M:The drink selected is empty, we have already sent an email to refilled your drink :)"
	}
	if j.moneyProvided < j.drink.Cost {
		missingMoney := j.drink.Cost - j.moneyProvided
		return constants.MissingMoneyMessage(missingMoney)
	}
	return constants.DrinkingMessage(j.drink.Id)
}

func (j *JuiceOrder) IsDrinkEmpty(drink *model.Drink) bool {
	if j.beverageQuantityChecker.IsEmpty(drink.Name) {
		j.emailNotifier.NotifyMissingDrink(drink.Name)
		return true
	}
	return false
}
