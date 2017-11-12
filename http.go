package main

import (
	"encoding/json"
	"fmt"
	"github.com/chengjf/mymoney/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/olebedev/config"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
	"github.com/gjvnq/go-logger"
	"os"
)

var log, err = logger.New("test", 1, os.Stdout)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Fprint(w, "Welcome!\n")
	Show(w, r)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
func Show(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("view/index.html")
	if err != nil {
		log.ErrorF("error:%v", err)
	}
	w.Write(b)
}

func Static(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}

func GetAllAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	accountDao := model.AccountDao{DB: DB}

	cashAccount, err := accountDao.QueryAllAccount()
	if err != nil {
		log.ErrorF("queryAllAccount error: %v", err)
	}
	if bytes, e := json.MarshalIndent(cashAccount, "", "\t"); e == nil {
		w.Write(bytes)
	}
}

func GetExpenseEntries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info("GetExpenseEntries start")
	entryDao := model.EntryDao{DB: DB}

	parentEntry := ps.ByName("parent")
	if parentEntry != "" {
		log.InfoF("GetExpenseEntries, parentEntry is %v", parentEntry)

		if parseUint, e := strconv.ParseUint(parentEntry, 10, 64); e == nil {

			entries, err := entryDao.QueryChildEntriesById(parseUint);
			if err != nil {
				log.ErrorF("GetExpenseEntries, error: %v", err)
			}
			if bytes, e := json.MarshalIndent(entries, "", "\t"); e == nil {
				w.Write(bytes)
			}
		} else {
			log.ErrorF("GetExpenseEntries, parentEntry parse error %v", e)
		}
	} else {
		entries, err := entryDao.QueryEntriesByLevel(1);
		if err != nil {
			log.ErrorF("GetExpenseEntries error: %v", err)
		}
		if bytes, e := json.MarshalIndent(entries, "", "\t"); e == nil {
			w.Write(bytes)
		}
	}
}

func GetAllEntries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Info("GetAllEntries start")
	entryDao := model.EntryDao{DB: DB}

	entries, err := entryDao.QueryAllEntries()
	if err != nil {
		log.ErrorF("GetAllEntries error: %v", err)
	}
	if bytes, e := json.MarshalIndent(entries, "", "\t"); e == nil {
		w.Write(bytes)
	}

}

type recordRequest struct {
	DebitAccount  uint64
	DebitEntry    uint64
	CreditAccount uint64
	CreditEntry   uint64
	Amount        float64
	Datetime      MyTime
	Counter       string
}

type MyTime struct {
	t time.Time
}

func (t *MyTime) UnmarshalJSON(s []byte) error {
	var (
		year int
		mon  int
		mday int
		hour int
		min  int
		sec  int
	)
	if len(s) <= 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return fmt.Errorf("invalid time: %s", s)
	}
	var str = string(s[1: len(s)-1])
	if n, err := fmt.Sscanf(str, "%d-%02d-%02d %02d:%02d:%02d", &year, &mon, &mday, &hour, &min, &sec); err != nil {
		return fmt.Errorf("invalid string(%s): %s", err.Error(), s)
	} else if n != 6 {
		return fmt.Errorf("invalid time: %s", s)
	}

	t.t = time.Date(year, time.Month(mon), mday, hour, min, sec, 0, time.UTC)
	return nil
}
func CreateRecord(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, _ := ioutil.ReadAll(r.Body)
	log.InfoF("CreateRecord requst data is: %v", string(body))
	request := recordRequest{}
	// {"debitAccount":3,"debitEntry":21,"creditAccount":1,"creditEntry":2,"amount":"11","datetime":"2017-11-05","counter":"ddd"}
	err := json.Unmarshal(body, &request)
	if err != nil {
		log.InfoF("CreateRecord unmarshal error: %v", err)
	}

	var record = model.Record{
		Type:      model.Debit,
		AccountId: request.DebitAccount,
		EntryId:   request.DebitEntry,
		Amount:    request.Amount,
		Counter:   request.Counter,
		Datetime:  request.Datetime.t,
	}

	record2 := model.Record{
		Type:      model.Credit,
		AccountId: request.CreditAccount,
		EntryId:   request.CreditEntry,
		Amount:    request.Amount,
		Counter:   request.Counter,
		Datetime:  request.Datetime.t,
	}

	recordDao := model.RecordDao{
		DB: DB,
	}

	err = recordDao.Insert(record, record2)
	if err != nil {
		log.Error(err)
	}

	//entryDao := model.EntryDao{DB: DB}
	//
	//entries, err := entryDao.QueryAllEntries()
	//if err != nil {
	//	log.Fatalln("GetExpenseEntries error: ", err)
	//}
	//if bytes, e := json.MarshalIndent(entries, "", "\t"); e == nil {
	//	w.Write(bytes)
	//}
}

var DB *sqlx.DB

func main() {
	DB = initDb()
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/accounts", GetAllAccounts)
	router.GET("/entries/:parent", GetExpenseEntries)
	router.GET("/entries", GetAllEntries)
	router.POST("/createRecord", CreateRecord)
	router.ServeFiles("/static/*filepath", http.Dir("static/"))
	fmt.Println("Server start at 8088")
	log.Fatal(http.ListenAndServe(":8088", router))
}

func initDb() *sqlx.DB {
	bytes, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Error(err)
	}
	cfg, err := config.ParseYaml(string(bytes))
	if err != nil {
		log.Error(err)
	}
	print(cfg)
	host, err := cfg.String("development.database.host")
	port, err := cfg.String("development.database.port")
	schema, err := cfg.String("development.database.schema")
	username, err := cfg.String("development.database.username")
	password, err := cfg.String("development.database.password")
	if err != nil {
		log.Error(err)
	}

	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4", username, password, host, port, schema)
	log.InfoF("Database usr is: %v", dataSourceName)

	db, err := sqlx.Connect("mysql", dataSourceName) //"mymoney:mymoney@tcp(127.0.0.1:3306)/mymoney?parseTime=true")
	if err != nil {
		log.Error(err)
	}
	return db
}
