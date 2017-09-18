package controllers

import (
	"gloo-server/models"
)

// Controller is a struct of controllers
type Controller struct {
	Service *models.Service
}

// ControllerScoped is initialize new Controller
func ControllerScoped(s *models.Service) *Controller {
	return &Controller{
		Service: s,
	}
}
