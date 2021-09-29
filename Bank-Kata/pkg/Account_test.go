package pkg

import (
	"bankkata/pkg/Builders"
	"bankkata/pkg/Mocks"
	model "bankkata/pkg/Model"
	"fmt"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

var builder = &Builders.TransactionBuilder{}

func Test_should_deposit_and_withdraw_an_amount_of_money_and_print_it_in_console(t *testing.T) {
	//GIVEN
	givenADepositMoney := 300
	givenAWithdrawMoney := 100
	ctrl := gomock.NewController(t)
	console := Mocks.NewMockLogger(ctrl)
	transactionStore := Mocks.NewMockTransactionStore(ctrl)
	account := NewAccount(console,transactionStore)
	ActualTime := time.Now().Format("02/01/2006")
	transaction := builder.
		WithAmount(givenADepositMoney).
		WithBalance(givenADepositMoney).
		WithDate(ActualTime).
		Build()
	secondTransaction := builder.
		WithAmount(-givenAWithdrawMoney).
		WithBalance(givenADepositMoney-givenAWithdrawMoney).
		WithDate(ActualTime).
		Build()
	//THEN
	transactionStore.EXPECT().AddDeposit(gomock.Eq(givenADepositMoney)).Times(1)
	transactionStore.EXPECT().AddWithdraw(gomock.Eq(givenAWithdrawMoney)).Times(1)
	transactionStore.EXPECT().GetTransactions().Return([]*model.Transaction{transaction,secondTransaction}).Times(1)
	console.EXPECT().Print(gomock.Eq("DATE | AMOUNT | BALANCE")).Times(1)
	console.EXPECT().Print(gomock.Eq(CreateTransactionInfo(ActualTime, givenADepositMoney,givenADepositMoney))).Times(1)
	console.EXPECT().Print(gomock.Eq(CreateTransactionInfo(ActualTime, -givenAWithdrawMoney,givenADepositMoney-givenAWithdrawMoney))).Times(1)

	//WHEN
	account.deposit(givenADepositMoney)
	account.withdraw(givenAWithdrawMoney)
	account.printStatement()
}

func Test_should_print_header_when_there_is_not_transaction(t *testing.T) {
	//GIVEN
	ctrl := gomock.NewController(t)
	console := Mocks.NewMockLogger(ctrl)
	transactionStore := Mocks.NewMockTransactionStore(ctrl)
	account := NewAccount(console,transactionStore)

	//THEN
	transactionStore.EXPECT().GetTransactions().Return([]*model.Transaction{}).Times(1)
	console.EXPECT().Print(gomock.Eq("DATE | AMOUNT | BALANCE")).Times(1)

	//WHEN
	account.printStatement()
}

func CreateTransactionInfo(ActualTime string, givenADepositMoney, balance int) string {
	return fmt.Sprintf("%s | %v | %v", ActualTime, givenADepositMoney, balance)
}

