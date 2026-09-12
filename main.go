package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", helloWorld)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("server serving at http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
