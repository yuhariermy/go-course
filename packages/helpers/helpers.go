package helpers

import (
	"math/rand"
	"time"
)

type PackageStruct struct {
	FirstName string
	Age       int
}

// It can defined function, interface, and any data structure
func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}