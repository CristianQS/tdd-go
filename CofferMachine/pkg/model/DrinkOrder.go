package model

import "fmt"

type DrinkOrder struct {
	character string
	sugarQuantity int
}

func NewDrinkOrder(character string, sugarQuantity int) *DrinkOrder {
	return &DrinkOrder{character: character, sugarQuantity: sugarQuantity}
}

func (d *DrinkOrder) CreateDrinkMakerCommand() string {
	if d.sugarQuantity > 0 {
		return fmt.Sprintf("%s:%d:0", d.character, d.sugarQuantity)
	}
	return fmt.Sprintf("%s::", d.character)
}
