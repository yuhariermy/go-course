package main

import (
	"errors"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Homepage")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d", sum))
}
func Divide( w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return // to stop the function as soon as possible
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0, f))
}

func divideValues(a, b float64) (float64, error) {
	if b == 0 {
		err := errors.New("Cannot divide by 0")
		return 0, err
	}
	result := a / b
	return result, nil
}

func addValues(x, y int) int {
	return x + y
}

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on server %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
