package pkg

import (
	"fmt"
	"strconv"
)

type Drink struct {
	DrinkType string
	NumberSugar int
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
	if c.drink.NumberSugar > 0 {
		sugar  = strconv.Itoa(c.drink.NumberSugar)
		sticks = "1"
	}
	c.drinkMaker.execute(fmt.Sprintf("%s:%s:%s", character, sugar,sticks))
	
}


