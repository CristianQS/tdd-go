package Infraestructure

import (
	"bankkata/pkg/Model"
)

type TransactionStore interface {
	AddDeposit(amount int)
	AddWithdraw(amount int)
	GetTransactions() []*Model.Transaction
}

