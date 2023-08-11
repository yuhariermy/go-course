package main

import (
	"log"
	"testPackage_name/helpers" // is created using go.mod into helpers folder
)

const numPool = 10

func CalculateValue(intChan chan int) { // 2.A executed when second way is running
	randomNumber := helpers.RandomNumber(numPool) // 3. while 2.A is running, helpers will check the RandomNumber function
	intChan <- randomNumber // random number pass to intChan
}

func main() {

	// stand for Packages
	log.Println("Hello")
	var testingPackage = helpers.PackageStruct { // it use like package log
		FirstName: "PackageStruct FirstName",
		Age: 10,
	}

	log.Println("testingPackage FirstName:", testingPackage.FirstName, "testingPackage Age:", testingPackage.Age)

	// stands for Channel, channel is third way to pass information
	intChan := make(chan int) // 1. first way that start in the beginning of channel program
	defer close(intChan) // 6. defer stand for whenever is program stop execute when is done

	go CalculateValue(intChan) // 2. second way to run program
	num := <- intChan // 4. the intChan information will pass to num
	log.Println(num) // 5. num is executed
}