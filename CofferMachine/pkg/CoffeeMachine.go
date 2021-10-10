package pkg

import (
	"CofferMachine/internal/enums"
	"CofferMachine/internal/model"
	"CofferMachine/pkg/Infraestructure"
	"fmt"
	"strconv"
)

type CoffeeMachine struct {
	drinkMaker Infraestructure.DrinkMaker
}

func NewCoffeeMachine(drinkMaker Infraestructure.DrinkMaker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker }
}

func (c *CoffeeMachine) Execute(orderDrink *model.OrderDrink) {
	drinkMakerCommand := CreateDrinkMakerCommand(orderDrink)
	c.drinkMaker.Execute(drinkMakerCommand)
}

func CreateDrinkMakerCommand(orderDrink *model.OrderDrink) string {
	var (
		sugar string
		sticks            string
		drinkMakerCommand string
	)
	if orderDrink.SugarQuantity > 0 {
		sugar = strconv.Itoa(orderDrink.SugarQuantity)
		sticks = "0"
	}
	if IsInfoMessageOrder(orderDrink) {
		drinkMakerCommand = WriteDrinkCommand(string(orderDrink.OrderType), sugar, sticks)
	} else {
		drinkMakerCommand = fmt.Sprintf("%s:%s", string(orderDrink.OrderType), orderDrink.Message)
	}
	return drinkMakerCommand
}

func IsInfoMessageOrder(order *model.OrderDrink) bool {
	return order.OrderType != enums.InfoMessage
}

func WriteDrinkCommand(character string, sugar string, sticks string) string {
	return fmt.Sprintf("%s:%s:%s", character, sugar, sticks)
}


