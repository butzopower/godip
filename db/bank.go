package db

import (
	"encoding/json"
	"example/dip/core"
	"github.com/xujiajun/nutsdb"
	"io/ioutil"
)

type dbBank struct {
	db *nutsdb.DB
}

var noOp = func() error { return nil }

func NewDbBank() (core.Bank, func() error, error) {
	tempDir, err := ioutil.TempDir("", "")

	if err != nil {
		return dbBank {}, noOp, err
	}

	opt := nutsdb.DefaultOptions
	opt.Dir = tempDir

	db, err := nutsdb.Open(opt)

	if err != nil {
		return dbBank {}, noOp, err
	}

	return dbBank { db: db }, db.Close, nil
}

func (d dbBank) AddTransaction(account string, amount int) {
	d.db.Update(func(tx *nutsdb.Tx) error {
		transaction, err := json.Marshal(dbTransaction{Amount: amount})

		if err != nil {
			return err
		}

		return tx.RPush("accounts", []byte(account), transaction)
	})
}

func (d dbBank) GetTransactions(account string) []core.Transaction {
	transactions := make([]core.Transaction, 0)

	d.db.View(func(tx *nutsdb.Tx) error {
		items, err := tx.LRange("accounts", []byte(account), 0, -1)

		if err != nil {
			return err
		}

		for _, item := range items {
			var transaction = dbTransaction{}

			err := json.Unmarshal(item, &transaction)

			if err != nil {
				return err
			}

			transactions = append(transactions, core.Transaction{Amount: transaction.Amount})
		}

		return nil
	})

	return transactions
}

type dbTransaction struct {
	Amount int `json:"amount"`
}