package pkg

import (
	"bankkata/pkg/Infraestructure"
	"fmt"
	"time"
)

type Account struct {
	logger Infraestructure.Logger
	Transactions  []*Transaction
	ActualBalance int
}

func NewAccount(logger Infraestructure.Logger) *Account {
	return &Account{logger: logger}
}

func (a *Account) GetTransactions() []*Transaction {
	return a.Transactions
}

func (a *Account) deposit(amount int) {
	a.ActualBalance += amount
	a.Transactions = append(a.Transactions, &Transaction {
		Amount:  amount,
		Balance: a.ActualBalance,
		Time:    time.Now().Format("02/01/2006")})
}

func (a *Account) withdraw(amount int) {
	a.ActualBalance -= amount
	a.Transactions = append(a.Transactions, &Transaction {
		Amount:  -amount,
		Balance: a.ActualBalance,
		Time:    time.Now().Format("02/01/2006")})
}

func (a *Account) printStatement() {
	a.logger.Print("DATE | AMOUNT | BALANCE")
	for _, transaction := range a.Transactions {
		var result = fmt.Sprintf("%s | %v | %v", transaction.Time, transaction.Amount, transaction.Balance)
		a.logger.Print(result)
	}
}
