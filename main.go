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

	a.Expense(&b, 10);
	print(a)
	print(b)

}

func print(v interface{}) {
	if bytes, e := json.MarshalIndent(v, "", "\t"); e == nil {
		fmt.Printf("%#s\n", string(bytes))
	}
}
