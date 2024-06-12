package main

import (
	"fmt"
	"sync"
)

type Loan struct {
	LoanNo string
}

func main() {
	loanGroups := make(map[int64][]*Loan)
	loanGroups[1] = []*Loan{{"1"}, {"2"}, {"3"}, {"4"}, {"5"}, {"6"}}
	loanGroups[2] = []*Loan{{"21"}, {"22"}, {"23"}, {"24"}, {"25"}, {"26"}}
	loanGroups[3] = []*Loan{{"31"}, {"32"}, {"33"}, {"34"}, {"35"}, {"36"}}
	loanGroups[4] = []*Loan{{"41"}, {"42"}, {"43"}, {"44"}, {"45"}, {"46"}}

	var wg sync.WaitGroup
	wg.Add(len(loanGroups))

	ch := make(chan string, 1000)

	for _, loanList := range loanGroups {
		loanListFinal := loanList
		println(&loanList)
		println(&loanListFinal)
		go func() {
			defer wg.Done()

			for _, loan := range loanListFinal {
				ch <- loan.LoanNo
			}
		}()
	}

	wg.Wait()
	close(ch)

	for e := range ch {
		fmt.Println(e)
	}
}
