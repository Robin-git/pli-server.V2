package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandlerGetUsers return all users
func (ctr *Controller) HandlerGetEtablishments(c *gin.Context) {
	etablishments, err := ctr.Service.GetEtablishments()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": etablishments})
}

// HandlerGetUser return one user
func (ctr *Controller) HandlerGetEtablishment(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err == nil {
		etablishment, err := ctr.Service.GetEtablishment(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"results": etablishment})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (ctr *Controller) HandlerGetDistanceEtablishment(c *gin.Context) {
	// params
	x := c.Query("paramX")
	y := c.Query("paramY")
	d := c.Query("paramDistance")
	// parse params
	fx, err := strconv.ParseFloat(x, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read paramX"})
		return
	}
	fy, err := strconv.ParseFloat(y, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read paramY"})
		return
	}
	fd, err := strconv.ParseFloat(d, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot read paramDistance"})
		return
	}
	// results
	res, err := ctr.Service.GetDistanceEtablishment(fx, fy, fd)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"results": res})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error 500"})
	}
}
