package controllers

import (
	"shopingCar_go/services"
)

type Controller struct {
	service *services.Service
}

func NewController(service *services.Service) *Controller {
	return &Controller{service}

}
