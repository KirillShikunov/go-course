package app

import (
	"14_layers/internal/api"
	"14_layers/internal/repositories"
	"14_layers/internal/services"
	"14_layers/internal/services/observers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type App struct {
}

func (a *App) Run(router *mux.Router) {
	log.Println("Starting web server on localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func NewApp(options ...Option) *App {
	app := &App{}

	for _, opt := range options {
		opt(app)
	}
	return app
}

type Option func(*App)

func WithAPIRoutes(router *mux.Router) Option {
	return func(app *App) {
		orderObservers := []services.OrderObserver{
			observers.NewEmailObserver(),
		}
		orderRepository := repositories.NewOrderRepository()
		orderManager := services.NewOrderManager(orderRepository, orderObservers)

		var apis = []api.Routable{
			api.NewOrderApi(orderManager),
		}

		for _, a := range apis {
			a.RegisterRoutes(router)
		}
	}
}
