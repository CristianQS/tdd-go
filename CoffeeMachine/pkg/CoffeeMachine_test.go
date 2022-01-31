package pkg

import (
	"CofferMachine/pkg/infraestructure"
	"CofferMachine/pkg/mocks"
	"CofferMachine/pkg/model"
	"CofferMachine/pkg/repository"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_create_drink_command_without_sugar(t *testing.T) {
	tests := []struct {
		name       string
		drinkOrder *model.Order
		expected   string
	}{
		{name: "Tea", drinkOrder: model.NewOrder(model.Tea, 0, 0.4, false), expected: "T::"},
		{name: "Chocolate", drinkOrder: model.NewOrder(model.Chocolate, 0, 0.5, false), expected: "H::"},
		{name: "Coffee", drinkOrder: model.NewOrder(model.Coffee, 0, 0.6, false), expected: "C::"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			ctrl := gomock.NewController(t)
			mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
			mockLogger := mocks.NewMockLogger(ctrl)
			mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)

			// WHEN
			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)
			mockDrinkMaker.EXPECT().Execute(gomock.Eq(tc.expected)).Times(1)

			// THEN
			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_drink_command_with_sugar(t *testing.T) {
	tests := []struct {
		name       string
		drinkOrder *model.Order
		expected   string
	}{
		{name: "Tea", drinkOrder: model.NewOrder(model.Tea, 1, 0.4, false), expected: "T:1:0"},
		{name: "Chocolate", drinkOrder: model.NewOrder(model.Chocolate, 2, 0.5, false), expected: "H:2:0"},
		{name: "Coffee", drinkOrder: model.NewOrder(model.Coffee, 1, 0.6, false), expected: "C:1:0"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// GIVEN
			ctrl := gomock.NewController(t)
			mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
			mockLogger := mocks.NewMockLogger(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
			reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)

			// WHEN
			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockDrinkMaker.EXPECT().Execute(gomock.Eq(tc.expected)).Times(1)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)

			// THEN
			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_info_command(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
	mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)

	// WHEN
	mockLogger.EXPECT().PrintLine(gomock.Any()).Times(2)
	mockDrinkMaker.EXPECT().Execute(gomock.Eq("M:info-message")).Times(1)

	// THEN
	coffeeMachine.Execute(model.NewOrderMessage(model.Message, 0, "info-message"))
}

func Test_create_drink_command_when_is_missing_money(t *testing.T) {
	// GIVEN
	tests := []struct {
		name       string
		drinkOrder *model.Order
		expected   string
	}{
		{name: "Tea", drinkOrder: model.NewOrder(model.Tea, 1, 0.3, false), expected: "M:Missing 0.10 € to get your drink"},
		{name: "Chocolate", drinkOrder: model.NewOrder(model.Chocolate, 2, 0.4, false), expected: "M:Missing 0.10 € to get your drink"},
		{name: "Coffee", drinkOrder: model.NewOrder(model.Coffee, 1, 0.5, false), expected: "M:Missing 0.10 € to get your drink"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
			mockLogger := mocks.NewMockLogger(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
			mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)

			// WHEN
			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(2)
			mockDrinkMaker.EXPECT().Execute(gomock.Eq(tc.expected)).Times(1)

			// THEN
			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_orange_command_without_sugar(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
	mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)

	// WHEN
	mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq("Orange")).Return(false)
	mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)
	mockDrinkMaker.EXPECT().Execute(gomock.Eq("O::")).Times(1)

	// THEN
	coffeeMachine.Execute(model.NewOrder(model.Orange, 0, 0.6, false))
}

func Test_create_extra_hot_drinks_command(t *testing.T) {
	// GIVEN
	tests := []struct {
		name       string
		drinkOrder *model.Order
		expected   string
	}{
		{name: "Tea", drinkOrder: model.NewOrder(model.Tea, 2, 0.4, true), expected: "Th:2:0"},
		{name: "Chocolate", drinkOrder: model.NewOrder(model.Chocolate, 1, 0.5, true), expected: "Hh:1:0"},
		{name: "Coffee", drinkOrder: model.NewOrder(model.Coffee, 0, 0.6, true), expected: "Ch::"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
			mockLogger := mocks.NewMockLogger(ctrl)
			repositoryInMemory := repository.NewOrderRepositoryInMemory()
			reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
			mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
			mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)

			// WHEN
			mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq(tc.name)).Return(false)
			mockLogger.EXPECT().PrintLine(gomock.Any()).Times(3)
			mockDrinkMaker.EXPECT().Execute(gomock.Eq(tc.expected)).Times(1)

			// THEN
			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_report_of_drinks_sold(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
	mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)
	order := model.NewOrder(model.Orange, 0, 0.6, false)

	// WHEN
	mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq("Orange")).Return(false)
	mockDrinkMaker.EXPECT().Execute(gomock.Any()).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Drink | Sold")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Orange | 1")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Total Money Earned: 0.60")).Times(1)

	// THEN
	coffeeMachine.Execute(order)
}

func Test_should_print_in_console_the_shortage_and_email_notification_when_BeverageQuantity_is_empty(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	mockDrinkMaker := mocks.NewMockDrinkMaker(ctrl)
	mockLogger := mocks.NewMockLogger(ctrl)
	repositoryInMemory := repository.NewOrderRepositoryInMemory()
	reportingLog := infraestructure.NewReportingLog(mockLogger, repositoryInMemory)
	mockEmailNotifier := mocks.NewMockEmailNotifier(ctrl)
	mockBeverageQuantityChecker := mocks.NewMockBeverageQuantityChecker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker, repositoryInMemory, reportingLog, mockEmailNotifier, mockBeverageQuantityChecker)
	order := model.NewOrder(model.Orange, 0, 0.6, false)

	// WHEN
	mockBeverageQuantityChecker.EXPECT().IsEmpty(gomock.Eq("Orange")).Return(true)
	mockEmailNotifier.EXPECT().NotifyMissingDrink(gomock.Eq("Orange")).Times(1)
	mockDrinkMaker.EXPECT().Execute(gomock.Eq("M:The drink selected is empty, we have already sent an email to refilled your drink :)")).Times(1)

	// THEN
	coffeeMachine.Execute(order)
}
