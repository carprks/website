package website

import (
	"net/http"
)

// PrivacyHandler ...
func AppHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "App",
		Page: "app",
		LoggedIn: false,
	})
}