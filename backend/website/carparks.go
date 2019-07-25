package website

import (
	"net/http"
)

// CarParksHandler ...
func CarParksHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title: "CarParks",
		Page: "carparks",
	})
}

// CarParkHandler ...
func CarParkHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, PageData{
		Title: "CarPark",
		Page: "carpark",
	})
}