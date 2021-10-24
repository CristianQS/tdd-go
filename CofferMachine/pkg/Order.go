package pkg

type Order struct {
	drinkType	string
	sugarQuantity int
	message string
	moneyProvided float64
}

func NewOrder(drinkType string, sugarQuantity int, moneyProvided float64) *Order {
	return &Order{drinkType: drinkType, sugarQuantity: sugarQuantity, message: "", moneyProvided: moneyProvided}
}

func NewOrderMessage(drinkType string, sugarQuantity int, message string) *Order {
	return &Order{drinkType: drinkType, sugarQuantity: sugarQuantity, message: message}
}
