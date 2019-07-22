package website

import (
	"net/http"
)

// CarParksHandler ...
func CarParksHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "CarParks",
		Page: "carparks",
	})
}

// CarParkHandler ...
func CarParkHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, PageData{
		Title: "CarPark",
		Page: "carpark",
	})
}