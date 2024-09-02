package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(5)

	go func (t *time.Ticker) {
		for tick := range t.C {
			fmt.Println("Tick at", tick)
			wg.Done()
		}
	}(ticker)

	wg.Wait()
	ticker.Stop()
	fmt.Println("Ticker stopped")
}