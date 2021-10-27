package repository

import (
	"CofferMachine/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_add_an_order_to_repository(t *testing.T) {
	repository := NewOrderRepositoryInMemory()
	drink := model.NewDrink(model.Tea, 0.4)

	repository.Add(drink)
	assert.Equal(t,	repository.drinks[drink],1)
}
