package models

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
	Name                 string `json:"name"`
	Model                string `json:"model"`
	Manufacturer         string `json:"manufacturer"`
	CostInCredits        string `json:"cost_in_credits"`
	Length               string `json:"length"`
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
	Crew                 string `json:"crew"`
	Passengers           string `json:"passengers"`
	CargoCapacity        string `json:"cargo_capacity"`
	Consumables          string `json:"consumables"`
	HyperdriveRating     string `json:"hyperdrive_rating"`
	Mglt                 string `json:"MGLT"`
	StarshipClass        string `json:"starship_class"`
}

type SwapiSpecies struct {
	Name            string `json:"name"`
	Classification  string `json:"classification"`
	Designation     string `json:"designation"`
	AverageHeight   string `json:"average_height"`
	SkinColors      string `json:"skin_colors"`
	HairColors      string `json:"hair_colors"`
	EyeColors       string `json:"eye_colors"`
	AverageLifespan string `json:"average_lifespan"`
	Homeworld       any    `json:"homeworld"`
	Language        string `json:"language"`
}

type SwapiPlanet struct {
	Name           string `json:"name"`
	RotationPeriod string `json:"rotation_period"`
	OrbitalPeriod  string `json:"orbital_period"`
	Diameter       string `json:"diameter"`
	Climate        string `json:"climate"`
	Gravity        string `json:"gravity"`
	Terrain        string `json:"terrain"`
	SurfaceWater   string `json:"surface_water"`
	Population     string `json:"population"`
}

type SwapiVehicles struct {
	Name                 string `json:"name"`
	Model                string `json:"model"`
	Manufacturer         string `json:"manufacturer"`
	CostInCredits        string `json:"cost_in_credits"`
	Length               string `json:"length"`
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
	Crew                 string `json:"crew"`
	Passengers           string `json:"passengers"`
	CargoCapacity        string `json:"cargo_capacity"`
	Consumables          string `json:"consumables"`
	VehicleClass         string `json:"vehicle_class"`
}
