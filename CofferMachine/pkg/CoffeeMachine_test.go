package pkg

import (
	"CofferMachine/internal/enums"
	"CofferMachine/internal/model"
	"CofferMachine/pkg/Mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func Test_should_create_a_tea_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := Mocks.NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().Execute(gomock.Eq("T::")).Times(1)

	coffeeMachine.Execute(model.NewOrderDrink(enums.Tea,0))
}

func Test_should_create_a_coffee_drink_with_no_sugar_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := Mocks.NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().Execute(gomock.Eq("C::")).Times(1)

	coffeeMachine.Execute(model.NewOrderDrink(enums.Coffee, 0))
}

func Test_should_create_a_chocolate_drink_with_two_sugar_and_one_stick_command(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := Mocks.NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().Execute(gomock.Eq("H:2:0")).Times(1)

	coffeeMachine.Execute(model.NewOrderDrink(enums.Chocolate, 2))
}

func Test_should_creates_an_info_message_when_customer_order_it(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDrinkMaker := Mocks.NewMockDrinkMaker(ctrl)
	var coffeeMachine = NewCoffeeMachine(mockDrinkMaker)

	mockDrinkMaker.EXPECT().Execute(gomock.Eq("M:content-message")).Times(1)

	coffeeMachine.Execute(&model.OrderDrink{OrderType: enums.InfoMessage, Message: "content-message"})
}
