package main

import "fmt"

func main() {
	c := make(chan int)

	go sendOnly(c)

	receiveOnly(c)

	fmt.Println("about to exit")
}

func sendOnly(c chan<- int) {
	c <- 10
}

func receiveOnly(c <-chan int) {
	fmt.Println(<-c)
}
