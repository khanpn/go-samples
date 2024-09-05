package main

import "fmt"

func add[T int | float64](a T, b T) T {
	return a + b
}

type myNumber interface {
	~int | ~float64
}

func addT[T myNumber](a T, b T) T {
	return a + b
}

type myAlias int

func main() {
	fmt.Println(add(2, 2.5))
	fmt.Println(addT(2, 2.5))

	var n myAlias = 2
	fmt.Println(addT(n, 2))
}