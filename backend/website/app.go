package website

import (
	"net/http"
)

// AppHandler ...
func AppHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "App",
		Page: "app",
		LoggedIn: false,
	})
}