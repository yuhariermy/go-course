package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// 1st trying to create a struct
type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "John",
			"last_name": "Doe",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Black",
			"last_name": "Mamba",
			"hair_color": "Blue",
			"has_dog": false
		}
	]`

	// read the JSON --> unmarshal
	// 2nd unmarshall the json of myJson
	var unmarshal []Person
	err := json.Unmarshal([]byte(myJson), &unmarshal)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	} 

	log.Printf("unmarshalled json: %v", unmarshal)

	// write into the json
	var mySliceAfterMarshal []Person
	
	var m1 Person
	m1.FirstName = "Testing"
	m1.LastName = "test"
	m1.HairColor = "black"
	m1.HasDog = true

	mySliceAfterMarshal = append(mySliceAfterMarshal, m1)
	
	var m2 Person
	m2.FirstName = "Blabla"
	m2.LastName = "blublu"
	m2.HairColor = "black"
	m2.HasDog = false

	mySliceAfterMarshal = append(mySliceAfterMarshal, m2)

	newJson, err := json.MarshalIndent(mySliceAfterMarshal, "", "        ")
	if err != nil {
		fmt.Println("Marshal error: ", err)
	}

	fmt.Println(string(newJson)) // the output was byte, therefore string() :> is stand for change the byte into string
}