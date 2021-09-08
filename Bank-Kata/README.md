## Goal
Develop a program to manage the transactions of a bank account.

The transactions are: deposit money into the account, and withdraw from the account.

We need to be able to print into the console the result.

## Requirement
You cannot change the signature of the public interface (the class AccountService).

## Code

```go
type AccountService interface {
  deposit(amount int)

  withdraw(amount int)

  printStatement()
}
```


## Tip
Start writing the acceptance test but then move to unit tests, and at the end focus again on the acceptance test.

