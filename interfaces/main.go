package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct{}
type VietnameseBot struct{}

func main() {
	englishBot := EnglishBot{}
	vietnameseBot := VietnameseBot{}

	printGreeting(englishBot)
	printGreeting(vietnameseBot)
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func (EnglishBot) getGreeting() string {
	return "Hi there!"
}

func (VietnameseBot) getGreeting() string {
	return "Xin chào!"
}