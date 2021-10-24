package model

type DegradableOrder interface {
	CreateDrinkMakerCommand() string
}
