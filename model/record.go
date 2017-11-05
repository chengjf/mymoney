package model

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Record struct {
	Id        uint64     `db:"id"`
	Type      RecordType `db:"type"`
	AccountId uint64     `db:"account_id"`
	EntryId   uint64     `db:"entry_id"`
	Amount    float64    `db:"amount"`
	Datetime  time.Time  `db:"datetime"`
	Counter   string     `db:"counter"`
}
type RecordType int

const (
	Debit  RecordType = 1 // 借
	Credit RecordType = 2 // 贷
)

const (
	insertRecordSql = "INSERT INTO t_record(type, account_id, entry_id, amount, datetime, counter) VALUES (:type,:account_id,:entry_id,:amount,:datetime,:counter)"
)

type RecordDao struct {
	DB *sqlx.DB
}

func (dao *RecordDao) Insert(debitRecord Record, creditRecord Record) (err error) {
	tx := dao.DB.MustBegin()
	defer tx.Commit()
	_, err = tx.NamedExec(insertRecordSql, debitRecord)
	_, err = tx.NamedExec(insertRecordSql, creditRecord)
	return
}
