package website

import (
	"net/http"
)

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "Login",
		Page: "login",
		LoggedIn: false,
	})
}

// RegisterHandler ...
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "Register",
		Page: "register",
		LoggedIn: false,
	})
}