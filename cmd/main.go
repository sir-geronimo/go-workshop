package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sir-geronimo/go-workshop/pkg/routes"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRouters(r)
	http.Handle("/", r)

	fmt.Println("Starting server at port: 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
