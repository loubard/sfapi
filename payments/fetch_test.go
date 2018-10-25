package payments

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/loubard/sfapi/sql"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	req := httptest.NewRequest("GET", "http://example.com/v1/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/payments/{id}", Fetch(db))
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, string(body), "{\"data\":")
}

func TestFetchNotFound(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	req := httptest.NewRequest("GET", "http://example.com/v1/payments/not-found", nil)
	w := httptest.NewRecorder()

	Fetch(db)(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}
