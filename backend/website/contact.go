package website

import (
	"net/http"
)

// PrivacyHandler ...
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "Contact",
		Page: "contact",
		LoggedIn: false,
	})
}