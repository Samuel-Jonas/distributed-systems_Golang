package routes

import (
	"dataViewer/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/createUser", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/getUserId/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/getUsers", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/updateUserId/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUserId/{id}", controllers.DeleteUser).Methods("DELETE")
}
