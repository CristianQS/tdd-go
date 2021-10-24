package pkg

import "CofferMachine/pkg/model"

type OrderFactory struct{}

func (f *OrderFactory) create(order *Order) model.DegradableOrder {
	switch order.drinkType {
	case model.Message:
		return model.NewInfoOrder(string(order.drinkType),order.message)
	default:
		return model.NewDrinkOrder(string(order.drinkType),order.sugarQuantity)
	}
}
