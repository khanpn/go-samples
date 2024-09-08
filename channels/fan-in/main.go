package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := fanIn(
		startJob("Job 1", 1*time.Second),
		startJob("Job 2", 3*time.Second),
		startJob("Job 3", 2*time.Second),
		startJob("Job 4", 2*time.Second),
		startJob("Job 5", 1*time.Second),
		startJob("Job 6", 3*time.Second),
		startJob("Job 7", 2*time.Second),
		startJob("Job 8", 1*time.Second),
		startJob("Job 9", 2*time.Second),
		startJob("Job 10", 1*time.Second))

	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println("All tasks were completed!")
}

func fanIn(sources ...<-chan string) <-chan string {
	destination := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(sources))

	for _, source := range sources {
		go func(s <-chan string) {
			defer wg.Done()
			for data := range s {
				destination <- data
			}
		}(source)
	}

	go func() {
		wg.Wait()
		close(destination)
	}()

	return destination
}

func startJob(jobName string, delay time.Duration) <-chan string {
	job := make(chan string)
	go func() {
		defer close(job)
		time.Sleep(delay)
		job <- jobName + " completed!"
	}()
	return job
}
