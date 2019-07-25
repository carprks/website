package website

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var (
	Version string
	Build string
)

// Link structure
type Link struct {
	Title string
	Link string
}

// Links structure
type Links struct {
	Navigation []Link
	Footer []Link
}

// PageData structure
type PageData struct {
	Title string
	Page string
	LoggedIn bool
	Links Links
	Version string
	Build string
}

// RenderTemplate ...
func RenderTemplate(w http.ResponseWriter, r *http.Request, data PageData) {
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
				Link: "carparks",
			},
			{
				Title: "Pricing",
				Link: "pricing",
			},
			{
				Title: "Companies",
				Link: "companies",
			},
			{
				Title: "App",
				Link: "app",
			},
			{
				Title: "About",
				Link: "about",
			},
		},
		Footer: []Link{
			{
				Title: "Contact Us",
				Link: "contact",
			},
			{
				Title: "About Us",
				Link: "about",
			},
			{
				Title: "Privacy Policy",
				Link: "privacy",
			},
		},
	}

	if checkJWT(r) {
		data.LoggedIn = true
	}

	// layout
	t := template.Must(template.ParseFiles(fmt.Sprintf("%s/frontend/layout.html", wd), fmt.Sprintf("%s/frontend/pages/%s.html", wd, data.Page)))
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(fmt.Sprintf("Template Err: %v", err))
	}
}