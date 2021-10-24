package pkg

import "CofferMachine/pkg/model"

type OrderFactory struct{}

func (f *OrderFactory) create(order *Order) model.DegradableOrder {
	switch order.drinkType {
	case model.Message:
		return model.NewInfoOrder(order.drinkType,order.message)
	case model.Orange:
		return model.NewJuiceOrder(order.drinkType,order.moneyProvided)
	default:
		return model.NewDrinkOrder(order.drinkType,order.sugarQuantity, order.moneyProvided)
	}
}
