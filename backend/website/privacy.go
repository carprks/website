package website

import (
	"net/http"
)

// PrivacyHandler ...
func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "Privacy",
		Page: "privacy",
		LoggedIn: false,
	})
}