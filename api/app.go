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

	r := gin.Default()

	r.GET("/api/user", userCtrl.HandlerGetUsers)
	r.GET("/api/user/:id", userCtrl.HandlerGetUser)

	r.GET("/api/etablishment", etablishmentCtrl.HandlerGetEtablishments)
	r.GET("/api/etablishment/:id", etablishmentCtrl.HandlerGetEtablishment)
	r.GET("/api/etablishment/:id/note", etablishmentCtrl.HandlerGetAverageNoteEtablishment)
	r.GET("/api/distance", etablishmentCtrl.HandlerGetDistanceEtablishment)

	r.GET("/api/opinion", opinionCtrl.HandlerGetOpinions)
	r.POST("/api/opinion", opinionCtrl.HandlerPostOpinion)

	r.GET("/api/item", itemCtrl.HandlerGetItems)
	r.GET("/api/item/:id", itemCtrl.HandlerGetItem)
	r.POST("/api/item", itemCtrl.HandlerPostItem)
	r.PUT("/api/item/:id", itemCtrl.HandlePutItem)

	return r
}
