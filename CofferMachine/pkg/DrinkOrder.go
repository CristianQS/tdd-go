package pkg

import pkg "CofferMachine/pkg/model"

type Order struct {
	drinkType     pkg.OrderType
	sugarQuantity int
	message string
}

func NewOrder(drinkType pkg.OrderType, sugarQuantity int) *Order {
	return &Order{drinkType: drinkType, sugarQuantity: sugarQuantity, message: ""}
}


func NewOrderMessage(drinkType pkg.OrderType, sugarQuantity int, message string) *Order {
	return &Order{drinkType: drinkType, sugarQuantity: sugarQuantity, message: message}
}
