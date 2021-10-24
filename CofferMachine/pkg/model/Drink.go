package model

type Drink struct {
	id string
	cost float64
}

func NewDrink(id string, cost float64) *Drink {
	return &Drink{id: id, cost: cost}
}
