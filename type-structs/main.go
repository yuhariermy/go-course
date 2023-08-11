package main

import (
	"fmt"
	"time"
)

// type is use for Variable
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthday    time.Time
}

func main() {
	user := User {
		FirstName: "John",
		LastName: "Chena",
	}

	fmt.Println("First Name: ", user.FirstName)
	fmt.Println("Last Name: ", user.LastName)

}