package controllers

import (
	"gloo-server/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlScopedSuggestion is initialize new Controller
func CtrlScopedSuggestion(db *gorm.DB) *CtrlSuggestion {
	return &CtrlSuggestion{
		Service: &models.ServiceSuggestion{
			DB: db,
		},
	}
}

// CtrlSuggestion is controller of Suggestion
type CtrlSuggestion struct {
	Service *models.ServiceSuggestion
}

func buildQueryParameter(c *gin.Context) *models.QueryParameterSuggestion {
	qp := &models.QueryParameterSuggestion{}
	// with etablishment
	we := c.Query("with")
	paramlist := strings.Split(we, ",")
	for _, param := range paramlist {
		switch param {
		case "etablishment":
			qp.Etablishment = true
		case "user":
			qp.User = true
		case "item":
			qp.Item = true
		}
	}
	return qp
}

// HandlerGetSuggestion return one Suggestion
// Params { id }
func (ctr *CtrlSuggestion) HandlerGetSuggestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	qp := buildQueryParameter(c)
	suggestion, err := ctr.Service.GetSuggestion(id, qp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: suggestion})
}

// HandlerGetSuggestions return all Suggestion
func (ctr *CtrlSuggestion) HandlerGetSuggestions(c *gin.Context) {
	qp := buildQueryParameter(c)
	queryID := c.Query("id_etablishment")
	ide, err := strconv.Atoi(queryID)
	if queryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{HError: "id_etablishment is required"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: "id_etablishment bad format"})
		return
	}
	suggestions, err := ctr.Service.GetSuggestions(qp, uint(ide))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: suggestions})
}

// HandlerPostSuggestion post one Suggestion
// Query { Suggestion }
func (ctr *CtrlSuggestion) HandlerPostSuggestion(c *gin.Context) {
	var json struct {
		Name           string `json:"name" binding:"required"`
		Description    string `json:"description" binding:"required"`
		EtablishmentID uint   `json:"etablishment_id" binding:"required"`
		UserID         uint   `json:"user_id" binding:"required"`
	}
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	if err := ctr.Service.PostSuggestion(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: "OK"})
}

// HandlePutSuggestion update one suggestion
// Query { Suggestion }
func (ctr *CtrlSuggestion) HandlePutSuggestion(c *gin.Context) {
	var json *models.Suggestion
	// Convert id string to id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	qp := buildQueryParameter(c)
	// Get one suggestion by id
	suggestion, err := ctr.Service.GetSuggestion(id, qp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	// Convert query by json
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	// Update Suggestion
	if err := ctr.Service.UpdateSuggestion(suggestion, json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	// Return OK
	c.JSON(http.StatusOK, gin.H{HResult: suggestion})
}

// HandlerDeleteSuggestion delete one Suggestion
// Params { id }
func (ctr *CtrlSuggestion) HandlerDeleteSuggestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	err = ctr.Service.DeleteSuggestion(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: "OK"})
}
