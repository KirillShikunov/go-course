package api

import (
	"14_layers/internal/db"
	"14_layers/internal/mail"
	"14_layers/internal/mapper"
	"14_layers/internal/repositories"
	"14_layers/internal/services"
	"14_layers/internal/services/observers"
	"github.com/gorilla/mux"
)

type Routable interface {
	RegisterRoutes(router *mux.Router)
}

func RegisterRoutes(router *mux.Router) {
	sender := mail.NewEmailSender()
	orderObservers := []services.OrderObserver{
		observers.NewEmailObserver(sender),
	}

	connection := db.GetConnection()
	orderRepository := repositories.NewOrderRepository(connection)

	orderManager := services.NewOrderManager(orderRepository, orderObservers)
	orderMapper := mapper.NewOrderMapper()

	var apis = []Routable{
		NewOrderAPI(orderManager, orderMapper),
	}

	for _, a := range apis {
		a.RegisterRoutes(router)
	}
}
