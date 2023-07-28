package service

import (
	"encoding/json"
	"fmt"

	"github.com/addamb/starwarsapi/pkg/models"
)

type SwapiI interface {
	GetHumans(name string) (models.SwapiResponse, error)
	GetPlanet(url string) (models.SwapiPlanet, error)
	GetStarships(urls []string) ([]models.SwapiStarship, error)
	GetSpecies(urls []string) ([]models.SwapiSpecies, error)
	GetVehicles(urls []string) ([]models.SwapiVehicles, error)
}

type Swapi struct {
	SwapiClient
	BaseURL string
}

func NewService(swapi SwapiClient, url string) Swapi {
	return Swapi{
		SwapiClient: swapi,
		BaseURL:     url,
	}
}

func (s *Swapi) GetHumans(name string) (models.SwapiResponse, error) {
	humans := models.SwapiResponse{}
	//Build person endpoint
	person := fmt.Sprintf("%speople/?search=%s", s.BaseURL, name)

	resp, err := s.SendSwapiRequest(person)
	if err != nil {
		return humans, err
	}

	if err := json.NewDecoder(resp).Decode(&humans); err != nil {
		return humans, err
	}

	return humans, nil
}

func (s *Swapi) GetPlanet(url string) (models.SwapiPlanet, error) {
	//If no url return empty
	if url == "" {
		return models.SwapiPlanet{}, nil
	}

	resp, err := s.SendSwapiRequest(url)
	if err != nil {
		return models.SwapiPlanet{}, err
	}

	planet := models.SwapiPlanet{}

	if err := json.NewDecoder(resp).Decode(&planet); err != nil {
		return models.SwapiPlanet{}, err
	}

	return planet, nil
}

func (s *Swapi) GetStarships(urls []string) ([]models.SwapiStarship, error) {
	var starships []models.SwapiStarship

	//If no urls return empty
	if len(urls) == 0 {
		return starships, nil
	}

	for _, u := range urls {
		resp, err := s.SendSwapiRequest(u)
		if err != nil {
			return starships, err
		}

		starship := models.SwapiStarship{}

		if err := json.NewDecoder(resp).Decode(&starship); err != nil {
			return starships, err
		}

		starships = append(starships, starship)
	}

	return starships, nil
}

func (s *Swapi) GetSpecies(urls []string) ([]models.SwapiSpecies, error) {
	var species []models.SwapiSpecies

	//If no urls return empty
	if len(urls) == 0 {
		return species, nil
	}

	for _, u := range urls {
		resp, err := s.SendSwapiRequest(u)
		if err != nil {
			return species, err
		}

		specie := models.SwapiSpecies{}

		if err := json.NewDecoder(resp).Decode(&specie); err != nil {
			return species, err
		}

		species = append(species, specie)
	}

	return species, nil
}

func (s *Swapi) GetVehicles(urls []string) ([]models.SwapiVehicles, error) {
	var vehicles []models.SwapiVehicles

	//If no urls return empty
	if len(urls) == 0 {
		return vehicles, nil
	}

	for _, u := range urls {
		resp, err := s.SendSwapiRequest(u)
		if err != nil {
			return vehicles, err
		}

		vehicle := models.SwapiVehicles{}

		if err := json.NewDecoder(resp).Decode(&vehicle); err != nil {
			return vehicles, err
		}

		vehicles = append(vehicles, vehicle)
	}

	return vehicles, nil
}
