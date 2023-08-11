package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// variable as string
	var someString string
	someString = "Some string"
	// variable as int

	var someInteger int
	someInteger = 1

	fmt.Println("Some string: ", someString)
	fmt.Println("Some integer: ", someInteger)

	// declare a function with variable, with :=
	someFunction, someFunctionInteger := someFunction()
	fmt.Println("some function: ", someFunction, "some function integer: ", someFunctionInteger)
}

func someFunction() (string, int) {
	return "something", 1
}
