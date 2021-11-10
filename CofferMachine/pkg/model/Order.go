package model

type Order struct {
	drinkType     string
	sugarQuantity int
	message string
	moneyProvided float64
	extraHot bool
}

func NewOrder(drinkType string, sugarQuantity int, moneyProvided float64, isExtraHot bool) *Order {
	return &Order{drinkType: drinkType, sugarQuantity: sugarQuantity, message: "",
		moneyProvided: moneyProvided, extraHot: isExtraHot}
}

func NewOrderMessage(drinkType string, sugarQuantity int, message string) *Order {
	return &Order{drinkType: drinkType, sugarQuantity: sugarQuantity, message: message}
}
