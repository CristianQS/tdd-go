package Infraestructure

import (
	"bankkata/pkg/Model"
)

type InMemoryTransactionStore struct {
	Transactions  []*Model.Transaction
	timeProvider TimeProvider
	ActualBalance int
}

func NewInMemoryTransactionStore(timeProvider TimeProvider) *InMemoryTransactionStore {
	return &InMemoryTransactionStore{timeProvider: timeProvider}
}

func (t *InMemoryTransactionStore) AddDeposit(amount int) {
	t.ActualBalance += amount
	t.Transactions = append(t.Transactions,&Model.Transaction{
		Amount:  amount,
		Balance: t.ActualBalance,
		Date: t.timeProvider.GetTodayDate(),
	})
}

func (t *InMemoryTransactionStore) AddWithdraw(amount int) {
	t.ActualBalance -= amount
	t.Transactions = append(t.Transactions, &Model.Transaction{
		Amount:  -amount,
		Balance: t.ActualBalance,
		Date:    t.timeProvider.GetTodayDate(),
	})
}

func (t *InMemoryTransactionStore) GetTransactions() []*Model.Transaction {
	return t.Transactions
}