package main

import "fmt"

type structWithFunction struct {
	FirstName string
}

// Reciever is a pointer that can be used to call a key of struct into function
func (reciever *structWithFunction) exampleOfReciever() string {
	return reciever.FirstName
}

func main() {
	// a way to declare variable
	var myVar structWithFunction // way 1
	myVar.FirstName = "Variable1"

	myVar2 := structWithFunction {
		FirstName: "Variable2",
	} // way 2

	// a manually way to call a value of variable
	// fmt.Println("Calling a value of Variable 1:", myVar.FirstName)
	// fmt.Println("Calling a value of Variable 2:", myVar2.FirstName)
	
	// a way to call variable using structs reciever function
	fmt.Println("Calling a value of Variable 1:", myVar.exampleOfReciever())
	fmt.Println("Calling a value of Variable 2:", myVar2.exampleOfReciever())
	
}