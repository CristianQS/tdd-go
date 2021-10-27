package model

import "fmt"

type HotDrinkOrder struct {
	drink *Drink
	sugarQuantity int
	moneyProvided float64
	extraHot bool
}

func NewHotDrinkOrder(drink *Drink, sugarQuantity int, provided float64, extraHot bool) *HotDrinkOrder {
	return &HotDrinkOrder{drink: drink, sugarQuantity: sugarQuantity, moneyProvided: provided,
		extraHot: extraHot}
}

func (d *HotDrinkOrder) CreateDrinkMakerCommand() string {
	if d.moneyProvided < d.drink.Cost {
		missingMoney := d.drink.Cost - d.moneyProvided
		return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
	}
	if d.sugarQuantity > 0 {
		return fmt.Sprintf("%s%s:%d:0", d.drink.Id,d.isExtraHot(), d.sugarQuantity)
	}
	return fmt.Sprintf("%s%s::", d.drink.Id, d.isExtraHot())
}

func (d *HotDrinkOrder) isExtraHot() string {
	if d.extraHot {return "h"}
	return ""
}
