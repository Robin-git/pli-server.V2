package controllers

import (
	"log"
	"net/http"
	"strconv"

	"gloo-server/models"

	"github.com/gin-gonic/gin"
)

func (ctr *Controller) HandlerGetOpinions(c *gin.Context) {
	id := c.Query("id_etablishment")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_etablishment is required"})
		return
	}
	if id, err := strconv.Atoi(id); err == nil {
		opinions, _ := ctr.Service.GetOpinion(id)
		c.JSON(http.StatusOK, gin.H{"results": opinions})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_etablishment bad format"})
		return
	}
}

func (ctr *Controller) HandlerPostOpinion(c *gin.Context) {
	var json models.Opinion
	if err := c.BindJSON(&json); err == nil {
		log.Println(json)
		if _, err := ctr.Service.GetUser(int(json.UserID)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}
		if _, err := ctr.Service.GetEtablishment(int(json.EtablishmentID)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Etablishment not found"})
			return
		}
		if err := ctr.Service.PostOpinion(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": json})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
