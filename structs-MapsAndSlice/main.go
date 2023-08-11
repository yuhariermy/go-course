package main

import (
	"log"
	"sort"
)

func main() {
	// maps :> is a way to store with key value
	var myFirstMap = make(map[string]string) // this string can be changed by any data type like int, struct, or interface{}
	myFirstMap["myFirst"] = "my first String Value" // this string value if changed by any data type should follow base on above data type
	log.Println("myFirstMap is:", myFirstMap["myFirst"])


	// slice :> is a way to store any data type into [] like an array
	var mySlice []string // 1st way to call slice
	mySlice = append(mySlice, "firstElement")
	mySlice = append(mySlice, "secondElement")

	mySlice2 := []int{1, 2, 3,4,5,7,6,8,9} // 2nd way to call slice
	sort.Ints(mySlice2) // a way to sort the slice for an integer


	log.Println("mySlice is:", mySlice)
	log.Println("mySlice is:", mySlice2[0:2]) // start with first element :> 0 index and end before index 2 which mean index 1
}