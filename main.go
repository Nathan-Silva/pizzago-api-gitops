package main

import (
	"github.com/Nathan-Silva/pizzago-api-gitops/handlers"
	"github.com/Nathan-Silva/pizzago-api-gitops/repository"
	"github.com/Nathan-Silva/pizzago-api-gitops/routes"
	"github.com/Nathan-Silva/pizzago-api-gitops/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := repository.NewInMemoryPizzaRepository()
	service := services.NewPizzaService(repo)
	handler := handlers.NewPizzaHandler(service)

	routes.SetupRoutes(router, handler)

	router.Run()
}
