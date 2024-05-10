package server

import (
	"21_uber-fx/internal/api"
	"21_uber-fx/internal/config"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"log"
	"net/http"
	"time"
)

type Server struct {
	handlers []api.Routable `group:"api.handlers"`
	router   *mux.Router
	server   *http.Server
}

func (s *Server) Start() error {
	for _, handler := range s.handlers {
		handler.RegisterRoutes(s.router)
	}

	port := config.Env("PORT")
	s.server = &http.Server{
		Addr:         ":" + port,
		Handler:      s.router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	log.Printf("Server started on port %s", port)

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	fmt.Println("Server has been shutdown gracefully.")

	return nil
}

func Run(lifecycle fx.Lifecycle, s *Server) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return s.Start()
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			return s.Shutdown(ctx)
		},
	})
}

func NewServer(apis []api.Routable, router *mux.Router) *Server {
	return &Server{handlers: apis, router: router}
}
