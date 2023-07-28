package handler

import (
	"github.com/addamb/starwarsapi/controller"
	"github.com/addamb/starwarsapi/internal/service"
)

type Handler struct {
	Controller *controller.Controller
}

func NewHandler(swapi service.SwapiI) *Handler {
	return &Handler{
		Controller: &controller.Controller{
			Swapi: swapi,
		},
	}
}
