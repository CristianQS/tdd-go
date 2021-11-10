package infraestructure

type EmailNotifier interface {
	NotifyMissingDrink(drink string)
}
