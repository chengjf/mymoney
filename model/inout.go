package model

import "sync"

type InOut struct {
	Name    string
	Balance int64
	sync.Mutex
}

func (a *InOut) Minus(i int64) {
	a.Lock()
	defer a.Unlock()
	a.Balance -= i
}

func (a *InOut) Add(i int64) {
	a.Lock()
	defer a.Unlock()
	a.Balance += i
}
