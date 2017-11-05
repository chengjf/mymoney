package main

import (
	"encoding/json"
	"fmt"
	"github.com/chengjf/mymoney/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/olebedev/config"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	fmt.Println("Hello, this is my first go application.")
	test()
}

func test() {
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	cfg, err := config.ParseYaml(string(bytes))
	if err != nil {
		log.Fatalln(err)
	}
	print(cfg)
	host, err := cfg.String("development.database.host")
	port, err := cfg.String("development.database.port")
	schema, err := cfg.String("development.database.schema")
	username, err := cfg.String("development.database.username")
	password, err := cfg.String("development.database.password")
	if err != nil {
		log.Fatalln(err)
	}

	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, schema)
	log.Println("Database usr is:", dataSourceName)

	db, err := sqlx.Connect("mysql", dataSourceName) //"mymoney:mymoney@tcp(127.0.0.1:3306)/mymoney?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	var cashAccount = model.Account{}
	err = db.Get(&cashAccount, "SELECT id,name,balance,entry_id AS entry  FROM t_account where name = '现金账户1'")
	if err != nil {
		log.Fatalln(err)
	}
	print(cashAccount)

	var expenseAccount = model.Account{}
	err = db.Get(&expenseAccount, "SELECT id,name,balance,entry_id AS entry  FROM t_account where name = '虚拟支出账户'")
	if err != nil {
		log.Fatalln(err)
	}
	print(expenseAccount)

	var launchEntry = model.Entry{}
	err = db.Get(&launchEntry, "SELECT id, name, level, parent_lvl AS parentlvl FROM t_entry WHERE name = '午餐' ")
	if err != nil {
		log.Fatalln(err)
	}
	print(launchEntry)

	var entry = model.Entry{}
	err = db.Get(&entry, "SELECT id, name, level, parent_lvl AS parentlvl FROM t_entry WHERE name = '现金' ")
	if err != nil {
		log.Fatalln(err)
	}
	print(entry)

	var record = model.Record{
		Type:     model.Debit,
		Account:  expenseAccount.Id,
		Entry:    entry.Id,
		Amount:   14.90,
		Counter:  "kfc",
		Datetime: time.Now(),
	}
	print(record)
	tx := db.MustBegin()
	result, err := tx.Exec("INSERT INTO t_record(type, account_id, entry_id, amount, datetime, counter) VALUES (?,?,?,?,?,?)", record.Type, record.Account, record.Entry, record.Amount, record.Datetime, record.Counter)
	if err != nil {
		log.Fatalln(err)
	}
	print(result)

	record = model.Record{
		Type:     model.Credit,
		Account:  cashAccount.Id,
		Entry:    launchEntry.Id,
		Amount:   14.90,
		Counter:  "kfc",
		Datetime: time.Now(),
	}
	print(record)

	result, err = tx.NamedExec("INSERT INTO t_record(type, account_id, entry_id, amount, datetime, counter) VALUES (:type,:account,:entry,:amount,:datetime,:counter)", &record)
	if err != nil {
		log.Fatalln("xx", err)
	}
	print(result)

	tx.Commit()
}

func print(v interface{}) {
	if bytes, e := json.MarshalIndent(v, "", "\t"); e == nil {
		fmt.Printf("%#s\n", string(bytes))
	}
}
