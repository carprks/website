package website

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	// Version of the build
	Version string

	// Build number
	Build string
)

// Link structure
type Link struct {
	Title string
	Link  string
}

// Links structure
type Links struct {
	Navigation []Link
	Footer     []Link
}

// PageData structure
type PageData struct {
	Title    string
	Page     string
	PagePath string
	LoggedIn bool
	Links    Links
	Version  string
	Build    string
	Content  interface{}
	Permission permission
}

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, r *http.Request, data PageData) {
	distPath := "frontend"
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Sprintf("Working Dir Err: %v", err))
	}

	// Version and Build
	data.Version = Version
	data.Build = Build

	// Navigation
	data.Links = Links{
		Navigation: []Link{
			{
				Title: "CarParks",
				Link:  "carparks",
			},
			{
				Title: "Pricing",
				Link:  "pricing",
			},
			{
				Title: "Companies",
				Link:  "companies",
			},
			{
				Title: "App",
				Link:  "app",
			},
			{
				Title: "About",
				Link:  "about",
			},
		},
		Footer: []Link{
			{
				Title: "Contact Us",
				Link:  "contact",
			},
			{
				Title: "About Us",
				Link:  "about",
			},
			{
				Title: "Privacy Policy",
				Link:  "privacy",
			},
		},
	}

	if checkJWT(r) {
		data.LoggedIn = true
	}

	if data.Permission.Name != "" && data.Permission.Action != "" {
		allowed := checkAllowed(data.Permission, r)
		if !allowed {
			data.Page = "home"
			data.PagePath = "/"
		}
	}

	// layout
	layoutPath := fmt.Sprintf("%s/%s/layout.html", wd, distPath)
	pagePath := fmt.Sprintf("%s/%s/pages/%s.html", wd, distPath, data.Page)
	if data.PagePath != "" {
		pagePath = fmt.Sprintf("%s/%s/pages/%s/%s.html", wd, distPath, data.PagePath, data.Page)
	}

	t := template.Must(template.ParseFiles(layoutPath, pagePath))
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(fmt.Sprintf("Template Err: %v", err))
	}
}
