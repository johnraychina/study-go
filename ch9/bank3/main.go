package bank3

import "sync"

//var mu sync.Mutex // guards balance
var mu sync.RWMutex
var balance int

func Deposit(amount int) {
	mu.Lock()
	deposit(amount)
	defer mu.Unlock()
}

func deposit(amount int) {
	balance += amount
}

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}

	return true
}

func Balance() int {
	mu.RLock()
	defer mu.RUnlock()
	return balance
}
