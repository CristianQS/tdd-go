package constants

import "fmt"

func MissingMoneyMessage(missingMoney float64) string {
	return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
}

func DrinkingMessage(drink string) string {
	return fmt.Sprintf("%s::", drink)
}
func DrinkingWithSugarMessage(drink string, sugarQuantity int) string {
	return fmt.Sprintf("%s:%d:0", drink, sugarQuantity)
}

func InfoMessage(character, message string) string {
	return fmt.Sprintf("%s:%s", character, message)
}
