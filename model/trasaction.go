package model

import "time"

type Transaction struct {
	From Account
	To Account
	Amount int32
	CreateDate time.Time
	UpdateDate time.Time
}
