package website

import (
	"net/http"
)

// CompaniesHandler ...
func CompaniesHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "Companies",
		Page: "companies",
		LoggedIn: false,
	})
}