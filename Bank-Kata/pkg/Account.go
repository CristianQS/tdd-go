package pkg

import "time"

type Account struct {
	transaction Transaction
}

func (a *Account) Transaction() Transaction {
	return a.transaction
}

func (a *Account) deposit(amount int) {
	a.transaction.balance += amount
	a.transaction.amount = amount
	a.transaction.timeTransaction = time.Now().Format("02-01-2006")
}

func (a *Account) withdraw(amount int) {
}

func (a *Account) printStatement() {
}
