package routes

import (
	"github.com/Nathan-Silva/infrastructure-golang-api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, pizzaHandler *handlers.PizzaHandler) {
	r.GET("/pizzas", pizzaHandler.GetPizzas)
	r.POST("/pizzas", pizzaHandler.PostPizza)
}
