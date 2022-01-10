package model

import "fmt"

type JuiceOrder struct {
	drink *Drink
	moneyProvided float64
}

func (d *JuiceOrder) GetDrink() *Drink {
	return d.drink
}

func NewJuiceOrder(drink *Drink, moneyProvided float64) *JuiceOrder {
	return &JuiceOrder{drink: drink, moneyProvided: moneyProvided}
}

func (d *JuiceOrder) CreateDrinkMakerCommand() string {
	if d.moneyProvided < d.drink.Cost {
		missingMoney := d.drink.Cost - d.moneyProvided
		return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
	}
	return fmt.Sprintf("%s::", d.drink.Id)
}
