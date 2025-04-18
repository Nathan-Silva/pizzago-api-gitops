package services

import (
	"github.com/Nathan-Silva/pizzago-api-gitops/models"
	"github.com/Nathan-Silva/pizzago-api-gitops/repository"
)

type PizzaService interface {
	GetAllPizzas() []models.Pizza
	AddPizza(pizza models.Pizza)
}

type pizzaService struct {
	repo repository.PizzaRepository
}

func NewPizzaService(repo repository.PizzaRepository) PizzaService {
	return &pizzaService{repo: repo}
}

func (s *pizzaService) GetAllPizzas() []models.Pizza {
	return s.repo.GetAll()
}

func (s *pizzaService) AddPizza(pizza models.Pizza) {
	s.repo.Add(pizza)
}
