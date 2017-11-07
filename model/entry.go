package model

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Entry struct {
	Id        uint64        `db:"id"`
	Name      string        `db:"name"`
	Level     uint64        `db:"level"`
	ParentLvl sql.NullInt64 `db:"parent_lvl"`
}

type EntryDao struct {
	DB *sqlx.DB
}

type EntryType string

const (
	Asset     EntryType = "资产类"
	Liability EntryType = "负债类"
	Income    EntryType = "收入类"
	Expense   EntryType = "支出类"
)

const (
	queryAllEntries       = "select id,name,level,parent_lvl from t_entry"
	queryEntriesByLevel   = "select id,name,level,parent_lvl from t_entry where level=?"
	queryEntryById        = "select id,name,level,parent_lvl from t_entry where id=?"
	queryEntryByName      = "select id,name,level,parent_lvl from t_entry where name=?"
	queryChildEntriesById = "select id,name,level,parent_lvl from t_entry where parent_lvl=?"
)

func (dao *EntryDao) QueryAllEntries() (entries []Entry, err error) {
	err = dao.DB.Select(&entries, queryAllEntries)
	return
}

func (dao *EntryDao) QueryEntryByName(name string) (entry Entry, err error) {
	dao.DB.Get(&entry, queryEntryByName, name)
	return
}

func (dao *EntryDao) QueryEntryById(id uint64) (entry Entry, err error) {
	dao.DB.Get(&entry, queryEntryById, id)
	return
}

func (dao *EntryDao) QueryEntriesByLevel(id uint64) (entries []Entry, err error) {
	err = dao.DB.Select(&entries, queryEntriesByLevel, id)
	return
}

func (dao *EntryDao) QueryChildEntriesById(id uint64) (entries []Entry, err error) {
	err = dao.DB.Select(&entries, queryChildEntriesById, id)
	return
}