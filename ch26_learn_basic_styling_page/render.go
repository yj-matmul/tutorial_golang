package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate renders the tempalte language
func renderTemplate(w http.ResponseWriter, tmpl string) {
	pasredTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := pasredTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
