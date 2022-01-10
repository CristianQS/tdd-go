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
	emailNotifier infraestructure.EmailNotifier
	beverageQuantityChecker infraestructure.BeverageQuantityChecker
}

func NewCoffeeMachine(drinkMaker DrinkMaker, repository repository.DrinkRepository, reportingLog *infraestructure.ReportingLog,
	emailNotifier infraestructure.EmailNotifier, beverageQuantityChecker infraestructure.BeverageQuantityChecker) *CoffeeMachine {
	return &CoffeeMachine{drinkMaker: drinkMaker, repository: repository, reportingLog: reportingLog,
		emailNotifier: emailNotifier, beverageQuantityChecker: beverageQuantityChecker}
}

func (c *CoffeeMachine) Execute(order *model.Order) {
	factory := model.OrderFactory{}
	degradableOrder := factory.Create(order)
	drink := degradableOrder.GetDrink()
	if  drink != nil {
		if c.beverageQuantityChecker.IsEmpty(drink.Name){
			c.emailNotifier.NotifyMissingDrink(drink.Name)
			c.drinkMaker.execute("M:The drink selected is empty, we have already sent an email to refilled your drink :)")
			return
		}
	}
	command := degradableOrder.CreateDrinkMakerCommand()
	c.drinkMaker.execute(command)
	if !IsOrderSold(command) {
		c.repository.Add(drink)
	}
	c.reportingLog.GetReport()
}

func IsOrderSold(command string) bool {
	return strings.HasPrefix(command, "M:")
}