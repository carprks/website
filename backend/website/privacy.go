package website

import (
  "net/http"
)

// Privacy ...
type Privacy struct {
  Title []TitlePart
  Content []PrivacyParts
}
// TitlePart ...
type TitlePart struct {
  Type string `json:"type"`
  Text string `json:"text"`
}
// PrivacyParts ...
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

// PrivacyCookieHandler ...
func PrivacyCookieHandler(w http.ResponseWriter, r *http.Request) {
  RenderTemplate(w, r, PageData{
    Title: "Cookie Policy",
    Page: "cookie",
  })
}

// PrivacyDataHandler ...
func PrivacyDataHandler(w http.ResponseWriter, r *http.Request) {
  RenderTemplate(w, r, PageData{
    Title: "Data Privacy Policy",
    Page: "data-privacy",
  })
}
