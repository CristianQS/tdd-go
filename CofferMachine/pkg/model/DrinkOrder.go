package model

import "fmt"

type HotDrinkOrder struct {
	character string
	sugarQuantity int
	moneyProvided float64
	extraHot bool
}

var drinks = map[string]*Drink{
	Tea: NewDrink(Tea,0.4),
	Chocolate: NewDrink(Chocolate,0.5),
	Coffee: NewDrink(Coffee,0.6),
}

func NewHotDrinkOrder(character string, sugarQuantity int, provided float64, extraHot bool) *HotDrinkOrder {
	return &HotDrinkOrder{character: character, sugarQuantity: sugarQuantity, moneyProvided: provided,
		extraHot: extraHot}
}

func (d *HotDrinkOrder) CreateDrinkMakerCommand() string {
	if drink := drinks[d.character]; d.moneyProvided < drink.cost {
		missingMoney := drink.cost - d.moneyProvided
		return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
	}
	if d.sugarQuantity > 0 {
		return fmt.Sprintf("%s%s:%d:0", d.character,d.isExtraHot(), d.sugarQuantity)
	}
	return fmt.Sprintf("%s%s::", d.character, d.isExtraHot())
}

func (d *HotDrinkOrder) isExtraHot() string {
	if d.extraHot {return "h"}
	return ""
}
