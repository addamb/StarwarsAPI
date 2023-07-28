package models

type JsonResponse struct {
	Message string `json:"message"`
}

type CharacterInput struct {
	Name string `json:"name"`
}

type CharacterResponse struct {
	Name      string          `json:"name"`
	Height    string          `json:"height"`
	Mass      string          `json:"mass"`
	HairColor string          `json:"hair_color"`
	SkinColor string          `json:"skin_color"`
	EyeColor  string          `json:"eye_color"`
	BirthYear string          `json:"birth_year"`
	Gender    string          `json:"gender"`
	Homeworld SwapiPlanet     `json:"homeworld"`
	Species   []SwapiSpecies  `json:"species"`
	Vehicles  []SwapiVehicles `json:"vehicles"`
	Starships []SwapiStarship `json:"starships"`
}

type SwapiResponse struct {
	Count    int          `json:"count"`
	Next     any          `json:"next"`
	Previous any          `json:"previous"`
	Results  []SwapiHuman `json:"results"`
}

type SwapiHuman struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	URL       string   `json:"url"`
}

type SwapiStarship struct {
	Name          string `json:"name"`
	CargoCapacity string `json:"cargo_capacity"`
	StarshipClass string `json:"starship_class"`
}

type SwapiSpecies struct {
	Name            string `json:"name"`
	AverageLifespan string `json:"average_lifespan"`
	Language        string `json:"language"`
}

type SwapiPlanet struct {
	Name       string `json:"name"`
	Climate    string `json:"climate"`
	Population string `json:"population"`
}

type SwapiVehicles struct {
	Name          string `json:"name"`
	CargoCapacity string `json:"cargo_capacity"`
	VehicleClass  string `json:"vehicle_class"`
}
