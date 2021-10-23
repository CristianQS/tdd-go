package pkg

import "fmt"

type CoffeeMachine struct {
	drinkMaker DrinkMaker
}

func NewCoffeeMachine(drinkMaker DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker}
}

func (c *CoffeeMachine) Execute(drinkOrder *DrinkOrder) {
	command := CreateDrinkMakerCommand(drinkOrder)
	c.drinkMaker.execute(command)
}

func CreateDrinkMakerCommand(drinkOrder *DrinkOrder) string {
	if drinkOrder.sugarQuantity > 0 {
		return fmt.Sprintf("%s:%d:0", drinkOrder.character,drinkOrder.sugarQuantity)
	}
	return fmt.Sprintf("%s::", drinkOrder.character)
}
