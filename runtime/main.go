package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
	fmt.Println("Num of CPUs", runtime.NumCPU())
	fmt.Println("Num of Goroutines", runtime.NumGoroutine())
	
	// runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	n := 20
	wg.Add(n)
	for i := 0; i<n;i++ {
		go func (index int) {
			fmt.Println("Foo: ", index)
			wg.Done()
		}(i)
		fmt.Println("Num of Goroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Num of Goroutines", runtime.NumGoroutine())
}