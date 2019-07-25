package website

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type loginObject struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	pd := PageData{
		Title: "Login",
		Page: "login",
	}

	if r.Method == "POST" {
		l := loginObject{}

		err := r.ParseForm()
		if err != nil {
			fmt.Println(fmt.Sprintf("Parse Form err: %v", err))
		}
		for key, value := range r.Form {
			switch key {
			case "loginEmail":
				l.Email = value[0]
			case "loginPassword":
				l.Password = value[0]
			}
		}


		j, err := json.Marshal(&l)
		if err != nil {
			fmt.Println(fmt.Sprintf("JSON login err: %v", err))
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", os.Getenv("SERVICE_LOGIN")), bytes.NewBuffer(j))
		if err != nil {
			fmt.Println(fmt.Sprintf("req err: %v", err))
		}
		req.Header.Set("X-Authorization", os.Getenv("AUTH_KEY"))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(fmt.Sprintf("client err: %v", err))
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			saveJWT(w, r)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	RenderTemplate(w, r, pd)
}

// RegisterHandler ...
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title: "Register",
		Page: "register",
		LoggedIn: false,
	})
}

// LogoutHandler ...
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	deleteJWT(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// AccountHandler ...
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title: "Account",
		Page: "account",
	})
}