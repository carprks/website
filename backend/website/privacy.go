package website

import (
	"net/http"
)

// PrivacyHandler ...
func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title: "Privacy",
		Page: "privacy",
		LoggedIn: false,
	})
}