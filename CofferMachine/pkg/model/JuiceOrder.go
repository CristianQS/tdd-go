package model

import "fmt"

type JuiceOrder struct {
	drink *Drink
	moneyProvided float64
}

func NewJuiceOrder(drink *Drink, moneyProvided float64) *JuiceOrder {
	return &JuiceOrder{drink: drink, moneyProvided: moneyProvided}
}

func (d *JuiceOrder) CreateDrinkMakerCommand() string {
	if d.moneyProvided < d.drink.cost {
		missingMoney := d.drink.cost - d.moneyProvided
		return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
	}
	return fmt.Sprintf("%s::", d.drink.id)
}

