package main

import (
	"fmt"
	"net/http"
)

func main() {
	api := &Api{addr: ":8080"}

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /", api.home)
	mux.HandleFunc("GET /users", api.getUsers)
	mux.HandleFunc("POST /users", api.createUsers)

	fmt.Println("server starting at http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
