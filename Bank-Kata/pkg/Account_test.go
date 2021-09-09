package pkg

import (
	"bankkata/pkg/Mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_deposit_and_withdraw_an_amount_of_money_and_print_it_in_console(t *testing.T) {
	givenADepositMoney := 300
	givenAWithdrawMoney := 100
	account := Account{}
	ctrl := gomock.NewController(t)
	console := Mocks.NewMockLogger(ctrl)

	account.deposit(givenADepositMoney)
	account.withdraw(givenAWithdrawMoney)
	account.printStatement()

	console.EXPECT().Print(gomock.Eq("DATE | AMOUNT | BALANCE")).Times(1)
	console.EXPECT().Print(gomock.Eq("08/09/2021 | 300 | 300")).Times(1)
	console.EXPECT().Print(gomock.Eq("08/09/2021 | -100 | 200")).Times(1)
}

func Test_should_deposit_an_amount_of_money(t *testing.T) {
	givenADepositMoney := 300
	account := Account{}

	account.deposit(givenADepositMoney)

	transaction := &Transaction{
		amount:          givenADepositMoney,
		balance:         givenADepositMoney,
		timeTransaction: time.Now().Format("02-01-2006")}
	assert.Equal(t, transaction, account.GetTransactions()[0])
}

func Test_should_deposit_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenAnotherDepositMoney := 200
	account := Account{}

	account.deposit(givenADepositMoney)
	account.deposit(givenAnotherDepositMoney)

	firstTransaction := &Transaction{
		amount:          givenADepositMoney,
		balance:         givenADepositMoney,
		timeTransaction: time.Now().Format("02-01-2006")}

	secondTransaction := &Transaction{
		amount:          givenAnotherDepositMoney,
		balance:         givenADepositMoney + givenAnotherDepositMoney,
		timeTransaction: time.Now().Format("02-01-2006")}
	assert.Equal(t, firstTransaction, account.GetTransactions()[0])
	assert.Equal(t, secondTransaction, account.GetTransactions()[1])
}

