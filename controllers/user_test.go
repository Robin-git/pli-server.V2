package controllers_test

import (
	"fmt"
	"gloo-server/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	a = api.NewRouter()
)

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.ServeHTTP(rr, req)
	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected response code %d. Got %d\n", expected, actual))
}

func TestHandlerGetUsers(t *testing.T) {
	// userJson := `{"FirstName": "Go", "LastName": "Gi", "Email": "go.go@gmail.com", Role: 1}`
	req, _ := http.NewRequest("GET", "/users/11", nil)
	response := ExecuteRequest(req)
	CheckResponseCode(t, http.StatusNotFound, response.Code)
	assert.Equal(t, response.Body.String(), "404 page not found\n", fmt.Sprintf("Expected the 'error' key of the response to be set to 'Product not found'. Got '%s'", response.Body.String()))
}
