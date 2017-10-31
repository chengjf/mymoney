package model

type Pool interface {
	Minus(i int64)
	Add(i int64)
	GetTransaction() []Transaction
}
