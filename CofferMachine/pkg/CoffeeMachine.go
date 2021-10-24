package pkg

import (
	pkg "CofferMachine/pkg/model"
	"fmt"
)

type CoffeeMachine struct {
	drinkMaker DrinkMaker
}

func NewCoffeeMachine(drinkMaker DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker}
}

func (c *CoffeeMachine) Execute(drinkOrder *Order) {
	command := CreateDrinkMakerCommand(drinkOrder)
	c.drinkMaker.execute(command)
}

func CreateDrinkMakerCommand(order *Order) string {
	if order.sugarQuantity > 0 {
		return fmt.Sprintf("%s:%d:0", order.drinkType, order.sugarQuantity)
	}
	if order.drinkType == pkg.Message {
		return fmt.Sprintf("%s:%s", order.drinkType, order.message)
	}
	return fmt.Sprintf("%s::", order.drinkType)
}
