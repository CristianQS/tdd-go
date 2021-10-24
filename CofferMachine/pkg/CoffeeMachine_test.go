package pkg

import (
	pkg "CofferMachine/pkg/model"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_create_drink_command_without_sugar(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *Order
		expected string
	}{
		{name: "Tea", drinkOrder: NewOrder(pkg.Tea, 0, 0.4), expected:"T::" },
		{name: "Chocolate", drinkOrder: NewOrder(pkg.Chocolate, 0, 0.5),expected:"H::" },
		{name: "Coffee", drinkOrder: NewOrder(pkg.Coffee, 0, 0.6),expected:"C::" },
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
		drinkOrder *Order
		expected string
	}{
		{name: "Tea", drinkOrder: NewOrder(pkg.Tea, 1, 0.4), expected:"T:1:0" },
		{name: "Chocolate", drinkOrder: NewOrder(pkg.Chocolate, 2, 0.5),expected:"H:2:0" },
		{name: "Coffee", drinkOrder: NewOrder(pkg.Coffee, 1, 0.6),expected:"C:1:0" },
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

func Test_create_info_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().execute(gomock.Eq("M:info-message")).Times(1)

	coffeeMachine.Execute(NewOrderMessage(pkg.Message,0,"info-message"))
}

func Test_create_drink_command_when_is_missing_money(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *Order
		expected string
	}{
		{name: "Tea", drinkOrder: NewOrder(pkg.Tea, 1, 0.3), expected:"M:Missing 0.10 € to get your drink" },
		{name: "Chocolate", drinkOrder: NewOrder(pkg.Chocolate, 2, 0.4),expected:"M:Missing 0.10 € to get your drink" },
		{name: "Coffee", drinkOrder: NewOrder(pkg.Coffee, 1, 0.5),expected:"M:Missing 0.10 € to get your drink" },
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

func Test_create_orange_command_without_sugar(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().execute(gomock.Eq("O::")).Times(1)

	coffeeMachine.Execute(NewOrder(pkg.Orange, 0, 0.6))

}


