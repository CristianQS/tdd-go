package pkg

type TransactionBuilder struct {
	Balance int
	Amount  int
	Date string
}

func (t *TransactionBuilder) WithAmount(amount int) *TransactionBuilder {
	t.Amount = amount
	return t
}

func (t *TransactionBuilder) WithBalance(balance int) *TransactionBuilder {
	t.Balance = balance
	return t
}

func (t *TransactionBuilder)WithDate(date string) *TransactionBuilder {
	t.Date = date
	return t
}

func (t *TransactionBuilder) Build() *Transaction {
	return &Transaction{
		Amount: t.Amount,
		Balance: t.Balance,
		Date: t.Date,
	}
}
