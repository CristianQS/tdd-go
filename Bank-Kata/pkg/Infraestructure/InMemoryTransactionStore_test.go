package Infraestructure

import (
	"bankkata/pkg/Builders"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var builder = &Builders.TransactionBuilder{}

func Test_should_deposit_an_amount_of_money(t *testing.T) {
	givenADepositMoney := 300
	memoryTransactionStore := InMemoryTransactionStore{}

	memoryTransactionStore.AddDeposit(givenADepositMoney)

	expectedTransaction := builder.
		WithAmount(givenADepositMoney).
		WithBalance(givenADepositMoney).
		WithDate(time.Now().Format("02/01/2006")).
		Build()
	assert.Equal(t, expectedTransaction, memoryTransactionStore.GetTransactions()[0])
}

func Test_should_deposit_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenAnotherDepositMoney := 200
	memoryTransactionStore := InMemoryTransactionStore{}

	memoryTransactionStore.AddDeposit(givenADepositMoney)
	memoryTransactionStore.AddDeposit(givenAnotherDepositMoney)
	firstTransaction := builder.
		WithAmount(givenADepositMoney).
		WithBalance(givenADepositMoney).
		WithDate(time.Now().Format("02/01/2006")).
		Build()

	secondTransaction := builder.
		WithAmount(givenAnotherDepositMoney).
		WithBalance(givenADepositMoney+givenAnotherDepositMoney).
		WithDate(time.Now().Format("02/01/2006")).
		Build()

	assert.Equal(t, firstTransaction, memoryTransactionStore.GetTransactions()[0])
	assert.Equal(t, secondTransaction, memoryTransactionStore.GetTransactions()[1])
}


func Test_should_withdraw_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenWithdraw := 100
	memoryTransactionStore := InMemoryTransactionStore{}

	memoryTransactionStore.AddDeposit(givenADepositMoney)
	memoryTransactionStore.AddWithdraw(givenWithdraw)

	firstTransaction := builder.
		WithAmount(givenADepositMoney).
		WithBalance(givenADepositMoney).
		WithDate(time.Now().Format("02/01/2006")).
		Build()

	withdrawTransaction := builder.
		WithAmount(-givenWithdraw).
		WithBalance( givenADepositMoney - givenWithdraw).
		WithDate(time.Now().Format("02/01/2006")).
		Build()

	assert.Equal(t, firstTransaction, memoryTransactionStore.GetTransactions()[0])
	assert.Equal(t, withdrawTransaction, memoryTransactionStore.GetTransactions()[1])
}