package pkg

type Account struct {
	balance int
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) deposit(amount int) {
	a.balance += amount
}

func (a *Account) withdraw(amount int) {
}

func (a *Account) printStatement() {
}
