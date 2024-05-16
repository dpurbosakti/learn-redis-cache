package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

/*
	localhost:8080/api/books?subject=love

hit the endpoint 2 times to see the time different
*/

func main() {
	h := NewBookHandler()

	router := chi.NewRouter()
	router.Get("/api/books", h.getBySubject)

	log.Printf("start http server at %s", "0.0.0.0:8080")
	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}
