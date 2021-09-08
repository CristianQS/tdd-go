package pkg

import (
	"testing"
)

func Test_should_deposit_and_withdraw_an_amount_of_money_and_print_it_in_console(t *testing.T) {
	givenADepositMoney := 300
	givenAWithdrawMoney := 100

	AccountService{}.Deposit(givenADepositMoney)
	AccountService{}.Withdraw(givenAWithdrawMoney)
	AccountService{}.PrintStatement()

	console.EXPECT().print("DATE | AMOUNT | BALANCE").Times(1)
	console.EXPECT().print("08/09/2021 | 300 | 300").Times(1)
	console.EXPECT().print("08/09/2021 | -100 | 200").Times(1)

}