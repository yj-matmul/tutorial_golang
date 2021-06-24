package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

//
var functions = template.FuncMap{}

// renderTemplate renders tempalte using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	_, err := RenderTemplateTest(w)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	pasredTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err = pasredTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

//
func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
