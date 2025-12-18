package main

import (
	"fmt"
	redisStorage "lab1812/internal/redis"
	"log"
	"net/http"

	"github.com/chetansj27/crud-go/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	routes.RegisterUserRoutes(route)
	http.Handle("/", route)
	log.Fatal(http.ListenAndServe("localhost:8182", route))
	err := redisStorage.SetKey("todo1", "{task:1}", 0)
	if err == nil {
		log.Printf("ERROR %v \n", err)
	}

	fmt.Println("res")
	log.Fatal(http.ListenAndServe("localhost:8182", route))
}
