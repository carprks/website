package website

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type loginObject struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Error       string       `json:"error"`
	Identifier  string       `json:"identifier"`
	Permissions []permission `json:"permissions"`
}

type permission struct {
	Name       string `json:"name"`
	Action     string `json:"action"`
	Identifier string `json:"identifier"`
}

type registerObject struct {
	Email           string `json:"email"`
	ConfirmEmail    string
	Password        string `json:"password"`
	ConfirmPassword string `json:"verify"`
}

type registerResponse struct {
	Error string `json:"error"`
	ID    string `json:"id"`
}

// LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	pd := PageData{
		Title:    "Login",
		Page:     "login",
		PagePath: "account",
	}

	if r.Method == "POST" {
		l := loginObject{}

		err := r.ParseForm()
		if err != nil {
			fmt.Println(fmt.Sprintf("Parse Form err: %v", err))
			pd.Content = "Invalid Login"

			RenderTemplate(w, r, pd)
			return
		}
		for key, value := range r.Form {
			switch key {
			case "login-email":
				l.Email = value[0]
			case "login-password":
				l.Password = value[0]
			}
		}

		if l.Password == "" && l.Email == "" {
			pd.Content = "Email and Password not recognised"

			RenderTemplate(w, r, pd)
			return
		}

		j, err := json.Marshal(&l)
		if err != nil {
			fmt.Println(fmt.Sprintf("JSON login err: %v", err))
		}

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", os.Getenv("SERVICE_ACCOUNT")), bytes.NewBuffer(j))
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
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(fmt.Sprintf("account resp err: %v", err))
				return
			}
			lr := loginResponse{}
			jerr := json.Unmarshal(body, &lr)
			if jerr != nil {
				fmt.Println(fmt.Sprintf("account decode err: %v", jerr))
				return
			}

			if canLogin(lr) {
				lr.Error = "account banned"
			}

			if lr.Error != "" {
				fmt.Println(fmt.Sprintf("lr err: %v", lr))
				// http.Redirect(w, r, "/account/login", http.StatusSeeOther)
				pd.Content = lr.Error

				RenderTemplate(w, r, pd)
				return
			}
			saveJWT(w, r, lr)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	RenderTemplate(w, r, pd)
}

// RegisterHandler ...
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	pd := PageData{
		Title:    "Register",
		Page:     "register",
		PagePath: "account",
	}
	content := []string{}

	if r.Method == "POST" {
		ro := registerObject{}
		err := r.ParseForm()
		if err != nil {
			fmt.Println(fmt.Sprintf("Parse Form Err: %v", err))
			content = append(content, "Invalid Register")
			pd.Content = content
			RenderTemplate(w, r, pd)
			return
		}

		for key, value := range r.Form {
			switch key {
			case "register-email":
				ro.Email = value[0]
			case "register-confirm-email":
				ro.ConfirmEmail = value[0]
			case "register-password":
				ro.Password = value[0]
			case "register-confirm-password":
				ro.ConfirmPassword = value[0]
			}
		}

		if strings.Compare(ro.Email, ro.ConfirmEmail) != 0 {
			content = append(content, "Emails don't match")
		}
		if strings.Compare(ro.Password, ro.ConfirmPassword) != 0 {
			content = append(content, "Passwords don't match")
		}

		if ro.Email == "" || ro.Password == "" {
			content = append(content, "You haven't filled all the form in")
			pd.Content = content
			RenderTemplate(w, r, pd)
			return
		}

		j, err := json.Marshal(&ro)
		if err != nil {
			content = append(content, "Invalid Register")
			pd.Content = content
			RenderTemplate(w, r, pd)
			return
		}
		fmt.Println(string(j))

		req, err := http.NewRequest("POST", fmt.Sprintf("%s/register", os.Getenv("SERVICE_ACCOUNT")), bytes.NewBuffer(j))
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
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(fmt.Sprintf("account resp err: %v", err))
				return
			}
			fmt.Println(string(body))
		}
	}

	pd.Content = content
	RenderTemplate(w, r, pd)
}

// LogoutHandler ...
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	deleteJWT(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// AccountHandler ...
func AccountHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title:    "Account",
		Page:     "account",
		PagePath: "account",
	})
}

// ForgotHandler ...
func ForgotHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		fmt.Println("POST")
	}

	fmt.Println(r.Method)
}

func canLogin(response loginResponse) bool {
	for _, perm := range response.Permissions {
		if perm.Action == "account" && perm.Name == "login" {
			return true
		}
	}

	return false
}
