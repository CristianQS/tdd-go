package pkg

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_create_drink_command_without_sugar(t *testing.T) {
	tests := []struct{
		name string
		drinkOrder *DrinkOrder
		expected string
	}{
		{name: "Tea", drinkOrder: NewDrinkOrder("T"), expected:"T::" },
		{name: "Chocolate", drinkOrder: NewDrinkOrder("H"),expected:"H::" },
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


