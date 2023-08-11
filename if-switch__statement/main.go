package main

import "log"

func main() {
	// If statement is stand for comparison, > greater then >= is equal greater then and etc...
	var myVar = 110
	isTrue := true

	if myVar >= 100 && !isTrue {
		log.Println("Your conditional work properly")
	} else if myVar == 0 || isTrue {
		log.Println("Your number is 0 or true")
	} else {
		log.Println("Try Again...")
	}

	// Switch Statement is stand something is equal
	myPet := "horse"

	switch myPet {
	case "cat":
		log.Println("Your pet is cat")
	case "dog":
		log.Println("Your pet is dog")
	case "fish":
		log.Println("Your pet is fish")
	default:
		log.Println("Try Again...")
	}
}
