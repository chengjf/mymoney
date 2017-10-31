package model

import (
	"sync"
	"time"
)

type InOut struct {
	Name         string
	Balance      int64
	sync.Mutex
	transactions []Transaction
}

func (io *InOut) Minus(i int64) {
	io.Lock()
	defer io.Unlock()
	io.Balance -= i
}

func (io *InOut) Add(i int64) {
	io.Lock()
	defer io.Unlock()
	io.Balance += i
}

func (a *InOut) Expense(dest Pool, c Category, i int64) {
	a.Minus(i)
	dest.Add(i)
	t := Transaction{
		From:       a,
		To:         dest,
		Type:       "expense",
		Category:   c,
		Amount:     i,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}
	a.transactions = append(a.transactions, t)
}

func (a *InOut) Income(from Pool, c Category, i int64) {
	a.Add(i)
	from.Minus(i)
	t := Transaction{
		From:       a,
		To:         from,
		Type:       "income",
		Category:   c,
		Amount:     i,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}
	a.transactions = append(a.transactions, t)
}

func (a *InOut) GetTransaction() []Transaction {
	return a.transactions
}
