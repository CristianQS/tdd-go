package pkg

type DrinkOrder struct {
	character string
	sugarQuantity int
}

func NewDrinkOrder(character string, sugarQuantity int) *DrinkOrder {
	return &DrinkOrder{character: character, sugarQuantity: sugarQuantity}
}
