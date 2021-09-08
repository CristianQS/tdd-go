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

	assert.Equal(t, givenADepositMoney, account.GetBalance())
	assert.Equal(t, givenADepositMoney, account.GetAmount())
	assert.Equal(t, time.Now().Format("DD/MM/YYYY"), account.GetDate())
}

