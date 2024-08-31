package main

import (
	"fmt"
	"sync"
)

type Counter struct {
    mu    sync.Mutex // Mutex to protect the counter
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()   // Lock the mutex before accessing the shared resource
    c.value++     // Safely increment the counter
    c.mu.Unlock() // Unlock the mutex after accessing the shared resource
}

func (c *Counter) Value() int {
    c.mu.Lock()   // Lock the mutex before accessing the shared resource
    defer c.mu.Unlock() // Ensure the mutex is unlocked even if function returns early
    return c.value
}

func main() {
    var wg sync.WaitGroup
    counter := &Counter{}

    // Spawn 1000 goroutines to increment the counter
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    // Wait for all goroutines to finish
    wg.Wait()

    // Safely read the counter value
    fmt.Printf("Final Counter Value: %d\n", counter.Value())
}
