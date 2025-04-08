package handlers

import (
	"net/http"

	"github.com/Nathan-Silva/infrastructure-golang-api/models"
	"github.com/Nathan-Silva/infrastructure-golang-api/services"
	"github.com/gin-gonic/gin"
)

type PizzaHandler struct {
	service services.PizzaService
}

func NewPizzaHandler(service services.PizzaService) *PizzaHandler {
	return &PizzaHandler{service: service}
}

func (h *PizzaHandler) GetPizzas(c *gin.Context) {
	pizzas := h.service.GetAllPizzas()
	c.JSON(http.StatusOK, gin.H{"pizzas": pizzas})
}

func (h *PizzaHandler) PostPizza(c *gin.Context) {
	var pizza models.Pizza
	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
		return
	}
	h.service.AddPizza(pizza)
	c.JSON(http.StatusCreated, pizza)
}
