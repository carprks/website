package website

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
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

	now := time.Now()
	cookie := http.Cookie{
		Name:   "ninjaToken",
		Value:  tokenString,
		MaxAge: 6000,
		Path:   "/",
		// Domain: os.Getenv("DOMAIN_NAME"),
		Expires: now.AddDate(0, 1, 0),
	}
	r.AddCookie(&cookie)
	http.SetCookie(w, &cookie)

	fmt.Println(fmt.Sprintf("cookie: %v", cookie))
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
