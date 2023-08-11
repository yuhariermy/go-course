package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divider(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("The result is: ", result)
}

// writing the test using divider function

func divider(x, y float32) (float32, error) {
	var result float32
	
	if y == 0 {
		return result, errors.New("Cannot divide by zero")
	}
	result = x / y
	// return result, error // should define the error
	return result, nil
}