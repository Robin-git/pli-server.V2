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

func (ctr *Controller) HandlerPostOpinion(c *gin.Context) {
	var json models.Opinion
	if err := c.BindJSON(&json); err == nil {
		log.Println(json)
		if err := ctr.Service.PostOpinion(json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{HError: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{HResult: "OK"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{HError: err.Error()})
	}
}
