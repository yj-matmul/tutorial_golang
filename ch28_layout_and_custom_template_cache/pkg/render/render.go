package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// implement custom function when go lang have no built-in function I want to
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
		fmt.Println("Error parsing template:", err)
		return
	}
}

// RenderTemplateTest matches page to layout and creates template cache
func RenderTemplateTest(w http.ResponseWriter) (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		fmt.Println("selected page path is", page)
		name := filepath.Base(page)
		fmt.Println("selected page is", name)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return templateCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseFiles("./templates/*.layout.html")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = ts
	}

	return templateCache, nil
}
