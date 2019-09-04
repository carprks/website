package website

import (
	"bytes"
	"encoding/json"
	"fmt"
	permissions "github.com/carprks/permissions/service"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"os"
)

// CustomClaims ...
type CustomClaims struct {
	Identifier string `json:"identifier"`
	jwt.StandardClaims
}

func saveJWT(w http.ResponseWriter, r *http.Request, lr loginResponse) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Identifier: lr.Identifier,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_SECRET")))
	if err != nil {
		fmt.Println(fmt.Errorf("signing err: %v", err))
	}

	cookie := http.Cookie{
		Name:   "ninjaToken",
		Value:  tokenString,
		MaxAge: 6000,
		Path:   "/",
	}
	r.AddCookie(&cookie)
	http.SetCookie(w, &cookie)
}

func deleteJWT(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   "ninjaToken",
		MaxAge: -500,
		Path:   "/",
		Value:  "",
	})
}

func checkJWT(r *http.Request) bool {
	cookie, err := r.Cookie("ninjaToken")
	if err != nil {
		// fmt.Println(fmt.Sprintf("cookie err: %v", err))
		return false
	}

	if cookie.Value == "" {
		return false
	}

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_SECRET")), nil
	})

	if token != nil {
		if token.Valid {
			return true
		}
	}

	return false
}

func getIdentifier(r *http.Request) string {
	identifier := ""

	cookie, err := r.Cookie("ninjaToken")
	if err != nil {
		return identifier
	}

	if checkJWT(r) {
		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNING_SECRET")), nil
		})

		if err != nil {
			return identifier
		}

		if token != nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				identifier = fmt.Sprintf("%v", claims["identifier"])
			}
		}
	}

	return identifier
}

func checkAllowed(p permission, r *http.Request) bool {
	ident := getIdentifier(r)

	if p.Identifier == "" {
		p.Identifier = ident
	}

	l := permissions.Permissions{
		Identifier: ident,
		Permissions: []permissions.Permission{
			{
				Name: p.Name,
				Action: p.Action,
				Identifier: p.Identifier,
			},
		},
	}

	j, err := json.Marshal(&l)
	if err != nil {
		fmt.Println(fmt.Sprintf("JSON login err: %v", err))
		return false
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/allowed", os.Getenv("SERVICE_ACCOUNT")), bytes.NewBuffer(j))
	if err != nil {
		fmt.Println(fmt.Sprintf("req err: %v", err))
	}
	req.Header.Set("X-Authorization", os.Getenv("AUTH_KEY_ACCOUNT"))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(fmt.Sprintf("client err: %v", err))
		return false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Sprintf("account resp err: %v", err))
		return false
	}
	if resp.StatusCode == 200 {
		lr := permissions.Permissions{}
		jerr := json.Unmarshal(body, &lr)
		if jerr != nil {
			fmt.Println(fmt.Sprintf("account decode err: %v", jerr))
			return false
		}

		if lr.Status == "allowed" {
			return true
		}
	}

	return false
}