package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"green": "#4bf745",
	}
	colors["white"] = "#ffffff"

	delete(colors, "red")

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Key: " + k + ", value: " + v)
	}
}