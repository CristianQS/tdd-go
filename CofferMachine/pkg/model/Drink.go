package model

type Drink struct {
	Id   string
	Cost float64
}

func NewDrink(id string, cost float64) *Drink {
	return &Drink{Id: id, Cost: cost}
}
