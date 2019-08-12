package website

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

// CustomClaims ...
type CustomClaims struct {
	Identifier string `json:"identifier"`
	jwt.StandardClaims
}

func saveJWT(w http.ResponseWriter, r *http.Request, lr LoginResponse) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		Identifier: lr.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNING_SECRET")))
	if err != nil {
		fmt.Println(fmt.Errorf("signing err: %v", err))
	}

	http.SetCookie(w, &http.Cookie{
		Name: "ninjaToken",
		Value: tokenString,
		MaxAge: 600,
		Path: "/",
	})
}

func deleteJWT(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name: "ninjaToken",
		MaxAge: -500,
		Path: "/",
		Value: "",
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

	token, err := jwt.Parse(cookie.Value, func (token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNING_SECRET")), nil
	})

	if token != nil {
	  if token.Valid {
		  return true
	  }
	}

	return false
}
