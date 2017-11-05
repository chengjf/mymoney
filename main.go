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

	accountDao := model.AccountDao{DB: db}

	cashAccount, err := accountDao.QueryAccountByName("现金账户1")
	if err != nil {
		log.Fatalln("query cash account", err)
	}
	print(cashAccount)

	expenseAccount, err := accountDao.QueryAccountByName("虚拟支出账户")
	if err != nil {
		log.Fatalln("query expense account", err)
	}
	print(expenseAccount)

	entryDao := model.EntryDao{db}
	launchEntry, err := entryDao.QueryEntryByName("午餐")

	if err != nil {
		log.Fatalln(err)
	}
	print(launchEntry)

	entry, err := entryDao.QueryEntryByName("现金")
	if err != nil {
		log.Fatalln(err)
	}
	print(entry)

	var record = model.Record{
		Type:      model.Debit,
		AccountId: expenseAccount.Id,
		EntryId:   entry.Id,
		Amount:    14.90,
		Counter:   "kfc",
		Datetime:  time.Now(),
	}

	record2 := model.Record{
		Type:      model.Credit,
		AccountId: cashAccount.Id,
		EntryId:   launchEntry.Id,
		Amount:    14.90,
		Counter:   "kfc",
		Datetime:  time.Now(),
	}

	recordDao := model.RecordDao{
		DB: db,
	}

	err = recordDao.Insert(record, record2)
	if err != nil {
		log.Fatalln(err)
	}

}

func print(v interface{}) {
	if bytes, e := json.MarshalIndent(v, "", "\t"); e == nil {
		fmt.Printf("%#s\n", string(bytes))
	}
}
