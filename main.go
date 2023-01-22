package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8000"

func Home(w http.ResponseWriter, r *http.Request){
	renderTemplate(w, "home.html")
	
}

func About(w http.ResponseWriter, r *http.Request){
	
	renderTemplate(w, "about.html")
	
}

func renderTemplate(w http.ResponseWriter, tmpl string){
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error pasing template", err)
		return
	}
}

func main() {


	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	

	fmt.Println("starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}