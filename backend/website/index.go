package website

import (
	"net/http"
)

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title:    "Home",
		Page:     "home",
		LoggedIn: false,
	})
}
