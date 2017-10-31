package main

import (
	"fmt"
	"github.com/chengjf/mymoney/model"
	"time"
	"encoding/json"
)

func main() {
	fmt.Println("Hello, this is my first go application.")
	test()
}

func test() {
	var a = model.Account{
		Name:    "Cash",
		Balance: 100,
	}
	print(a)
	var b = model.Account{
		Name: "美丽源",
	}
	print(b)

	var c = model.Category{
		Name: "晚饭",
	}
	print(c)

	var t = model.Transaction{
		From:       a,
		To:         b,
		Amount:     22,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}

	print(t)

}

func print(v interface{}) {
	if bytes, e := json.MarshalIndent(v,"","\t"); e == nil {
		fmt.Printf("%#s\n", string(bytes))
	}
}
