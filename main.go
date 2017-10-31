package main

import (
	"fmt"
	"github.com/chengjf/mymoney/model"
	"encoding/json"
)

func main() {
	fmt.Println("Hello, this is my first go application.")
	test()
}

func test() {
	var a = model.Account{
		model.InOut{
			Name:    "Cash",
			Balance: 100,
		},
	}
	print(a)
	var b = model.Account{
		model.InOut{Name: "美丽源",},
	}
	print(b)

	var c = model.Category{
		Name: "晚饭",
	}
	print(c)

	var s = model.Category{
		Name: "工资",
	}
	print(s)

	fmt.Println("--------Expense---------")
	a.Expense(&b, c,10);
	a.Expense(&b, c,20);
	a.Expense(&b, c,30);
	print(a)
	print(b)
	print(a.GetTransaction())

	fmt.Println("--------Income---------")
	a.Income(&b, s, 1000)
	print(a)
	print(b)
	print(a.GetTransaction())

	fmt.Println("--------Query---------")
	rs := a.QueryTransactionsByCategory(c)
	print(rs)

}

func print(v interface{}) {
	if bytes, e := json.MarshalIndent(v, "", "\t"); e == nil {
		fmt.Printf("%#s\n", string(bytes))
	}
}
