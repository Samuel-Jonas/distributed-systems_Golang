package routes

import (
	"dataViewer/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
}
