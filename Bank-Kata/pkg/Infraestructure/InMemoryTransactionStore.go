package Infraestructure

import (
	"bankkata/pkg/Model"
	"time"
)

type InMemoryTransactionStore struct {
	Transactions  []*Model.Transaction
	ActualBalance int
}

func (t *InMemoryTransactionStore) AddDeposit(amount int) {
	t.ActualBalance += amount
	t.Transactions = append(t.Transactions,&Model.Transaction{
		Amount:  amount,
		Balance: t.ActualBalance,
		Date:    time.Now().Format("02/01/2006")})
}

func (t *InMemoryTransactionStore) AddWithdraw(amount int) {
	t.ActualBalance -= amount
	t.Transactions = append(t.Transactions, &Model.Transaction{
		Amount:  -amount,
		Balance: t.ActualBalance,
		Date:    time.Now().Format("02/01/2006")})
}

func (t *InMemoryTransactionStore) GetTransactions() []*Model.Transaction {
	return t.Transactions
}