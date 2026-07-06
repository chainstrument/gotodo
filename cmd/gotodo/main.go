package main

import (
	"log"
	"net/http"

	"github.com/chainstrument/gotodo/internal/todo"
)

func main() {
	store := todo.NewStore()
	handler := todo.NewHandler(store)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /todos", handler.List)
	mux.HandleFunc("POST /todos", handler.Create)
	mux.HandleFunc("GET /todos/{id}", handler.Get)
	mux.HandleFunc("PUT /todos/{id}", handler.Update)
	mux.HandleFunc("DELETE /todos/{id}", handler.Delete)

	const addr = ":8080"
	log.Printf("gotodo listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
