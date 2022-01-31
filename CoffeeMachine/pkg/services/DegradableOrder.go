package services

import "CofferMachine/pkg/model"

type DegradableOrder interface {
	CreateDrinkMakerCommand() string
	GetDrink() *model.Drink
}
