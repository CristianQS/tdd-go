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
	factory := factories.OrderFactory{}
	degradableOrder := factory.Create(order)
	drink := degradableOrder.GetDrink()
	// TODO Check if a drink is empty in CreateDrinkMakerCommand
	if IsDrinkOrder(order.DrinkType) && c.IsDrinkEmpty(drink) {
		c.drinkMaker.Execute("M:The drink selected is empty, we have already sent an email to refilled your drink :)")
		return
	}
	command := degradableOrder.CreateDrinkMakerCommand()
	c.drinkMaker.Execute(command)
	if IsDrinkOrder(command) {
		c.repository.Add(drink)
	}
	c.reportingLog.GetReport()
}

func (c *CoffeeMachine) IsDrinkEmpty(drink *model.Drink) bool {
	if c.beverageQuantityChecker.IsEmpty(drink.Name) {
		c.emailNotifier.NotifyMissingDrink(drink.Name)
		return true
	}
	return false
}

func IsDrinkOrder(command string) bool {
	return !strings.HasPrefix(command, model.Message)
}
