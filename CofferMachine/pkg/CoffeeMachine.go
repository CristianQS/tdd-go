package pkg

import (
	"CofferMachine/pkg/infraestructure"
	"CofferMachine/pkg/model"
	"CofferMachine/pkg/repository"
	"strings"
)

type CoffeeMachine struct {
	drinkMaker DrinkMaker
	repository repository.DrinkRepository
	reportingLog *infraestructure.ReportingLog
}

func NewCoffeeMachine(drinkMaker DrinkMaker, repository repository.DrinkRepository, reportingLog *infraestructure.ReportingLog) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker, repository: repository, reportingLog: reportingLog}
}

func (c *CoffeeMachine) Execute(order *model.Order) {
	factory := model.OrderFactory{}
	degradableOrder := factory.Create(order)
	command := degradableOrder.CreateDrinkMakerCommand()
	c.drinkMaker.execute(command)
	if !strings.HasPrefix(command,"M:"){
		c.repository.Add(degradableOrder.GetDrink())
	}
	c.reportingLog.GetReport()
}
