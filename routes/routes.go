package routes

import (
	"github.com/Nathan-Silva/pizzago-api-gitops/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, pizzaHandler *handlers.PizzaHandler) {
	r.GET("/pizzas", pizzaHandler.GetPizzas)
	r.POST("/pizzas", pizzaHandler.PostPizza)
	r.GET("/healthz", handlers.GetHealth)
	r.GET("/readyz", handlers.GetHealth)
}
