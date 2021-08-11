package pkg

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_create_a_tea_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker)

	coffeeMachine.Execute()

	mockDrinkMaker.EXPECT().execute(gomock.Eq("T::"))
}
