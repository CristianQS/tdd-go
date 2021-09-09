package pkg

import (
	"bankkata/pkg/Mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_should_deposit_and_withdraw_an_amount_of_money_and_print_it_in_console(t *testing.T) {
	//GIVEN
	givenADepositMoney := 300
	givenAWithdrawMoney := 100
	ctrl := gomock.NewController(t)
	console := Mocks.NewMockLogger(ctrl)
	account := NewAccount(console)

	//THEN
	console.EXPECT().Print(gomock.Eq("DATE | AMOUNT | BALANCE")).Times(1)
	console.EXPECT().Print(gomock.Eq("09/09/2021 | 300 | 300")).Times(1)
	console.EXPECT().Print(gomock.Eq("09/09/2021 | -100 | 200")).Times(1)

	//WHEN
	account.deposit(givenADepositMoney)
	account.withdraw(givenAWithdrawMoney)
	account.printStatement()
}

func Test_should_deposit_an_amount_of_money(t *testing.T) {
	givenADepositMoney := 300
	account := Account{}

	account.deposit(givenADepositMoney)

	transaction := &Transaction{
		Amount:  givenADepositMoney,
		Balance: givenADepositMoney,
		Time:    time.Now().Format("02/01/2006")}
	assert.Equal(t, transaction, account.GetTransactions()[0])
}

func Test_should_deposit_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenAnotherDepositMoney := 200
	account := Account{}

	account.deposit(givenADepositMoney)
	account.deposit(givenAnotherDepositMoney)

	firstTransaction := &Transaction{
		Amount:  givenADepositMoney,
		Balance: givenADepositMoney,
		Time:    time.Now().Format("02/01/2006")}

	secondTransaction := &Transaction{
		Amount:  givenAnotherDepositMoney,
		Balance: givenADepositMoney + givenAnotherDepositMoney,
		Time:    time.Now().Format("02/01/2006")}
	assert.Equal(t, firstTransaction, account.GetTransactions()[0])
	assert.Equal(t, secondTransaction, account.GetTransactions()[1])
}


func Test_should_withdraw_various_amounts_of_money_on_the_account(t *testing.T) {
	givenADepositMoney := 300
	givenWithdraw := 100
	account := Account{}

	account.deposit(givenADepositMoney)
	account.withdraw(givenWithdraw)

	firstTransaction := &Transaction{
		Amount:  givenADepositMoney,
		Balance: givenADepositMoney,
		Time:    time.Now().Format("02/01/2006")}

	withdrawTransaction := &Transaction{
		Amount:  -givenWithdraw,
		Balance: givenADepositMoney - givenWithdraw,
		Time:    time.Now().Format("02/01/2006")}
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


