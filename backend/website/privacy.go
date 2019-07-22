package website

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// PrivacyHandler ...
func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Sprintf("Working Dir Err: %v", err))
	}

	html := template.Must(template.ParseFiles(wd + "/frontend/privacy.html"))
	err = html.Execute(w, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("Template Err: %v", err))
	}
}