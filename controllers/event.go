package controllers

import (
	"gloo-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlScopedEvent is initialize new Controller
func CtrlScopedEvent(db *gorm.DB) *CtrlEvent {
	return &CtrlEvent{
		Service: &models.ServiceEvent{
			DB: db,
		},
	}
}

// CtrlEvent is controller of Event
type CtrlEvent struct {
	Service *models.ServiceEvent
}

// HandlerGetEvent return one Event
// Params { id }
func (ctr *CtrlEvent) HandlerGetEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	qp := &models.QueryParameterEvent{}
	// with etablishment
	we := c.Query("with")
	if we == "etablishment" {
		qp.Wetablishment = true
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	event, err := ctr.Service.GetEvent(id, qp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: event})
}

// HandlerGetEvents return all Event
func (ctr *CtrlEvent) HandlerGetEvents(c *gin.Context) {
	qp := &models.QueryParameterEvent{}
	// with etablishment
	we := c.Query("with")
	if we == "etablishment" {
		qp.Wetablishment = true
	}
	events, err := ctr.Service.GetEvents(qp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: events})
}

// HandlerPostEvent post one Event
// Query { Event }
func (ctr *CtrlEvent) HandlerPostEvent(c *gin.Context) {
	var json *models.Event
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	if err := ctr.Service.PostEvent(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: "OK"})
}

// HandlePutEvent update one event
// Query { Event }
func (ctr *CtrlEvent) HandlePutEvent(c *gin.Context) {
	var json *models.Event
	// Convert id string to id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	// Get one event by id
	event, err := ctr.Service.GetEvent(id, nil)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	// Convert query by json
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	// Update Event
	if err := ctr.Service.UpdateEvent(event, json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	// Return OK
	c.JSON(http.StatusOK, gin.H{HResult: event})
}
