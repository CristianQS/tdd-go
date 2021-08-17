package pkg

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_create_a_tea_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker, &Drink{DrinkType: "Tea", NumberSugar: 0})

	mockDrinkMaker.EXPECT().execute(gomock.Eq("T::")).Times(1)

	coffeeMachine.Execute()
}

func Test_should_create_a_coffee_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker,&Drink{DrinkType: "Coffee", NumberSugar: 0})

	mockDrinkMaker.EXPECT().execute(gomock.Eq("C::")).Times(1)

	coffeeMachine.Execute()
}
