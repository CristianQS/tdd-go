package pkg

import "CofferMachine/pkg/model"

type CoffeeMachine struct {
	drinkMaker DrinkMaker
}

func NewCoffeeMachine(drinkMaker DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker}
}

func (c *CoffeeMachine) Execute(order *model.Order) {
	factory := model.OrderFactory{}
	degradableOrder := factory.Create(order)
	c.drinkMaker.execute(degradableOrder.CreateDrinkMakerCommand())
}
