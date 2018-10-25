package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/loubard/sfapi/payments"
	"github.com/loubard/sfapi/sql"
	log "github.com/sirupsen/logrus"
)

func main() {
	db, err := gorm.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.WithError(err).Fatal("Error opening the database")
	}
	defer db.Close()
	db = db.Set("gorm:auto_preload", true)

	sql.Seed(db)

	r := mux.NewRouter()
	s := r.PathPrefix("/v1/payments").Subrouter()
	s.HandleFunc("/", payments.List(db)).Methods("GET")
	s.HandleFunc("/", payments.Create(db)).Methods("POST")
	s.HandleFunc("/{id}", payments.Fetch(db)).Methods("GET")
	s.HandleFunc("/{id}", payments.Delete(db)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", r))
}
