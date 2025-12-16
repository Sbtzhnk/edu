package main

import (
	"lab1312lab1/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterUserRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8182", route))
}
