package pkg

import "fmt"

type CoffeeMachine struct {
	drinkMaker DrinkMaker
}

func NewCoffeeMachine(drinkMaker DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker}
}

func (c *CoffeeMachine) Execute(drinkOrder *DrinkOrder) {
	command := fmt.Sprintf("%s::", drinkOrder.character)
	c.drinkMaker.execute(command)
}
