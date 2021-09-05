package pkg

import (
	"fmt"
	"strconv"
)

type Drink struct {
	DrinkType     string
	SugarQuantity int
}

type CoffeeMachine struct {
	drinkMaker DrinkMaker
	drink *Drink
}

func NewCoffeeMachine(drinkMaker DrinkMaker, drink *Drink) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker, drink: drink}
}

func (c *CoffeeMachine) Execute() {
	var (
		character string
		sugar string
		sticks    string
	)
	if c.drink.DrinkType == "Tea"{ character += "T"}
	if c.drink.DrinkType == "Coffee"{ character += "C"}
	if c.drink.DrinkType == "Chocolate"{ character += "H"}
	if c.drink.SugarQuantity > 0 {
		sugar  = strconv.Itoa(c.drink.SugarQuantity)
		sticks = "0"
	}
	c.drinkMaker.execute(c.CreateDrinkMakerOrder(character, sugar, sticks))

}

func (c *CoffeeMachine) CreateDrinkMakerOrder(character string, sugar string, sticks string) string {
	return fmt.Sprintf("%s:%s:%s", character, sugar, sticks)
}


