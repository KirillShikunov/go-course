package api

import (
	"14_layers/internal/repositories"
	"14_layers/internal/services"
	"14_layers/internal/services/observers"
	"github.com/gorilla/mux"
)

type Routable interface {
	RegisterRoutes(router *mux.Router)
}

func RegisterRoutes(router *mux.Router) {
	orderObservers := []services.OrderObserver{
		observers.NewEmailObserver(),
	}
	orderRepository := repositories.NewOrderRepository()
	orderManager := services.NewOrderManager(orderRepository, orderObservers)

	var apis = []Routable{
		NewOrderAPI(orderManager),
	}

	for _, a := range apis {
		a.RegisterRoutes(router)
	}
}
