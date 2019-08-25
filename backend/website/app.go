package website

import (
	"net/http"
)

// AppHandler ...
func AppHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title:    "App",
		Page:     "app",
		LoggedIn: false,
	})
}
