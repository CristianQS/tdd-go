package model

import "fmt"

type JuiceOrder struct {
	character string
	moneyProvided float64
}

var juices = map[string]*Drink{
	Orange: NewDrink(Orange,0.6),
}

func NewJuiceOrder(character string, moneyProvided float64) *JuiceOrder {
	return &JuiceOrder{character: character, moneyProvided: moneyProvided}
}

func (d *JuiceOrder) CreateDrinkMakerCommand() string {
	if juice := juices[d.character]; d.moneyProvided < juice.cost {
		missingMoney := juice.cost - d.moneyProvided
		return fmt.Sprintf("M:Missing %.2f € to get your drink", missingMoney)
	}
	return fmt.Sprintf("%s::", d.character)
}

