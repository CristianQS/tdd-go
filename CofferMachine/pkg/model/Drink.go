package model

type Drink struct {
	Id   string
	Name string
	Cost float64
}

func NewDrink(id string, name string, cost float64) *Drink {
	return &Drink{Id: id, Name: name, Cost: cost}
}
