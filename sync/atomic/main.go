package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
    var wg sync.WaitGroup
    var counter int64

    // Spawn 1000 goroutines to increment the counter
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
			fmt.Println("Goroutines", runtime.NumGoroutine())
			runtime.Gosched()
            atomic.AddInt64(&counter, 1)
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()

    // Safely read the counter value
    fmt.Printf("Final Counter Value: %d\n", counter)
}
