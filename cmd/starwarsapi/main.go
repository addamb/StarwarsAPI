package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/addamb/starwarsapi/internal/handler"
	"github.com/addamb/starwarsapi/internal/service"
	"github.com/labstack/echo/v4"
)

const (
	defaultPort = "8080"
	baseURL     = "https://swapi.dev/api/"
)

func main() {
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if port == ":" {
		port = fmt.Sprintf(":%s", defaultPort)
	}

	swapi := service.NewSwapiClient(&http.Client{})
	service := service.NewService(swapi, baseURL)
	h := handler.NewHandler(&service)

	e := echo.New()
	api := e.Group("/api")

	h.RegisterRoutes(api)

	e.Logger.Fatal(e.Start(port))
}
