package website

import (
	"net/http"
)

// PricingHandler ...
func PricingHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title:    "Pricing",
		Page:     "pricing",
		LoggedIn: false,
	})
}
