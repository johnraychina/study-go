package main

import (
	"os"
	"os/signal"
	"time"
)

func StartTickerWithTimer() {

	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			t := <-ticker.C
			println("tick:", t.String())
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
