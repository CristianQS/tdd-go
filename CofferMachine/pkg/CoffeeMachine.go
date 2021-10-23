package pkg

type CoffeeMachine struct {
	drinkMaker DrinkMaker
}

func NewCoffeeMachine(drinkMaker DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker}
}


func (c *CoffeeMachine) Execute() {
	c.drinkMaker.execute("T::")
}
