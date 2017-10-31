package model

import "time"

type Transaction struct {
	From Pool
	To Pool
	Type string
	Category Category
	Amount int64
	CreateDate time.Time
	UpdateDate time.Time
}
