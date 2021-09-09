package pkg

import "time"

type Account struct {
	transactions  []*Transaction
	ActualBalance int
}

func (a *Account) GetTransactions() []*Transaction {
	return a.transactions
}

func (a *Account) deposit(amount int) {
	a.ActualBalance += amount
	a.transactions = append(a.transactions, &Transaction {
		amount: amount,
		balance: a.ActualBalance,
		timeTransaction: time.Now().Format("02-01-2006")})
}

func (a *Account) withdraw(amount int) {
}

func (a *Account) printStatement() {
}
