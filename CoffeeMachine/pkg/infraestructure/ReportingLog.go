package infraestructure

import (
	"CofferMachine/pkg/repository"
	"fmt"
)

type ReportingLog struct{
	logger          Logger
	drinkRepository repository.DrinkRepository
}

func NewReportingLog(logger Logger, drinkRepository repository.DrinkRepository) *ReportingLog {
	return &ReportingLog{logger: logger, drinkRepository: drinkRepository}
}

func (r *ReportingLog) GetReport() {
	drinks := r.drinkRepository.GetDrinks()
	r.logger.PrintLine("Drink | Sold")
	totalAmount := 0.0
	for drink, i := range drinks {
		totalAmount += drink.Cost * float64(i)
		r.logger.PrintLine(fmt.Sprintf("%s | %d",drink.Name, i))
	}
	r.logger.PrintLine(fmt.Sprintf("Total Money Earned: %0.2f",totalAmount))
}

