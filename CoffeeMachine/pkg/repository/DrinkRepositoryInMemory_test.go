package repository

import (
	"CofferMachine/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_add_a_drink_to_repository(t *testing.T) {
	repository := NewOrderRepositoryInMemory()
	drink := model.NewDrink(model.Tea, "Tea", 0.4)

	repository.Add(drink)

	assert.Equal(t, repository.drinks[drink], 1)
}

func Test_add_some_drink_to_repository(t *testing.T) {
	repository := NewOrderRepositoryInMemory()
	teaDrink := model.NewDrink(model.Tea, "Tea", 0.4)
	coffeeDrink := model.NewDrink(model.Coffee, "Coffee", 0.5)

	repository.Add(teaDrink)
	repository.Add(teaDrink)
	repository.Add(coffeeDrink)

	assert.Equal(t, repository.drinks[teaDrink], 2)
	assert.Equal(t, repository.drinks[coffeeDrink], 1)
}

func Test_get_all_drinks_from_repository(t *testing.T) {
	repository := NewOrderRepositoryInMemory()
	teaDrink := model.NewDrink(model.Tea, "Tea", 0.4)
	coffeeDrink := model.NewDrink(model.Coffee, "Coffee", 0.5)
	expectedDrinks := map[*model.Drink]int{
		teaDrink:    2,
		coffeeDrink: 1,
	}

	repository.Add(teaDrink)
	repository.Add(teaDrink)
	repository.Add(coffeeDrink)
	drinks := repository.GetDrinks()

	assert.Equal(t, drinks, expectedDrinks)
}
