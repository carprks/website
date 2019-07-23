package website

import (
	"net/http"
)

// ContactHandler ...
func ContactHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "Contact",
		Page: "contact",
		LoggedIn: false,
	})
}