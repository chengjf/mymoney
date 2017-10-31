package model

type Report interface {
	QueryTransactionsByCategory(c Category) []Transaction
}
