package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandlerGetUsers return all users
func (ctr *Controller) HandlerGetUsers(c *gin.Context) {
	users, err := ctr.Service.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: users})
}

// HandlerGetUser return one user
func (ctr *Controller) HandlerGetUser(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err == nil {
		user, err := ctr.Service.GetUser(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{HResult: user})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{HError: "bad arguments"})
		return
	}
}

// HandlerAddUser post a new User
func (ctr *Controller) HandlerAddUser(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{HError: http.NotFound})
}

// HandlerDelUser post a new User
func (ctr *Controller) HandlerDelUser(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{HError: http.NotFound})
}
