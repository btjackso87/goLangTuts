package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2,2)
	fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum ))
}

func addValues(x,y int) int{
	var sum int
	sum = x + y
	return sum 
}

func main() {


	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}