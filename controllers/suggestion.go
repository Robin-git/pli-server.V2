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
			qp.WEtablishment = true
		case "user":
			qp.WUser = true
		case "item":
			qp.WItem = true
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
	suggestions, err := ctr.Service.GetSuggestions(qp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: suggestions})
}

// HandlerPostSuggestion post one Suggestion
// Query { Suggestion }
func (ctr *CtrlSuggestion) HandlerPostSuggestion(c *gin.Context) {
	var json *models.Suggestion
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
