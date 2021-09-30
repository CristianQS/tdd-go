package model

import "CofferMachine/internal/enums"

type OrderDrink struct {
	DrinkType     enums.DrinkType
	SugarQuantity int
	Message       string
}
