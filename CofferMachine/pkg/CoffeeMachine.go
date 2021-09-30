package pkg

import (
	"CofferMachine/internal/enums"
	"CofferMachine/internal/model"
	"fmt"
	"strconv"
)

type CoffeeMachine struct {
	drinkMaker DrinkMaker
	drink *model.OrderDrink
}

func NewCoffeeMachine(drinkMaker DrinkMaker, drink *model.OrderDrink) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker, drink: drink}
}

func (c *CoffeeMachine) Execute() {
	var (
		sugar string
		sticks    string
		order string
	)

	if c.drink.SugarQuantity > 0{
		sugar  = strconv.Itoa(c.drink.SugarQuantity)
		sticks = "0"
	}
	if c.drink.DrinkType != enums.InfoMessage{
		order = c.CreateDrinkMakerOrder(string(c.drink.DrinkType), sugar, sticks)
	} else {
		order = fmt.Sprintf("%s:%s", string(c.drink.DrinkType), c.drink.Message)
	}
	c.drinkMaker.execute(order)

}

func (c *CoffeeMachine) CreateDrinkMakerOrder(character string, sugar string, sticks string) string {
	return fmt.Sprintf("%s:%s:%s", character, sugar, sticks)
}


