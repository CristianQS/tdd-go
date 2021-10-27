package model

type OrderFactory struct{}


var hotDrinks = map[string]*Drink{
	Tea: NewDrink(Tea,"Tea",0.4),
	Chocolate: NewDrink(Chocolate,"Chocolate",0.5),
	Coffee: NewDrink(Coffee,"Coffee",0.6),
}
var juices = map[string]*Drink{
	Orange: NewDrink(Orange,"Orange",0.6),
}

func (f *OrderFactory) Create(order *Order) DegradableOrder {
	if hotDrinks[order.drinkType] != nil {
		return NewHotDrinkOrder(hotDrinks[order.drinkType],order.sugarQuantity,order.moneyProvided, order.extraHot)
	}
	if juices[order.drinkType] != nil {
		return NewJuiceOrder(juices[order.drinkType],order.moneyProvided)
	}
	return NewInfoOrder(order.drinkType,order.message)
}
