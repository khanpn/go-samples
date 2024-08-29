package main

import "fmt"

func main() {
	n := 7

	if (n & 1 == 1) {
		fmt.Printf("%d is odd\n", n)
	} else {
		fmt.Printf("%d is even\n", n)
	}

}