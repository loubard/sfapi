package payments

import (
	"encoding/json"
	"io/ioutil"
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
	}
}

// Create a payment resource
func Create(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var p models.Payment
		err = json.Unmarshal(requestBody, &p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = sql.Create(db, &p)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// Update a payment resource
func Update(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var p models.Payment
		err = json.Unmarshal(requestBody, &p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		vars := mux.Vars(r)
		err = sql.Update(db, vars["id"], &p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
