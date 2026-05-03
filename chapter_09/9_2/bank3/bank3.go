package bank3

import "sync"

var (
	mu sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

func Withdraw(amount int) bool {
	Deposit(amount)
	if Balance() < 0 {
		Deposit(amount)
		return false
	}
	return true
}