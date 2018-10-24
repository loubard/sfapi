package payments

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/v1/payments/id", nil)
	w := httptest.NewRecorder()
	Fetch(w, req)

	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Can't read body from get detail response")
	}

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Contains(t, string(body), "{\"data\":")
}
