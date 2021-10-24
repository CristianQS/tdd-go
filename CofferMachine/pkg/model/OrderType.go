package model

type OrderType string

const (
	Tea       OrderType = "T"
	Chocolate OrderType = "H"
	Coffee    OrderType = "C"
	Message    OrderType = "M"
)
