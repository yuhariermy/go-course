package main

import "fmt"

func main() {

	// normal variable
	var someVariable = "Variable"
	fmt.Println("normal Variable must be:", someVariable)

	// variable is changed by pointers function using &someVariableName
	pointersFunction(&someVariable)
	fmt.Println("new Variable must be:", someVariable)
}

// create function that includes a pointers, its represent by *someDataType
func pointersFunction(s *string) { // *string is a pointer
	newValue := "New Variable"
	*s = newValue
}
