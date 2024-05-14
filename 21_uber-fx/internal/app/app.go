package app

import (
	"21_uber-fx/internal/api"
	"21_uber-fx/internal/config"
	"21_uber-fx/internal/db"
	"21_uber-fx/internal/mapper"
	"21_uber-fx/internal/repositories"
	"21_uber-fx/internal/server"
	"21_uber-fx/internal/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
)

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return nil
}

func ProvideRegisters() fx.Option {
	if err := loadEnv(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}

	return fx.Provide(
		config.NewDBConfig,
		db.NewConnection,

		repositories.NewOrderRepository,
		fx.Annotate(repositories.NewOrderRepository, fx.As(new(services.Repository))),

		services.NewOrderManager,
		mapper.NewOrderMapper,

		fx.Annotate(services.NewOrderManager, fx.As(new(api.OrderManager))),

		mux.NewRouter,

		fx.Annotate(api.NewOrderAPI, fx.ResultTags(`group:"api.handlers"`), fx.As(new(api.Routable))),
		fx.Annotate(server.NewServer, fx.ParamTags(`group:"api.handlers"`)),
	)
}
