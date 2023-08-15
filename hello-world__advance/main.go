package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the Homepage")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2,2)

	_,_ = fmt.Fprintf(w, fmt.Sprintf("This is the About page and 2 + 2 is %d",sum))
}

func addValues(x,y int) int {
	return x + y
}

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on server %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
