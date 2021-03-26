package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepositingIncreasesBalance(t *testing.T) {
	var bank = newInMemoryBank()

	assert.NoError(t, Deposit(bank)("ABC", 100))
	assert.NoError(t, Deposit(bank)("ABC", 200))

	assert.Equal(t, 300, Balance(bank)("ABC"))
}

func TestNotAllowedToDepositNegativeAmounts(t *testing.T) {
	var bank = newInMemoryBank()

	assert.NoError(t, Deposit(bank)("ABC", 100))
	assert.Error(t, Deposit(bank)("ABC", -100))

	assert.Equal(t, 100, Balance(bank)("ABC"))
}

func TestWithdrawDecreasesBalance(t *testing.T) {
	var bank = newInMemoryBank()

	assert.NoError(t, Deposit(bank)("ABC", 400))
	assert.NoError(t, Withdraw(bank)("ABC", 100))

	assert.Equal(t, 300, Balance(bank)("ABC"))
}

func TestNotAllowedToWithdrawNegativeAmounts(t *testing.T) {
	var bank = newInMemoryBank()

	assert.NoError(t, Deposit(bank)("ABC", 100))
	assert.Error(t, Withdraw(bank)("ABC", -100))

	assert.Equal(t, 100, Balance(bank)("ABC"))
}

func TestAllowedToTransferFromOneAccountToAnother(t *testing.T) {
	var bank = newInMemoryBank()

	assert.NoError(t, Deposit(bank)("ABC", 100))
	assert.NoError(t, Deposit(bank)("DEF", 100))
	assert.NoError(t, Transfer(bank)("ABC", "DEF", 50))

	assert.Equal(t, 50, Balance(bank)("ABC"))
	assert.Equal(t, 150, Balance(bank)("DEF"))
}

type inMemoryBank struct {
	balances map[string]int
}

func newInMemoryBank() Bank {
	return inMemoryBank{
		balances: map[string]int{},
	}
}

func (i inMemoryBank) AddTransaction(account string, amount int) {
	i.balances[account] += amount
}

func (i inMemoryBank) GetTransactions(account string) []Transaction {
	return []Transaction{{Amount: i.balances[account]}}
}
