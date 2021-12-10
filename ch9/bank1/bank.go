package main

import (
	"fmt"
	"time"
)

func main() {
	// Alice:
	go func() {
		Deposit(200) // A1
	}()

	// Bob:
	go Deposit(100) // B

	go func() {
		for {
			fmt.Println("=", Balance())
		}
	}()

	time.Sleep(time.Millisecond)
}

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func init() {
	go teller() // start the monitor goroutine
}

func teller() {
	var balance int // balance is confined to the teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func Balance() int {
	return <-balances
}
