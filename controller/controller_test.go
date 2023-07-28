package controller_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/addamb/starwarsapi/internal/handler"
	"github.com/addamb/starwarsapi/mocks"
	"github.com/addamb/starwarsapi/pkg/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	userJSON = `{"name":"Anakin Skywalker"}`
)

func Test_GetCharacterByNameHandler(t *testing.T) {
	t.Parallel()
	url := "http://test.com/"
	assert := assert.New(t)

	swapiHuman := models.SwapiHuman{
		Name:      "Luke Skywalker",
		Height:    "172",
		Homeworld: url,
		Species:   []string{url},
		Vehicles:  []string{url},
		Starships: []string{url},
	}

	swapiResponse := models.SwapiResponse{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []models.SwapiHuman{
			swapiHuman,
		},
	}

	characterResp := models.CharacterResponse{
		Name:      "Luke Skywalker",
		Height:    "172",
		Homeworld: models.SwapiPlanet{},
		Species:   []models.SwapiSpecies{},
		Vehicles:  []models.SwapiVehicles{},
		Starships: []models.SwapiStarship{},
	}

	expectedResponse := []models.CharacterResponse{
		characterResp,
	}

	expectedJson, _ := json.Marshal(expectedResponse)

	tests := map[string]struct {
		CharacterJSON     string
		GetHumansResp     models.SwapiResponse
		GetHumansError    error
		GetPlanetResp     models.SwapiPlanet
		GetPlanetError    error
		GetSpeciesResp    []models.SwapiSpecies
		GetSpeciesError   error
		GetStarshipsResp  []models.SwapiStarship
		GetStarshipsError error
		GetVehiclesResp   []models.SwapiVehicles
		GetVehiclesError  error
		ExpectedStatus    int
		ExpectedResult    string
		ExpectError       bool
	}{
		"error swapi human request": {
			CharacterJSON:  userJSON,
			GetHumansError: errors.New("test error client"),
			ExpectedStatus: http.StatusBadRequest,
			ExpectedResult: "could not process request: test error client",
		},
		"error swapi planet request": {
			CharacterJSON: userJSON,
			GetHumansResp: swapiResponse,
			GetPlanetResp: models.SwapiPlanet{
				Name: "Test Planet",
			},
			GetPlanetError: errors.New("test error client"),
			ExpectedStatus: http.StatusOK,
			ExpectedResult: "",
			ExpectError:    true,
		},
		"success": {
			CharacterJSON: userJSON,
			GetHumansResp: swapiResponse,
			GetPlanetResp: models.SwapiPlanet{
				Name: "Test Planet",
			},
			GetSpeciesResp: []models.SwapiSpecies{
				models.SwapiSpecies{Name: "unknown"},
			},
			GetStarshipsResp: []models.SwapiStarship{
				models.SwapiStarship{Name: "X-Wing"},
			},
			GetVehiclesResp: []models.SwapiVehicles{
				models.SwapiVehicles{Name: "N/A"},
			},
			ExpectedStatus: http.StatusAccepted,
			ExpectedResult: string(expectedJson) + "\n",
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockService := mocks.NewSwapiI(t)
			h := handler.NewHandler(mockService)

			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(tcase.CharacterJSON))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			//if (tcase.GetHumansResp != models.SwapiResponse{}) {
			mockService.EXPECT().GetHumans("Anakin Skywalker").Return(
				tcase.GetHumansResp,
				tcase.GetHumansError,
			)

			//}

			if (tcase.GetPlanetResp != models.SwapiPlanet{}) {
				mockService.EXPECT().GetPlanet(url).Return(
					models.SwapiPlanet{},
					tcase.GetPlanetError,
				)
			}

			if len(tcase.GetSpeciesResp) != 0 {
				mockService.EXPECT().GetSpecies([]string{url}).Return(
					[]models.SwapiSpecies{},
					tcase.GetSpeciesError,
				)
			}

			if len(tcase.GetStarshipsResp) != 0 {
				mockService.EXPECT().GetStarships([]string{url}).Return(
					[]models.SwapiStarship{},
					tcase.GetStarshipsError,
				)
			}

			if len(tcase.GetVehiclesResp) != 0 {
				mockService.EXPECT().GetVehicles([]string{url}).Return(
					[]models.SwapiVehicles{},
					tcase.GetVehiclesError,
				)
			}

			err := h.Controller.GetCharacterByNameHandler(c)
			assert.Equal(tcase.ExpectError, err != nil)
			assert.Equal(tcase.ExpectedStatus, rec.Code)
			if len(rec.Body.Bytes()) > 0 {
				assert.Equal(tcase.ExpectedResult, rec.Body.String())
			}
		})
	}
}
