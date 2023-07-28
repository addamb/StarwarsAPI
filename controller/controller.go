package controller

import (
	"fmt"
	"net/http"
	"sort"

	"github.com/addamb/starwarsapi/internal/service"
	"github.com/addamb/starwarsapi/pkg/models"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Swapi service.SwapiI
}

func NewController(swapi service.SwapiI) *Controller {
	return &Controller{
		Swapi: swapi,
	}
}

func (c *Controller) GetCharacterByNameHandler(e echo.Context) (err error) {
	input := new(models.CharacterInput)

	if err := e.Bind(input); err != nil {
		return e.String(http.StatusBadRequest, "bad request")
	}

	humans, err := c.Swapi.GetHumans(input.Name)
	if err != nil {
		message := fmt.Errorf("could not process request: %w", err)

		return e.String(http.StatusBadRequest, message.Error())
	}

	var characters []models.CharacterResponse

	for _, h := range humans.Results {
		planet, err := c.Swapi.GetPlanet(h.Homeworld)
		if err != nil {
			return err
		}

		species, err := c.Swapi.GetSpecies(h.Species)
		if err != nil {
			return err
		}

		starships, err := c.Swapi.GetStarships(h.Starships)
		if err != nil {
			return err
		}

		vehicles, err := c.Swapi.GetVehicles(h.Vehicles)
		if err != nil {
			return err
		}

		character := models.CharacterResponse{
			Name:      h.Name,
			Height:    h.Height,
			Mass:      h.Mass,
			HairColor: h.HairColor,
			SkinColor: h.SkinColor,
			EyeColor:  h.EyeColor,
			BirthYear: h.BirthYear,
			Gender:    h.Gender,
			Homeworld: planet,
			Species:   species,
			Vehicles:  vehicles,
			Starships: starships,
		}

		characters = append(characters, character)
	}

	if len(characters) > 1 {
		sort.SliceStable(characters, func(i, j int) bool {
			return characters[i].Name < characters[j].Name
		})
	}

	return e.JSON(http.StatusAccepted, characters)
}
