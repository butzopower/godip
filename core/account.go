package core

import "errors"

type Transaction struct {
	Amount int
}

type Bank interface {
	AddTransaction(account string, amount int)
	GetTransactions(account string) []Transaction
}

func Deposit(b Bank) func(string, int) error {
	return func(account string, amount int) error {
		if amount < 0 {
			return errors.New("can't deposit negative Amount")
		}

		b.AddTransaction(account, amount)
		return nil
	}
}

func Withdraw(b Bank) func(string, int) error {
	return func(account string, amount int) error {
		if amount < 0 {
			return errors.New("can't withdraw negative Amount")
		}

		b.AddTransaction(account, -amount)
		return nil
	}
}


func Balance(b Bank) func(string) int {
	return func(account string) int {
		balance := 0

		for _, t := range b.GetTransactions(account) {
			balance += t.Amount
		}

		return balance
	}
}

func Transfer(b Bank) func(string, string, int) error {
	return func(payer string, payee string, amount int) error {
		b.AddTransaction(payer, -amount)
		b.AddTransaction(payee, amount)
		return nil
	}
}