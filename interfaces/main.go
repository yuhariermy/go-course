package main

import (
	"fmt"
)

// interface is stands for the rule of something it must have same datatype functions inside of it
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// struct is stands for format of Object
type Dog struct {
	FirstName string
	LastName string
}

type Gorilla struct { 
	Name string
	Age int
	NumberOfTeeth int
}

func main() {
	dog := Dog {
		FirstName: "Lilo",
		LastName: "Docker",
	}

	gorilla := Gorilla{
		Name: "Docker",
		Age: 10,
		NumberOfTeeth: 20,
	}

	PrintInfo(&dog) // must be called using new function to integrate between interface and struct using same info with interface (Says(), and NumberOfLegs())
	PrintInfo(&gorilla) // must be called using new function to integrate between interface and struct
}

func PrintInfo(a Animal) {
	fmt.Println("This Animalsays: ", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

// function to integrate between interface and struct

func (d *Dog) Says() string { // the reciever it has to be has its own pointers * with his pair
	return "Woof"
}

func (d *Dog) NumberOfLegs() int { // the reciever it has to be has its own pointers * with his pair
	return 4
}

func (g *Gorilla) Says() string { // the reciever it has to be has its own pointers 
	return "Auwowo"
}

func (g *Gorilla) NumberOfLegs() int { // the reciever it has to be has its own 
	return 2
}