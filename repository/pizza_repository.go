package repository

import "github.com/Nathan-Silva/pizzago-api-gitops/models"

type PizzaRepository interface {
	GetAll() []models.Pizza
	Add(pizza models.Pizza)
}

type InMemoryPizzaRepository struct {
	pizzas []models.Pizza
}

func NewInMemoryPizzaRepository() *InMemoryPizzaRepository {
	return &InMemoryPizzaRepository{
		pizzas: []models.Pizza{
			{Id: 1, Nome: "Toscana", Preco: 49.5},
			{Id: 2, Nome: "Calabresa", Preco: 28.5},
			{Id: 3, Nome: "Frango", Preco: 70.5},
		},
	}
}

func (r *InMemoryPizzaRepository) GetAll() []models.Pizza {
	return r.pizzas
}

func (r *InMemoryPizzaRepository) Add(pizza models.Pizza) {
	r.pizzas = append(r.pizzas, pizza)
}
