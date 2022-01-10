package model

import (
	"fmt"
)

type InfoOrder struct {
	character string
	message string
}

func NewInfoOrder(character string, message string) *InfoOrder {
	return &InfoOrder{character: character, message: message}
}

func (o *InfoOrder) CreateDrinkMakerCommand() string {
	return fmt.Sprintf("%s:%s", o.character, o.message)
}

func (d *InfoOrder) GetDrink() *Drink {
	return nil
}


