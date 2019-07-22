package website

import (
	"net/http"
)

// PrivacyHandler ...
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "About",
		Page: "about",
		LoggedIn: false,
	})
}