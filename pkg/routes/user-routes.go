package routes

import (
	"github.com/gorilla/mux"
	_ "github.com/rushabh2390/gousersmodule/docs"
	"github.com/rushabh2390/gousersmodule/pkg/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterUserstore(router *mux.Router) {

	router.HandleFunc("/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/login", controllers.LoginUser).Methods("POST")

	// Serve Swagger UI
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
