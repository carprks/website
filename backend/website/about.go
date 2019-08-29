package website

import (
	"net/http"
)

// AboutHandler ...
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title:    "About",
		Page:     "about",
		PagePath: "about",
	})
}
