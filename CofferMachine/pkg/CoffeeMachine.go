package pkg

type CoffeeMachine struct {
	drinkMaker DrinkMaker
}

func NewCoffeeMachine(drinkMaker DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker}
}

func (c *CoffeeMachine) Execute(order *Order) {
	factory := OrderFactory{}
	degradableOrder := factory.create(order)
	c.drinkMaker.execute(degradableOrder.CreateDrinkMakerCommand())
}
