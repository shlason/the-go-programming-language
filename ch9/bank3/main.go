package main

import "sync"

var (
	mu      sync.Mutex
	balance int
)

func Deposite(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {

}
