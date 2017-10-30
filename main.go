package main

import (
	"fmt"
	"github.com/chengjf/mymoney/model"
	"time"
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
	fmt.Println(a)
	var b = model.Account{
		Name: "美丽源",
	}
	fmt.Println(b)

	var c = model.Category{
		Name: "晚饭",
	}
	fmt.Println(c)

	var t = model.Transaction{
		From:       a,
		To:         b,
		Amount:     22,
		CreateDate: time.Now(),
		UpdateDate: time.Now(),
	}

	fmt.Println(t)

}
