package main

import (
	"dataViewer/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)

	routes.RegisterUserRoutes(router)

	log.Println(fmt.Sprintf("Starting server on port %s", "5000"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 5000), router))
}
