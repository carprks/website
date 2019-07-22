package website

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// PageData structure
type PageData struct {
	Title string
	Page string
	LoggedIn bool
	Links []string
}

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, data PageData) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Sprintf("Working Dir Err: %v", err))
	}

	data.Links = []string{
		"carparks",
		"pricing",
		"companies",
		"app",
		"about",
	}

	// layout
	t := template.Must(template.ParseFiles(fmt.Sprintf("%s/frontend/layout.html", wd), fmt.Sprintf("%s/frontend/pages/%s.html", wd, data.Page)))
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(fmt.Sprintf("Template Err: %v", err))
	}
}