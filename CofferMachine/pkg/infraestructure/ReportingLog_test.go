package infraestructure

import (
	"CofferMachine/pkg/model"
	"CofferMachine/pkg/repository"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_print_the_total_of_each_drink_was_sold(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockLogger := NewMockLogger(ctrl)
	mockRepository := repository.NewMockDrinkRepository(ctrl)
	reportingLog := NewReportingLog(mockLogger,mockRepository)
	expectedDrinks := GetExpectedDrinks()
	mockRepository.EXPECT().GetDrinks().Return(expectedDrinks)

	mockLogger.EXPECT().PrintLine(gomock.Eq("Drink | Soled")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("T | 2")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("C | 1")).Times(1)
	mockLogger.EXPECT().PrintLine(gomock.Eq("Total Money Earned: 1.30")).Times(1)

	reportingLog.GetReport()
}

func GetExpectedDrinks() map[*model.Drink]int {
	teaDrink := model.NewDrink(model.Tea, 0.4)
	coffeeDrink := model.NewDrink(model.Coffee, 0.5)
	return map[*model.Drink]int{
		teaDrink:    2,
		coffeeDrink: 1,
	}
}
