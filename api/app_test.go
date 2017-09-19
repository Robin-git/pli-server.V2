package api_test

import (
	"fmt"
	"net/http/httptest"

	"gloo-server/api"
)

var (
	server   *httptest.Server
	usersUrl string
)

func init() {
	server = httptest.NewServer(api.NewRouter()) //Creating new server with the user handlers

	usersUrl = fmt.Sprintf("%s/api/users", server.URL) //Grab the address for the API endpoint
}
