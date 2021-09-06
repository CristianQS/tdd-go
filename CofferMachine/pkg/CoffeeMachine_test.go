package pkg

import (
	"CofferMachine/internal/enums"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_create_a_tea_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker, &Drink{DrinkType: enums.Tea, SugarQuantity: 0})

	mockDrinkMaker.EXPECT().execute(gomock.Eq("T::")).Times(1)

	coffeeMachine.Execute()
}

func Test_should_create_a_coffee_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker,&Drink{DrinkType: enums.Coffee, SugarQuantity: 0})

	mockDrinkMaker.EXPECT().execute(gomock.Eq("C::")).Times(1)

	coffeeMachine.Execute()
}

func Test_should_create_a_chocolate_drink_with_two_sugar_and_one_stick_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker,&Drink{DrinkType: enums.Chocolate, SugarQuantity: 2})

	mockDrinkMaker.EXPECT().execute(gomock.Eq("H:2:0")).Times(1)

	coffeeMachine.Execute()
}
