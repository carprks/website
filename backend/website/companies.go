package website

import (
	"net/http"
)

// CompaniesHandler ...
func CompaniesHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title:    "Companies",
		Page:     "companies",
		LoggedIn: false,
	})
}
