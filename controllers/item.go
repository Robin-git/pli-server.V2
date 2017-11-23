package controllers

import (
	"gloo-server/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CtrlScopedItem is initialize new Controller
func CtrlScopedItem(db *gorm.DB) *CtrlItem {
	return &CtrlItem{
		Service: &models.ServiceItem{
			DB: db,
		},
	}
}

// CtrlItem is controller of Item
type CtrlItem struct {
	Service *models.ServiceItem
}

// HandlerGetItem return one Item
// Params { id }
func (ctr *CtrlItem) HandlerGetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	item, err := ctr.Service.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: item})
}

// HandlerGetItems return all Item
func (ctr *CtrlItem) HandlerGetItems(c *gin.Context) {
	id := c.Query("id_etablishment")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{HError: "id_etablishment is required"})
		return
	}
	idconv, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: "id_etablishment is malformed"})
		return
	}
	items, err := ctr.Service.GetItems(uint(idconv))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: items})
}

// HandlerPostItem post one Item
// Query { Item }
func (ctr *CtrlItem) HandlerPostItem(c *gin.Context) {
	var json *models.Item
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	if err := ctr.Service.PostItem(json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{HResult: "OK"})
}

// HandlePutItem update one item
// Query { Item }
func (ctr *CtrlItem) HandlePutItem(c *gin.Context) {
	var json *models.Item
	// Convert id string to id int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{HError: err.Error()})
		return
	}
	// Get one item by id
	item, err := ctr.Service.GetItem(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{HError: err.Error()})
		return
	}
	// Convert query by json
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	// Update Item
	if err := ctr.Service.UpdateItem(item, json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{HError: http.StatusBadRequest})
		return
	}
	// Return OK
	c.JSON(http.StatusOK, gin.H{HResult: item})
}
