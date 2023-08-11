package main

import "log"

func main() {
	// basic for loop
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	// Loop using range
	rangeLoopString := []string{"dog","cat","horse", "bird"}

	// for i, loopString := range rangeLoopString { // i stand for index
	for _, loopString := range rangeLoopString { // _ stand for ignoring index
		log.Println(loopString)
	}

	// loop using range on map
	myMapVariable := make(map[string]int)
	myMapVariable["Number1"] = 1
	myMapVariable["Number2"] = 2
	myMapVariable["Number3"] = 3

	// for _, number := range myMapVariable { // just print value
	for nameOfKey, number := range myMapVariable {
		// log.Println(number) // to print value
		log.Println(nameOfKey, number) // to print pair key-value
	}

	// using struct into loop
	type User struct {
		FirstName string
		LastName  string
		Age       int
		Email     string
	}

	var newUser []User
	// create or update the struct using append format
	newUser = append(newUser, User{"John", "Smith", 10, "john@smith.com"})
	newUser = append(newUser, User{"Conro", "Michale", 15, "Conro@Michale.com"})
	newUser = append(newUser, User{"John", "Blabla", 30, "john@Blabla.com"})
	newUser = append(newUser, User{"John", "blublu", 50, "john@blublu.com"})

	for _, user := range newUser {
		log.Println(user.FirstName, user.LastName, user.Email, user.Age)
	}

}
