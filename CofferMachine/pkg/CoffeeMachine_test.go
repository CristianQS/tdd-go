package pkg

import (
	pkg "CofferMachine/pkg/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_create_drink_command_without_sugar(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *DrinkOrder
		expected string
	}{
		{name: "Tea", drinkOrder: NewDrinkOrder(pkg.Tea, 0), expected:"T::" },
		{name: "Chocolate", drinkOrder: NewDrinkOrder(pkg.Chocolate, 0),expected:"H::" },
		{name: "Coffee", drinkOrder: NewDrinkOrder(pkg.Coffee, 0),expected:"C::" },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := NewMockDrinkMaker(ctrl)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker)

			mockDrinkMaker.EXPECT().execute(gomock.Eq(tc.expected)).Times(1)

			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}

func Test_create_drink_command_with_sugar(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *DrinkOrder
		expected string
	}{
		{name: "Tea", drinkOrder: NewDrinkOrder(pkg.Tea,1), expected:"T:1:0" },
		{name: "Chocolate", drinkOrder: NewDrinkOrder(pkg.Chocolate,2),expected:"H:2:0" },
		{name: "Coffee", drinkOrder: NewDrinkOrder(pkg.Coffee,1),expected:"C:1:0" },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDrinkMaker := NewMockDrinkMaker(ctrl)
			coffeeMachine := NewCoffeeMachine(mockDrinkMaker)

			mockDrinkMaker.EXPECT().execute(gomock.Eq(tc.expected)).Times(1)

			coffeeMachine.Execute(tc.drinkOrder)
		})
	}
}


