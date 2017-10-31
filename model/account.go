package model

type Account struct {
	InOut
}

func (a *Account) Expense(o Pool, i int64) {
	a.Minus(i)
	o.Add(i)
}

func (a *Account) Income(o Pool, i int64) {
	o.Minus(i)
	a.Add(i)
}

func (a *Account) Transfer(o Pool, i int64) {
	a.Minus(i)
	o.Add(i)
}
