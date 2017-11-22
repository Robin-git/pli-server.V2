package api

import (
	"gloo-server/controllers"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter return new mux Router
func NewRouter() *gin.Engine {
	// MariaDB is database of gloo app
	MariaDB := &SQLConnection{
		Host:     os.Getenv("DB_HOST"),
		Login:    os.Getenv("DB_LOGIN"),
		Password: os.Getenv("DB_PWD"),
		Database: os.Getenv("DB_DATABASE"),
		Port:     os.Getenv("DB_PORT"),
	}
	db := MariaDB.GetConnection()

	etablishmentCtrl := controllers.CtrlScopedEtablishment(db)
	userCtrl := controllers.CtrlScopedUser(db)
	opinionCtrl := controllers.CtrlScopedOpinion(db)
	itemCtrl := controllers.CtrlScopedItem(db)
	eventCtrl := controllers.CtrlScopedEvent(db)
	suggestionCtrl := controllers.CtrlScopedSuggestion(db)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		// User
		api.GET("/user", userCtrl.HandlerGetUsers)
		api.GET("/user/:id", userCtrl.HandlerGetUser)

		// Etablishment
		api.GET("/etablishment", etablishmentCtrl.HandlerGetEtablishments)
		api.GET("/etablishment/:id", etablishmentCtrl.HandlerGetEtablishment)
		api.GET("/etablishment/:id/note", etablishmentCtrl.HandlerGetAverageNoteEtablishment)
		api.GET("/distance", etablishmentCtrl.HandlerGetDistanceEtablishment)

		// Option
		api.GET("/opinion", opinionCtrl.HandlerGetOpinions)
		api.POST("/opinion", opinionCtrl.HandlerPostOpinion)

		// Item
		api.GET("/item", itemCtrl.HandlerGetItems)
		api.GET("/item/:id", itemCtrl.HandlerGetItem)
		api.POST("/item", itemCtrl.HandlerPostItem)
		api.PUT("/item/:id", itemCtrl.HandlePutItem)

		// Event
		api.GET("/event", eventCtrl.HandlerGetEvents)
		api.GET("/event/:id", eventCtrl.HandlerGetEvent)
		api.POST("/event", eventCtrl.HandlerPostEvent)
		api.PUT("/event/:id", eventCtrl.HandlePutEvent)

		// Suggestion
		api.GET("/suggestion", suggestionCtrl.HandlerGetSuggestions)
		api.GET("/suggestion/:id", suggestionCtrl.HandlerGetSuggestion)
		api.POST("/suggestion", suggestionCtrl.HandlerPostSuggestion)
		api.PUT("/suggestion/:id", suggestionCtrl.HandlePutSuggestion)
	}

	return r
}
