package main

import (
	"github.com/Nathan-Silva/infrastructure-golang-api/handlers"
	"github.com/Nathan-Silva/infrastructure-golang-api/repository"
	"github.com/Nathan-Silva/infrastructure-golang-api/routes"
	"github.com/Nathan-Silva/infrastructure-golang-api/services"
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
