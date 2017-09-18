package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HandlerGetUsers return all users
func (c *Controller) HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	res := c.Service.GetUsers()
	users, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(users)
}

// HandlerGetUser return one user
func (c *Controller) HandlerGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id, err := strconv.Atoi(id); err == nil {
		res := c.Service.GetUser(id)
		user, err := json.Marshal(res)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(user)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
