package payments

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/loubard/sfapi/models"
	"github.com/loubard/sfapi/sql"
)

// Fetch returns a payment resource based on the id
func Fetch(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		p, err := sql.GetByID(db, vars["id"])
		if err != nil || vars["id"] == "" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		d := models.FetchResponse{Data: p}
		j, err := json.Marshal(d)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
		}
	}
}

// List returns all payment resources
func List(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p := sql.GetAll(db)

		d := models.ListResponse{Data: p}
		j, err := json.Marshal(d)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
		}
	}
}

// Delete a payment resource
func Delete(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		err := sql.Delete(db, vars["id"])

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
		}
	}
}
