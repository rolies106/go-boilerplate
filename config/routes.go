package config

import (
	"github.com/gorilla/mux"

	"net/http"

	"mortred/app"
	"mortred/controllers"
)

/**
 * setup routes
 * @return mux.Router
 */
func Setup() *mux.Router {

	router := mux.NewRouter()

	// Version 1 of API that not using JWT auth
	v1NoAuth := router.PathPrefix("/v1").Subrouter()
	v1NoAuth.HandleFunc("/users/buyer", controllers.BuyerCreate).Methods("POST")
	v1NoAuth.HandleFunc("/users/buyer/auth", controllers.BuyerAuthenticate).Methods("POST")

	// Version 1 of API that need auth
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.Use(app.JwtAuthentication) //attach JWT auth middleware

	// Adding default not found handler
	router.NotFoundHandler = http.HandlerFunc(app.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.NotFoundHandler)

	return router
}
