package model

type Order struct {
	DrinkType     string
	SugarQuantity int
	Message       string
	MoneyProvided float64
	ExtraHot      bool
}

func NewOrder(drinkType string, sugarQuantity int, moneyProvided float64, isExtraHot bool) *Order {
	return &Order{DrinkType: drinkType, SugarQuantity: sugarQuantity, Message: "",
		MoneyProvided: moneyProvided, ExtraHot: isExtraHot}
}

func NewOrderMessage(drinkType string, sugarQuantity int, message string) *Order {
	return &Order{DrinkType: drinkType, SugarQuantity: sugarQuantity, Message: message}
}
