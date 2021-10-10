package model

import "CofferMachine/internal/enums"

type OrderDrink struct {
	OrderType     enums.OrderType
	SugarQuantity int
	Message       string
}

func NewOrderDrink(drinkType enums.OrderType, sugarQuantity int ) *OrderDrink {
	return &OrderDrink{OrderType: drinkType, SugarQuantity: sugarQuantity}
}
