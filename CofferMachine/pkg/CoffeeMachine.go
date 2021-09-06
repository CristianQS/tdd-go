package pkg

import (
	"CofferMachine/internal/model"
	"fmt"
	"strconv"
)

type CoffeeMachine struct {
	drinkMaker DrinkMaker
	drink *model.Drink
}

func NewCoffeeMachine(drinkMaker DrinkMaker, drink *model.Drink) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker, drink: drink}
}

func (c *CoffeeMachine) Execute() {
	var (
		sugar string
		sticks    string
	)

	if c.drink.SugarQuantity > 0 {
		sugar  = strconv.Itoa(c.drink.SugarQuantity)
		sticks = "0"
	}
	c.drinkMaker.execute(c.CreateDrinkMakerOrder(string(c.drink.DrinkType), sugar, sticks))

}

func (c *CoffeeMachine) CreateDrinkMakerOrder(character string, sugar string, sticks string) string {
	return fmt.Sprintf("%s:%s:%s", character, sugar, sticks)
}


