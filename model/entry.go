package model

import "github.com/jmoiron/sqlx"

type Entry struct {
	Id        uint64 `db:"id"`
	Name      string `db:"name"`
	Level     uint64 `db:"level"`
	ParentLvl uint64 `db:"parent_lvl"`
}

type EntryDao struct {
	DB *sqlx.DB
}

const (
	queryAllEntries  = "select id,name,level,parent_lvl from t_entry"
	queryEntryByName = "select id,name,level,parent_lvl from t_entry where name=?"
)

func (dao *EntryDao) QueryAllEntries() (entries []Entry, err error) {
	err = dao.DB.Select(&entries, queryAllEntries)
	return
}

func (dao *EntryDao) QueryEntryByName(name string) (entry Entry, err error) {
	dao.DB.Get(&entry, queryEntryByName, name)
	return
}
