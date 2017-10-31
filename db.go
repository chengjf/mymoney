package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"time"
	"github.com/jmoiron/sqlx"
	"fmt"
)

type QA struct {
	Id       sql.NullInt64  `db:"id"`
	Uid      sql.NullInt64  `db:"uid"`
	QATime   time.Time      `db:"qa_time"`
	Question sql.NullString `db:"question"`
	Answer   sql.NullString `db:"answer"`
	Result   sql.NullInt64  `db:"result"`
}

func main() {
	testDB()
	//testDB2()
}

func testDB() {
	db, err := sqlx.Connect("mysql", "root:uxin.com@tcp(127.0.0.1:3306)/slive?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	var qa = QA{}
	err = db.Get(&qa, "SELECT id,uid,qa_time,question,answer,result FROM t_qa WHERE id = ?", 237)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(qa)
	if qa.Answer.Valid {
		fmt.Println(qa.Answer.String)
	}

	i := 0
	err = db.Get(&i, "select count(1) from t_qa where id < ?", 300)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(i)

	qas := []QA{}
	err = db.Select(&qas, "SELECT id,uid,qa_time,question,answer,result FROM t_qa WHERE id < ?", 300)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(len(qas), qas)
}

//func testDB2() {
//	db, err := sql.Open("mysql",
//		"root:uxin.com@tcp(127.0.0.1:3306)/slive?parseTime=true")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	stmt, err := db.Prepare("select id,uid,qa_time,question,answer,result from t_qa where id = ?")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stmt.Close()
//
//	var (
//		id       int64
//		uid      int64
//		qa_time  time.Time
//		question string
//		answer   string
//		result   int64
//	)
//	err = stmt.QueryRow(237).Scan(&id, &uid, &qa_time, &question, &answer, &result)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(id)
//	log.Println(uid)
//	log.Println(qa_time)
//	log.Println(question)
//	log.Println(answer)
//	log.Println(result)
//}
