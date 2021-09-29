package pkg

import (
	"bankkata/pkg/Infraestructure"
	"fmt"
)

type Account struct {
	logger           Infraestructure.Logger
	transactionStore Infraestructure.TransactionStore
}

func NewAccount(logger Infraestructure.Logger, transactionStore Infraestructure.TransactionStore) *Account {
	return &Account{logger: logger, transactionStore: transactionStore}
}

func (a *Account) deposit(amount int) {
	a.transactionStore.AddDeposit(amount)
}

func (a *Account) withdraw(amount int) {
	a.transactionStore.AddWithdraw(amount)
}

func (a *Account) printStatement() {
	a.logger.Print("DATE | AMOUNT | BALANCE")
	for _, transaction := range a.transactionStore.GetTransactions() {
		var result = fmt.Sprintf("%s | %v | %v", transaction.Date, transaction.Amount, transaction.Balance)
		a.logger.Print(result)
	}
}
