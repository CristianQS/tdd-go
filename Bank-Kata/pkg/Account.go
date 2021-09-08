package pkg

import "time"

type Account struct {
	balance int
	timeTransaction time.Time
	amount int
}


func (a *Account) deposit(amount int) {
	a.balance += amount
	a.amount = amount
	a.timeTransaction = time.Now()
}

func (a *Account) withdraw(amount int) {
}

func (a *Account) printStatement() {
}

func (a *Account) GetDate() string {
	return a.timeTransaction.Format("DD/MM/YYYY")
}
func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) GetAmount() int {
	return a.amount
}