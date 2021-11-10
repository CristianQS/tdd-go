package model

type DegradableOrder interface {
	CreateDrinkMakerCommand() string
	GetDrink() *Drink
}
