package main

import (
	"21_uber-fx/internal/app"
	"21_uber-fx/internal/server"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		app.ProvideRegisters(),
		fx.Invoke(
			server.Run,
		),
	).Run()
}
