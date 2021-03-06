package controllers

import (
	"github.com/FooedBar/fooedserver/controllers/api"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateRouter() http.Handler {
	router := mux.NewRouter()
	apiV0Router := router.PathPrefix("/v0").Subrouter()
	apiV0Router = apiV0Router.StrictSlash(true)
	apiV0Router.HandleFunc("/", Use(api.V0_API, GetContext)).Methods("GET")
	apiV0Router.HandleFunc("/sessions", Use(api.V0_API_Create_Session, GetContext)).Methods("POST")
	apiV0Router.HandleFunc("/menuItems", Use(api.V0_API_Get_Menu_Items, RequireSessionId, GetContext)).Methods("GET")
	apiV0Router.HandleFunc("/restaurants/{restaurantId}", Use(api.V0_API_Get_Restaurant, RequireSessionId, GetContext)).Methods("GET")
	apiV0Router.HandleFunc("/selections", Use(api.V0_API_Post_Selection, RequireSessionId, GetContext)).Methods("POST")
	apiV0Router.HandleFunc("/suggestions/restaurants", Use(api.V0_API_Get_Restaurant_Suggestions, RequireSessionId, GetContext)).Methods("GET")
	apiV0Router.HandleFunc("/suggestions/restaurants/{restaurantId}/menu", Use(api.V0_API_Get_Restaurant_Menu_Suggestions, RequireSessionId, GetContext)).Methods("GET")
	//router.PathPrefix("/uploads/").Handler(uploadsFS)
	return router
}

// `Use` allows us to stack middleware to process the request
// Example taken from https://github.com/gorilla/mux/pull/36#issuecomment-25849172
func Use(handler http.HandlerFunc, mid ...func(http.Handler) http.HandlerFunc) http.HandlerFunc {
	for _, m := range mid {
		handler = m(handler)
	}
	return handler
}
