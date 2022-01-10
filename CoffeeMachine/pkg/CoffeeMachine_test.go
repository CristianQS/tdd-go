package pkg

import (
	"CofferMachine/pkg/infraestructure"
	pkg "CofferMachine/pkg/model"
	"CofferMachine/pkg/repository"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_create_drink_command_without_sugar(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *pkg.Order
		expected string
	}{
		{name: "Tea", drinkOrder: pkg.NewOrder(pkg.Tea, 0, 0.4, false), expected:"T::" },
		{name: "Chocolate", drinkOrder: pkg.NewOrder(pkg.Chocolate, 0, 0.5, false),expected:"H::" },
		{name: "Coffee", drinkOrder: pkg.NewOrder(pkg.Coffee, 0, 0.6, false),expected:"C::" },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := NewMockDrinkMaker(ctrl)
			mockLogger := infraestructure.NewMockLogger(ctrl)
			mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)

			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)
			mockDrinkMaker.EXPECT().execute(gomock.Eq(tc.expected)).Times(1)

			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_drink_command_with_sugar(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *pkg.Order
		expected string
	}{
		{name: "Tea", drinkOrder: pkg.NewOrder(pkg.Tea, 1, 0.4, false), expected:"T:1:0" },
		{name: "Chocolate", drinkOrder: pkg.NewOrder(pkg.Chocolate, 2, 0.5, false),expected:"H:2:0" },
		{name: "Coffee", drinkOrder: pkg.NewOrder(pkg.Coffee, 1, 0.6, false),expected:"C:1:0" },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := NewMockDrinkMaker(ctrl)
			mockLogger := infraestructure.NewMockLogger(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
			reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)

			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockDrinkMaker.EXPECT().execute(gomock.Eq(tc.expected)).Times(1)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)

			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_info_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	mockLogger := infraestructure.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
	mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)

	mockLogger.EXPECT().PrintLine(gomock.Any()).Times(2)
	mockDrinkMaker.EXPECT().execute(gomock.Eq("M:info-message")).Times(1)

	coffeeMachine.Execute(pkg.NewOrderMessage(pkg.Message,0,"info-message"))
}

func Test_create_drink_command_when_is_missing_money(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *pkg.Order
		expected string
	}{
		{name: "Tea", drinkOrder: pkg.NewOrder(pkg.Tea, 1, 0.3, false), expected:"M:Missing 0.10 € to get your drink" },
		{name: "Chocolate", drinkOrder: pkg.NewOrder(pkg.Chocolate, 2, 0.4, false),expected:"M:Missing 0.10 € to get your drink" },
		{name: "Coffee", drinkOrder: pkg.NewOrder(pkg.Coffee, 1, 0.5, false),expected:"M:Missing 0.10 € to get your drink" },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := NewMockDrinkMaker(ctrl)
			mockLogger := infraestructure.NewMockLogger(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
			mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)

			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(2)
			mockDrinkMaker.EXPECT().execute(gomock.Eq(tc.expected)).Times(1)

			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_orange_command_without_sugar(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	mockLogger := infraestructure.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
	mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)

	mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq("Orange")).Return(false)
	mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)
	mockDrinkMaker.EXPECT().execute(gomock.Eq("O::")).Times(1)

	coffeeMachine.Execute(pkg.NewOrder(pkg.Orange, 0, 0.6, false))
}


func Test_create_extra_hot_drinks_command(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *pkg.Order
		expected string
	}{
		{name: "Tea", drinkOrder: pkg.NewOrder(pkg.Tea, 2, 0.4, true), expected:"Th:2:0" },
		{name: "Chocolate", drinkOrder: pkg.NewOrder(pkg.Chocolate, 1, 0.5, true),expected:"Hh:1:0" },
		{name: "Coffee", drinkOrder: pkg.NewOrder(pkg.Coffee, 0, 0.6, true),expected:"Ch::" },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := NewMockDrinkMaker(ctrl)
			mockLogger := infraestructure.NewMockLogger(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
			mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)

			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)
			mockDrinkMaker.EXPECT().execute(gomock.Eq(tc.expected)).Times(1)

			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_report_of_drinks_sold(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	mockLogger := infraestructure.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
	mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker,repositoryInMemory,reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)
	order := pkg.NewOrder(pkg.Orange, 0, 0.6, false)

	mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq("Orange")).Return(false)
	mockDrinkMaker.EXPECT().execute(gomock.Any()).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Drink | Sold")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Orange | 1")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Total Money Earned: 0.60")).Times(1)

	coffeeMachine.Execute(order)
}


func Test_should_print_in_console_the_shortage_and_email_notification_when_BeverageQuantity_is_empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	mockLogger := infraestructure.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger,repositoryInMemory)
	mockEmailNotifier := infraestructure.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker:= infraestructure.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker,repositoryInMemory,reportingLog,mockEmailNotifier,mockBeverageQuantityChecker)
	order := pkg.NewOrder(pkg.Orange, 0, 0.6, false)

	mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq("Orange")).Return(true)
	mockEmailNotifier.EXPECT().NotifyMissingDrink(gomock.Eq("Orange")).Times(1)
	mockDrinkMaker.EXPECT().execute(gomock.Eq("M:The drink selected is empty, we have already sent an email to refilled your drink :)")).Times(1)


	coffeeMachine.Execute(order)
}


