package pkg

type DrinkOrder struct {
	character string
}

func NewDrinkOrder(character string) *DrinkOrder {
	return &DrinkOrder{character: character}
}
