package bank2

var semaphore = make(chan struct{}, 1)
var balance int

func Deposit(amount int) {
	semaphore <- struct{}{} // acquire token
	balance += amount
	<-semaphore // release token
}

func Balance() int {
	semaphore <- struct{}{}
	b := balance
	<-semaphore
	return b
}
