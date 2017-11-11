package controllers

import (
	"net/http"
	"strconv"

	"gloo-server/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlScopedEtablishment is initialize new Controller
func CtrlScopedEtablishment(db *gorm.DB) *CtrlEtablsihment {
	return &CtrlEtablsihment{
		Service: &models.ServiceEtablishment{
			DB: db,
		},
	}
}

// CtrlEtablsihment is controller of Etablisment
type CtrlEtablsihment struct {
	Service *models.ServiceEtablishment
}

// HandlerGetEtablishments return all Etablishments
func (ctr *CtrlEtablsihment) HandlerGetEtablishments(c *gin.Context) {
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

// HandlerGetEtablishment return one Etablishment
func (ctr *CtrlEtablsihment) HandlerGetEtablishment(c *gin.Context) {
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

// HandlerGetDistanceEtablishment return distance of user and x & y & distance
func (ctr *CtrlEtablsihment) HandlerGetDistanceEtablishment(c *gin.Context) {
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

// HandlerGetAverageNoteEtablishment return average note of one Etablishment
func (ctr *CtrlEtablsihment) HandlerGetAverageNoteEtablishment(c *gin.Context) {
	id := c.Param("id")
	if id, err := strconv.Atoi(id); err == nil {
		average, _ := ctr.Service.GetAverageNoteEtablishment(id)
		c.JSON(http.StatusOK, gin.H{HResult: average})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
	}
}
