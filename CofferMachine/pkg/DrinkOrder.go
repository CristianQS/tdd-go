package pkg

import pkg "CofferMachine/pkg/model"

type DrinkOrder struct {
	drinkType     pkg.DrinkType
	sugarQuantity int
}

func NewDrinkOrder(drinkType pkg.DrinkType, sugarQuantity int) *DrinkOrder {
	return &DrinkOrder{drinkType: drinkType, sugarQuantity: sugarQuantity}
}
