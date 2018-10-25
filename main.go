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
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		log.WithError(err).Fatal("Error opening the database")
	}
	defer db.Close()

	sql.Seed(db)

	r := mux.NewRouter()
	s := r.PathPrefix("/v1/payments").Subrouter()
	s.HandleFunc("/{id}", payments.Fetch)
	log.Fatal(http.ListenAndServe(":8080", r))
}
