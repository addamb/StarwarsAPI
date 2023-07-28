package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/addamb/starwarsapi/internal/service"
	"github.com/addamb/starwarsapi/mocks"
	"github.com/addamb/starwarsapi/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_GetHumans(t *testing.T) {
	t.Parallel()
	url := "http://test.com/"
	assert := assert.New(t)

	swapiHuman := models.SwapiHuman{
		Name:   "Luke Skywalker",
		Height: "172",
	}

	swapiResponse := models.SwapiResponse{
		Count:    1,
		Next:     nil,
		Previous: nil,
		Results: []models.SwapiHuman{
			swapiHuman,
		},
	}

	tests := map[string]struct {
		SendSwapiRequestResp  models.SwapiResponse
		SendSwapiRequestError error
		ExpectedError         bool
		RequestResponse       string
	}{
		"error swapi request": {
			SendSwapiRequestError: errors.New("test swapi request error"),
			ExpectedError:         true,
		},
		"error decoder": {
			SendSwapiRequestResp: models.SwapiResponse{},
			ExpectedError:        true,
		},
		"success": {
			SendSwapiRequestResp: swapiResponse,

			ExpectedError:   false,
			RequestResponse: fmt.Sprintf("%+v", swapiResponse.Results[0]),
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := mocks.NewHTTPClient(t)

			swapi := service.NewSwapiClient(client)
			service := service.NewService(swapi, url)

			var out []byte
			if tcase.RequestResponse != "" {
				out, _ = json.Marshal(tcase.SendSwapiRequestResp)
			}

			r := ioutil.NopCloser(bytes.NewReader(out))
			request, _ := http.NewRequest(http.MethodGet, url+"people/?search=Luke", nil)
			client.EXPECT().Do(request).Return(
				&http.Response{
					Body: r,
				},
				tcase.SendSwapiRequestError,
			)

			resp, err := service.GetHumans("Luke")

			assert.Equal(tcase.ExpectedError, err != nil)
			assert.Equal(tcase.SendSwapiRequestResp, resp)

		})
	}
}

func Test_GetPlanet(t *testing.T) {
	t.Parallel()
	url := "http://test.com/"
	assert := assert.New(t)

	swapiResponse := models.SwapiPlanet{
		Name:           "Tatooine",
		RotationPeriod: "23",
		OrbitalPeriod:  "304",
		Diameter:       "10465",
		Climate:        "arid",
		Gravity:        "1 standard",
		Terrain:        "desert",
		SurfaceWater:   "1",
		Population:     "200000",
	}

	tests := map[string]struct {
		SendSwapiRequestResp  models.SwapiPlanet
		SendSwapiRequestError error
		ExpectedError         bool
		RequestResponse       string
		Url                   string
	}{
		"return empty planet": {
			Url:                  "",
			SendSwapiRequestResp: models.SwapiPlanet{},
		},
		"error swapi request": {
			Url:                   url + "planets/1/",
			SendSwapiRequestError: errors.New("test swapi request error"),
			ExpectedError:         true,
		},
		"error decoder": {
			Url:                  url + "planets/1/",
			SendSwapiRequestResp: models.SwapiPlanet{},
			ExpectedError:        true,
		},
		"success": {
			Url:                  url + "planets/1/",
			SendSwapiRequestResp: swapiResponse,

			ExpectedError:   false,
			RequestResponse: fmt.Sprintf("%+v", swapiResponse),
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := mocks.NewHTTPClient(t)

			swapi := service.NewSwapiClient(client)
			service := service.NewService(swapi, url)

			var out []byte
			if tcase.RequestResponse != "" {
				out, _ = json.Marshal(tcase.SendSwapiRequestResp)
			}

			if tcase.Url != "" {
				r := ioutil.NopCloser(bytes.NewReader(out))
				request, _ := http.NewRequest(http.MethodGet, tcase.Url, nil)
				client.EXPECT().Do(request).Return(
					&http.Response{
						Body: r,
					},
					tcase.SendSwapiRequestError,
				)
			}

			resp, err := service.GetPlanet(tcase.Url)

			assert.Equal(tcase.ExpectedError, err != nil)
			assert.Equal(tcase.SendSwapiRequestResp, resp)

		})
	}
}

func Test_GetStarships(t *testing.T) {
	t.Parallel()
	url := "http://test.com/starship/1/"
	assert := assert.New(t)

	swapiStarship := models.SwapiStarship{
		Name:                 "X-wing",
		Model:                "T-65 X-wing",
		Manufacturer:         "Incom Corporation",
		CostInCredits:        "149999",
		Length:               "12.5",
		MaxAtmospheringSpeed: "1050",
		Crew:                 "1",
		Passengers:           "0",
		CargoCapacity:        "110",
		Consumables:          "1 week",
		HyperdriveRating:     "1.0",
		Mglt:                 "100",
		StarshipClass:        "Starfighter",
	}

	swapiResponse := []models.SwapiStarship{
		swapiStarship,
	}

	tests := map[string]struct {
		SendSwapiRequestResp  []models.SwapiStarship
		SendSwapiRequestError error
		ExpectedError         bool
		RequestResponse       string
		Urls                  []string
	}{
		"return empty planet": {
			Urls:                 []string{},
			SendSwapiRequestResp: nil,
		},
		"error swapi request": {
			Urls:                  []string{url},
			SendSwapiRequestError: errors.New("test swapi request error"),
			ExpectedError:         true,
		},
		"error decoder": {
			Urls:                 []string{url},
			SendSwapiRequestResp: nil,
			ExpectedError:        true,
		},
		"success": {
			Urls:                 []string{url},
			SendSwapiRequestResp: swapiResponse,

			ExpectedError:   false,
			RequestResponse: fmt.Sprintf("%+v", swapiResponse),
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := mocks.NewHTTPClient(t)

			swapi := service.NewSwapiClient(client)
			service := service.NewService(swapi, url)

			var out []byte
			if tcase.RequestResponse != "" {
				out, _ = json.Marshal(tcase.SendSwapiRequestResp[0])
			}

			if len(tcase.Urls) != 0 {
				r := ioutil.NopCloser(bytes.NewReader(out))
				request, _ := http.NewRequest(http.MethodGet, tcase.Urls[0], nil)
				client.EXPECT().Do(request).Return(
					&http.Response{
						Body: r,
					},
					tcase.SendSwapiRequestError,
				)
			}

			resp, err := service.GetStarships(tcase.Urls)

			assert.Equal(tcase.ExpectedError, err != nil)
			assert.Equal(tcase.SendSwapiRequestResp, resp)

		})
	}
}

func Test_GetSpecies(t *testing.T) {
	t.Parallel()
	url := "http://test.com/species/2/"
	assert := assert.New(t)

	swapiSpecies := models.SwapiSpecies{
		Name:            "Droid",
		Classification:  "artificial",
		Designation:     "sentient",
		AverageHeight:   "n/a",
		SkinColors:      "n/a",
		HairColors:      "n/a",
		EyeColors:       "n/a",
		AverageLifespan: "indefinite",
		Homeworld:       nil,
		Language:        "n/a",
	}

	swapiResponse := []models.SwapiSpecies{
		swapiSpecies,
	}

	tests := map[string]struct {
		SendSwapiRequestResp  []models.SwapiSpecies
		SendSwapiRequestError error
		ExpectedError         bool
		RequestResponse       string
		Urls                  []string
	}{
		"return empty planet": {
			Urls:                 []string{},
			SendSwapiRequestResp: nil,
		},
		"error swapi request": {
			Urls:                  []string{url},
			SendSwapiRequestError: errors.New("test swapi request error"),
			ExpectedError:         true,
		},
		"error decoder": {
			Urls:                 []string{url},
			SendSwapiRequestResp: nil,
			ExpectedError:        true,
		},
		"success": {
			Urls:                 []string{url},
			SendSwapiRequestResp: swapiResponse,

			ExpectedError:   false,
			RequestResponse: fmt.Sprintf("%+v", swapiResponse),
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := mocks.NewHTTPClient(t)

			swapi := service.NewSwapiClient(client)
			service := service.NewService(swapi, url)

			var out []byte
			if tcase.RequestResponse != "" {
				out, _ = json.Marshal(tcase.SendSwapiRequestResp[0])
			}

			if len(tcase.Urls) != 0 {
				r := ioutil.NopCloser(bytes.NewReader(out))
				request, _ := http.NewRequest(http.MethodGet, tcase.Urls[0], nil)
				client.EXPECT().Do(request).Return(
					&http.Response{
						Body: r,
					},
					tcase.SendSwapiRequestError,
				)
			}

			resp, err := service.GetSpecies(tcase.Urls)

			assert.Equal(tcase.ExpectedError, err != nil)
			assert.Equal(tcase.SendSwapiRequestResp, resp)

		})
	}
}

func Test_GetVehicles(t *testing.T) {
	t.Parallel()
	url := "http://test.com/vehicles/14/"
	assert := assert.New(t)

	swapiSpecies := models.SwapiVehicles{
		Name:                 "Snowspeeder",
		Model:                "t-47 airspeeder",
		Manufacturer:         "Incom corporation",
		CostInCredits:        "unknown",
		Length:               "4.5",
		MaxAtmospheringSpeed: "650",
		Crew:                 "2",
		Passengers:           "0",
		CargoCapacity:        "10",
		Consumables:          "none",
		VehicleClass:         "airspeeder",
	}

	swapiResponse := []models.SwapiVehicles{
		swapiSpecies,
	}

	tests := map[string]struct {
		SendSwapiRequestResp  []models.SwapiVehicles
		SendSwapiRequestError error
		ExpectedError         bool
		RequestResponse       string
		Urls                  []string
	}{
		"return empty planet": {
			Urls:                 []string{},
			SendSwapiRequestResp: nil,
		},
		"error swapi request": {
			Urls:                  []string{url},
			SendSwapiRequestError: errors.New("test swapi request error"),
			ExpectedError:         true,
		},
		"error decoder": {
			Urls:                 []string{url},
			SendSwapiRequestResp: nil,
			ExpectedError:        true,
		},
		"success": {
			Urls:                 []string{url},
			SendSwapiRequestResp: swapiResponse,

			ExpectedError:   false,
			RequestResponse: fmt.Sprintf("%+v", swapiResponse),
		},
	}

	for name, tcase := range tests {
		tcase := tcase //used to avoid lint error 'loop variable tcase captured by func literal'

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			client := mocks.NewHTTPClient(t)

			swapi := service.NewSwapiClient(client)
			service := service.NewService(swapi, url)

			var out []byte
			if tcase.RequestResponse != "" {
				out, _ = json.Marshal(tcase.SendSwapiRequestResp[0])
			}

			if len(tcase.Urls) != 0 {
				r := ioutil.NopCloser(bytes.NewReader(out))
				request, _ := http.NewRequest(http.MethodGet, tcase.Urls[0], nil)
				client.EXPECT().Do(request).Return(
					&http.Response{
						Body: r,
					},
					tcase.SendSwapiRequestError,
				)
			}

			resp, err := service.GetVehicles(tcase.Urls)

			assert.Equal(tcase.ExpectedError, err != nil)
			assert.Equal(tcase.SendSwapiRequestResp, resp)

		})
	}
}
