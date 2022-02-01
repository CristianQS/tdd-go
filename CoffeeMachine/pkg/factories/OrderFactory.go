package factories

import (
	"CofferMachine/pkg/infraestructure"
	"CofferMachine/pkg/model"
	"CofferMachine/pkg/services"
)

type OrderFactory struct {
	emailNotifier           infraestructure.EmailNotifier
	beverageQuantityChecker infraestructure.BeverageQuantityChecker
}

func NewOrderFactory(emailNotifier infraestructure.EmailNotifier, beverageQuantityChecker infraestructure.BeverageQuantityChecker) *OrderFactory {
	return &OrderFactory{emailNotifier: emailNotifier, beverageQuantityChecker: beverageQuantityChecker}
}

var hotDrinks = map[string]*model.Drink{
	model.Tea:       model.NewDrink(model.Tea, "Tea", 0.4),
	model.Chocolate: model.NewDrink(model.Chocolate, "Chocolate", 0.5),
	model.Coffee:    model.NewDrink(model.Coffee, "Coffee", 0.6),
}
var juices = map[string]*model.Drink{
	model.Orange: model.NewDrink(model.Orange, "Orange", 0.6),
}

func (f *OrderFactory) Create(order *model.Order) services.DegradableOrder {
	if IsAHotDrink(order) {
		return services.NewHotDrinkOrder(hotDrinks[order.DrinkType], order.SugarQuantity, order.MoneyProvided,
			order.ExtraHot, f.emailNotifier, f.beverageQuantityChecker)
	}
	if IsAJuice(order) != nil {
		return services.NewJuiceOrder(juices[order.DrinkType], order.MoneyProvided, f.emailNotifier, f.beverageQuantityChecker)
	}
	return services.NewInfoOrder(order.DrinkType, order.Message)
}

func IsAJuice(order *model.Order) *model.Drink {
	return juices[order.DrinkType]
}

func IsAHotDrink(order *model.Order) bool {
	return hotDrinks[order.DrinkType] != nil
}
