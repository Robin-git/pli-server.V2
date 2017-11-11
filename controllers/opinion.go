package controllers

import (
	"log"
	"net/http"
	"strconv"

	"gloo-server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlScopedOpinion is initialize new Controller
func CtrlScopedOpinion(db *gorm.DB) *CtrlOpinion {
	return &CtrlOpinion{
		Service: &models.ServiceOpinion{
			DB: db,
		},
	}
}

// CtrlOpinion is controller of Opinion
type CtrlOpinion struct {
	Service *models.ServiceOpinion
}

// HandlerGetOpinions return multiple opinions
func (ctr *CtrlOpinion) HandlerGetOpinions(c *gin.Context) {
	id := c.Query("id_etablishment")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{HError: "id_etablishment is required"})
		return
	}
	if id, err := strconv.Atoi(id); err == nil {
		opinions, _ := ctr.Service.GetOpinion(uint(id))
		c.JSON(http.StatusOK, gin.H{HResult: opinions})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{HError: "id_etablishment bad format"})
		return
	}
}

// HandlerPostOpinion post one opinion
func (ctr *CtrlOpinion) HandlerPostOpinion(c *gin.Context) {
	var json models.Opinion
	if err := c.BindJSON(&json); err == nil {
		log.Println(json)
		if err := ctr.Service.PostOpinion(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{HError: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{HResult: "OK"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
	}
}
