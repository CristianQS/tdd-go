package infraestructure

import (
	"CofferMachine/pkg/mocks"
	"CofferMachine/pkg/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_print_the_total_of_each_drink_was_sold(t *testing.T) {
	// GIVEN
	ctrl := gomock.NewController(t)
	mockLogger := mocks.NewMockLogger(ctrl)
	mockRepository := mocks.NewMockDrinkRepository(ctrl)
	reportingLog := NewReportingLog(mockLogger, mockRepository)
	expectedDrinks := GetExpectedDrinks()
	mockRepository.EXPECT().GetDrinks().Return(expectedDrinks)

	// THEN
	mockLogger.EXPECT().PrintLine(gomock.Eq("Drink | Sold")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Tea | 2")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Coffee | 1")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Total Money Earned: 1.30")).Times(1)

	// WHEN
	reportingLog.GetReport()
}

func GetExpectedDrinks() map[*model.Drink]int {
	teaDrink := model.NewDrink(model.Tea, "Tea", 0.4)
	coffeeDrink := model.NewDrink(model.Coffee, "Coffee", 0.5)
	return map[*model.Drink]int{
		teaDrink:    2,
		coffeeDrink: 1,
	}
}
