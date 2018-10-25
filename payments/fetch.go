package payments

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/loubard/sfapi/models"
)

// Fetch returns a payment resource based on the id
func Fetch(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d := models.FetchResponse{Data: &models.Payment{}}
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
