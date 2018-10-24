package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/loubard/sfapi/payments"
)

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/v1/payments").Subrouter()
	s.HandleFunc("/{id}", payments.Fetch)
	log.Fatal(http.ListenAndServe(":8080", r))
}
