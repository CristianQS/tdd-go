package infraestructure

type BeverageQuantityChecker interface {
	IsEmpty(drink string)bool
}
