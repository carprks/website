package website

import (
  "net/http"
)

type Privacy struct {
  Title []TitlePart
  Content []PrivacyParts
}
type TitlePart struct {
  Type string `json:"type"`
  Text string `json:"text"`
}
type PrivacyParts struct {
  Type string `json:"type"`
  Text string `json:"text"`
}

// PrivacyHandler ...
func PrivacyHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title: "Privacy Policy",
		Page: "privacy",
		LoggedIn: false,
	})
}

func PrivacyCookieHandler(w http.ResponseWriter, r *http.Request) {
  RenderTemplate(w, r, PageData{
    Title: "Cookie Policy",
    Page: "cookie",
  })
}

func PrivacyDataHandler(w http.ResponseWriter, r *http.Request) {
  RenderTemplate(w, r, PageData{
    Title: "Data Privacy Policy",
    Page: "data-privacy",
  })
}
