package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandleFunc
}

var Routes []Route

var routes = Routes{
	Route{
		"GetAllReservations",
		"GET",
		"/api/v1/resto/{venue_id}/getAllReservations",
		api.GetAllReservations,
	},
	Route{
		"GetReservations",
		"GET",
		"/api/v1/resto/{venue_id}/getReservations/{id}",
		api.GetReservations,
	},
	Route{
		"AddReservations",
		"POST",
		"/api/v1/resto/{venue_id}/addReservations",
		api.AddReservations,
	},
	Route{
		"UpdateReservations",
		"PUT",
		"/api/v1/resto/{venue_id}/updateReservations/{id}",
		api.UpdateReservations,
	},
	Route{
		"CancelReservations",
		"DELETE",
		"/api/v1/resto/{venue_id}/cancelReservations/{id}",
		api.CancelReservations,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
