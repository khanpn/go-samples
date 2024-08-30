package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	contact ContactInfo
}

type ContactInfo struct {
	email string
	zipCode int
}

func (p Person) print() {
	fmt.Printf("%+v",p)
}

func (pointerToPerson *Person) updateFirstName(firstName string) {
	pointerToPerson.firstName = firstName
}

func main() {
	person := Person{
		firstName: "Alex",
		lastName: "Anderson",
		contact: ContactInfo{
			email: "alex@gmail.com",
			zipCode: 17109,
		},
	}
	person.updateFirstName("Khanh")
	updateLastName(&person, "Nguyen")
	person.print()
}

func updateLastName(p *Person, lastName string) {
	p.lastName = lastName
}