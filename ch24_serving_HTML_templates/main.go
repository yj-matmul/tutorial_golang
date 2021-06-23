package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

// renderTemplate is
func renderTemplate(w http.ResponseWriter, tmpl string) {
	pasredTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := pasredTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on prot %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
