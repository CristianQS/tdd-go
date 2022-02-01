package pkg

import (
	"CofferMachine/pkg/factories"
	"CofferMachine/pkg/infraestructure"
	"CofferMachine/pkg/model"
	"CofferMachine/pkg/repository"
	"strings"
)

type CoffeeMachine struct {
	drinkMaker              DrinkMaker
	repository              repository.DrinkRepository
	reportingLog            *infraestructure.ReportingLog
	emailNotifier           infraestructure.EmailNotifier
	beverageQuantityChecker infraestructure.BeverageQuantityChecker
}

func NewCoffeeMachine(drinkMaker DrinkMaker, repository repository.DrinkRepository, reportingLog *infraestructure.ReportingLog,
	emailNotifier infraestructure.EmailNotifier, beverageQuantityChecker infraestructure.BeverageQuantityChecker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker, repository: repository, reportingLog: reportingLog,
		emailNotifier: emailNotifier, beverageQuantityChecker: beverageQuantityChecker}
}

func (c *CoffeeMachine) Execute(order *model.Order) {
	factory := factories.NewOrderFactory(c.emailNotifier, c.beverageQuantityChecker)
	degradableOrder := factory.Create(order)
	command := degradableOrder.CreateDrinkMakerCommand()
	c.drinkMaker.Execute(command)
	if IsDrinkOrder(command) {
		c.repository.Add(degradableOrder.GetDrink())
	}
	c.reportingLog.GetReport()
}

func IsDrinkOrder(command string) bool {
	return !strings.HasPrefix(command, model.Message)
}
