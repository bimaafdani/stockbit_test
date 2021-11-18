package main

import (
	"log"
	"net/http"
	"stocbit_exam/HTTP/rest"

	"github.com/gorilla/mux"
)

func main() {
	routes := rest.InitPackage()

	//init mux routes
	router := mux.NewRouter()
	routes.InitRoutes(router)

	log.Printf("API Gateway up and running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
