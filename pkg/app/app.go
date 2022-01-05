package app

import (
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/server"
)

type Application struct {
	config *config.Config
	server *server.Server
}

func NewApplication(config *config.Config) *Application {
	return &Application{
		config: config,
	}
}

func (a *Application) Run() {
	println("App starting...")

	a.server = server.NewServer(&a.config.Tls, &a.config.App)
	go a.server.Run()
}

func (a *Application) Shutdown() {
	println("\nApp shutting down...")

	// _ = a.server.Shutdown(context.Background())

	println("OFF")
}
