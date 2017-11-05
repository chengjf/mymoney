package model

import "github.com/jmoiron/sqlx"

type Account struct {
	Id      uint64  `db:"id"`
	Name    string  `db:"name"`
	EntryId uint64  `db:"entry_id"`
	Balance float64 `db:"balance"`
}

type AccountDao struct {
	DB *sqlx.DB
}

const (
	queryAllAccount    = "select id,name,entry_id,balance from t_account"
	queryAccountByName = "select id,name,entry_id,balance from t_account where name=?"
)

func (dao *AccountDao) QueryAllAccount() (accounts []Account, err error) {
	err = dao.DB.Select(&accounts, queryAllAccount)
	return
}

func (dao *AccountDao) QueryAccountByName(name string) (account Account, err error) {
	err = dao.DB.Get(&account, queryAccountByName, name)
	return
}
