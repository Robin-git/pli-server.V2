package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// HandlerGetUsers return all users
func (c *Controller) HandlerGetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := c.Service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
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
		res, err := c.Service.GetUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
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

// HandlerAddUser post a new User
func (c *Controller) HandlerAddUser(w http.ResponseWriter, r *http.Request) {
	res, err := c.Service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	users, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(users)
}
