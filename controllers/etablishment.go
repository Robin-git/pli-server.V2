package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HandlerGetUsers return all users
func (ctr *Controller) HandlerGetEtablishments(c *gin.Context) {
	// if search
	name := c.DefaultQuery("Name", "")
	if name != "" {
		res, err := ctr.Service.SearchEtablishmentByName(name)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{HResult: res})
		}
		return
	}
	// else
	etablishments, err := ctr.Service.GetEtablishments()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: etablishments})
}

// HandlerGetUser return one user
func (ctr *Controller) HandlerGetEtablishment(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err == nil {
		etablishment, err := ctr.Service.GetEtablishment(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{HResult: etablishment})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{HError: "cannot read paramX"})
		return
	}
	fy, err := strconv.ParseFloat(y, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: "cannot read paramY"})
		return
	}
	fd, err := strconv.ParseFloat(d, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: "cannot read paramDistance"})
		return
	}
	// result
	res, err := ctr.Service.GetDistanceEtablishment(fx, fy, fd)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{HResult: res})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{HError: "error 500"})
	}
}

func (ctr *Controller) HandlerGetAverageNoteEtablishment(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err == nil {
		average, _ := ctr.Service.GetAverageNoteEtablishment(id)
		c.JSON(http.StatusOK, gin.H{HResult: average})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
	}
}
