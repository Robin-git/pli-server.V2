package controllers

import (
	"gloo-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlScopedUser is initialize new Controller
func CtrlScopedUser(db *gorm.DB) *CtrlUser {
	return &CtrlUser{
		Service: &models.ServiceUser{
			DB: db,
		},
	}
}

// CtrlUser is controller of User
type CtrlUser struct {
	Service *models.ServiceUser
}

// HandlerGetUsers return all users
func (ctr *CtrlUser) HandlerGetUsers(c *gin.Context) {
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
func (ctr *CtrlUser) HandlerGetUser(c *gin.Context) {
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
func (ctr *CtrlUser) HandlerAddUser(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{HError: http.NotFound})
}

// HandlerDelUser post a new User
func (ctr *CtrlUser) HandlerDelUser(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{HError: http.NotFound})
}
