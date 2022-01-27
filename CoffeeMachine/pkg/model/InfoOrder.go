package model

import (
	"CofferMachine/pkg/constants"
)

type InfoOrder struct {
	character string
	message   string
}

func NewInfoOrder(character string, message string) *InfoOrder {
	return &InfoOrder{character: character, message: message}
}

func (o *InfoOrder) CreateDrinkMakerCommand() string {
	return constants.InfoMessage(o.character, o.message)
}

func (o *InfoOrder) GetDrink() *Drink {
	return nil
}
