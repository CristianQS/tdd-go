package pkg

import (
	"bankkata/pkg/Mocks"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var builder = &TransactionBuilder{}


func Test_should_deposit_and_withdraw_an_amount_of_money_and_print_it_in_console(t *testing.T) {
	//GIVEN
	givenADepositMoney := 300
	givenAWithdrawMoney := 100
	ctrl := gomock.NewController(t)
	console := Mocks.NewMockLogger(ctrl)
	account := NewAccount(console)
	ActualTime := time.Now().Format("02/01/2006")
	//THEN
	console.EXPECT().Print(gomock.Eq("DATE | AMOUNT | BALANCE")).Times(1)
	console.EXPECT().Print(gomock.Eq(CreateTransactionInfo(ActualTime, givenADepositMoney,givenADepositMoney))).Times(1)
	console.EXPECT().Print(gomock.Eq(CreateTransactionInfo(ActualTime, -givenAWithdrawMoney,givenADepositMoney-givenAWithdrawMoney))).Times(1)

	//WHEN
	account.deposit(givenADepositMoney)
	account.withdraw(givenAWithdrawMoney)
	account.printStatement()
}

func Test_should_deposit_an_amount_of_money(t *testing.T) {
	givenADepositMoney := 300
	account := Account{}

	account.deposit(givenADepositMoney)

	transaction := builder.
		WithAmount(givenADepositMoney).
		WithBalance(givenADepositMoney).
		WithDate(time.Now().Format("02/01/2006")).
		Build()
	assert.Equal(t, transaction, account.GetTransactions()[0])
}

func Test_should_deposit_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenAnotherDepositMoney := 200
	account := Account{}

	account.deposit(givenADepositMoney)
	account.deposit(givenAnotherDepositMoney)
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
	
	assert.Equal(t, firstTransaction, account.GetTransactions()[0])
	assert.Equal(t, secondTransaction, account.GetTransactions()[1])
}


func Test_should_withdraw_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenWithdraw := 100
	account := Account{}

	account.deposit(givenADepositMoney)
	account.withdraw(givenWithdraw)

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

	assert.Equal(t, firstTransaction, account.GetTransactions()[0])
	assert.Equal(t, withdrawTransaction, account.GetTransactions()[1])
}


func Test_should_print_header_when_there_is_not_transaction(t *testing.T) {
	//GIVEN
	ctrl := gomock.NewController(t)
	console := Mocks.NewMockLogger(ctrl)
	account := NewAccount(console)

	//THEN
	console.EXPECT().Print(gomock.Eq("DATE | AMOUNT | BALANCE")).Times(1)

	//WHEN
	account.printStatement()
}

func CreateTransactionInfo(ActualTime string, givenADepositMoney, balance int) string {
	return fmt.Sprintf("%s | %v | %v", ActualTime, givenADepositMoney, balance)
}

