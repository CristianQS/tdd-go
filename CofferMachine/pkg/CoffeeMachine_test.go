package pkg

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_create_tea_command_without_sugar(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := NewMockDrinkMaker(ctrl)
	coffeeMachine := NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().execute(gomock.Eq("T::")).Times(1)

	coffeeMachine.Execute()
}

