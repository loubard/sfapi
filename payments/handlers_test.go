package payments

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/loubard/sfapi/models"
	"github.com/loubard/sfapi/sql"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	req := httptest.NewRequest(
		"GET",
		"http://example.com/v1/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		nil,
	)
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

func TestList(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	req := httptest.NewRequest("GET", "http://example.com/v1/payments/", nil)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/payments/", List(db))
	router.ServeHTTP(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, string(body), "{\"data\":[")
	d := &models.ListResponse{}
	json.Unmarshal(body, d)
	assert.Equal(t, 2, len(*d.Data))
}

func TestDelete(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	req := httptest.NewRequest(
		"DELETE",
		"http://example.com/v1/payments/4ee3a8d8-ca7b-4290-a52c-dd5b6165ec43",
		nil,
	)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/payments/{id}", Delete(db))
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCreate(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	json := `{"payment_id":"42","attributes":{"amount":"100"}}`

	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/payments/",
		strings.NewReader(json),
	)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/payments/", Create(db))
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	p := &models.Payment{}

	db.Where("payments.payment = ?", "42").First(&p)
	assert.Equal(t, "42", p.Payment)
}

func TestCreateBadJSON(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	json := `{`

	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/payments/",
		strings.NewReader(json),
	)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/payments/", Create(db))
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateInvalid(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	assert.NoError(t, err)
	sql.Seed(db)

	json := `{"payment_id": 42}`

	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/payments/",
		strings.NewReader(json),
	)
	w := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/v1/payments/", Create(db))
	router.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
