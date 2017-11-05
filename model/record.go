package model

import "time"

type Record struct {
	Id       int
	Type     Type
	Account  int
	Entry    int
	Amount   float64
	Datetime time.Time
	Counter  string
}
type Type int

const (
	Debit  Type = 1 // 借
	Credit Type = 2 // 贷

)
