package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Employee struct {
	FullName string			`json:"full_name"`
	BirthDate time.Time		`json:"birthdate"`
	Email string			`json:"email"`
	WorkPhone string		`json:"work_phone"`
	HomePhone string		`json:"home_phone,omitempty"`
}

func main() {
	birthdate, _ := time.Parse("yyyy-mm-dd", "1890-04-15")
	emp := Employee {
		FullName: "Oliver Montgomery",
		BirthDate: birthdate,
		Email: "Oliver@gmail.com",
		WorkPhone: "(641) 451-0000",
	}

	json_data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}
	json_string := string(json_data)
	fmt.Println("JSON string", json_string)

	var empFromJson Employee
	json.Unmarshal(json_data, &empFromJson)

	fmt.Println("Employee from JSON", empFromJson)
}