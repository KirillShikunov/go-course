package app

import (
	"14_layers/internal/api"
	"14_layers/internal/config"
	"14_layers/internal/di"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	container *di.ServiceContainer
}

func (a *App) RegisterRoutes(router *mux.Router) {
	var apis = []api.Routable{
		api.NewOrderAPI(a.container.OrderManager(), a.container.OrderMapper()),
	}

	for _, a := range apis {
		a.RegisterRoutes(router)
	}
}

func (a *App) RunServer(router *mux.Router) {
	a.RegisterRoutes(router)

	port := config.Env("PORT")

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()
	log.Printf("Server started on port %s", port)

	<-ctx.Done()

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server exited properly")
}

func NewApp(container *di.ServiceContainer) *App {
	return &App{container}
}
