package pkg

import "time"

type Account struct {
	Transactions  []*Transaction
	ActualBalance int
}

func (a *Account) GetTransactions() []*Transaction {
	return a.Transactions
}

func (a *Account) deposit(amount int) {
	a.ActualBalance += amount
	a.Transactions = append(a.Transactions, &Transaction {
		Amount:  amount,
		Balance: a.ActualBalance,
		Time:    time.Now().Format("02-01-2006")})
}

func (a *Account) withdraw(amount int) {
}

func (a *Account) printStatement() {
}
