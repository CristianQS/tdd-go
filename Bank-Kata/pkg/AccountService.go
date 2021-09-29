package pkg

type AccountService interface {
	deposit(amount int)
	withdraw(amount int)
	printStatement()
}
