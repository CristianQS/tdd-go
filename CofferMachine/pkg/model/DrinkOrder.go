package model

import "fmt"

type HotDrinkOrder struct {
	character string
	sugarQuantity int
	moneyProvided float64
}

var drinks = map[string]*Drink{
	Tea: NewDrink(Tea,0.4),
	Chocolate: NewDrink(Chocolate,0.5),
	Coffee: NewDrink(Coffee,0.6),
}

func NewDrinkOrder(character string, sugarQuantity int, provided float64) *HotDrinkOrder {
	return &HotDrinkOrder{character: character, sugarQuantity: sugarQuantity, moneyProvided: provided}
}

func (d *HotDrinkOrder) CreateDrinkMakerCommand() string {
	if drink := drinks[d.character]; d.moneyProvided < drink.cost {
		missingMoney := drink.cost - d.moneyProvided
		return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
	}
	if d.sugarQuantity > 0 {
		return fmt.Sprintf("%s:%d:0", d.character, d.sugarQuantity)
	}
	return fmt.Sprintf("%s::", d.character)
}
