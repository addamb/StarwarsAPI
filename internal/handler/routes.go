package handler

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRoutes(api *echo.Group) {
	//Character group
	character := api.Group("/character")
	character.GET("", h.Controller.GetCharacterByNameHandler)

	//Could add more routes to get specific starships or to translate into Wookie
}
