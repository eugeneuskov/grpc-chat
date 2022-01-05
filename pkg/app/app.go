package app

import "github.com/eugeneuskov/grpc-chat/config"

type Application struct {
	config *config.Config
}

func NewApplication(config *config.Config) *Application {
	return &Application{
		config,
	}
}

func (a *Application) Shutdown() {
	println("OFF")
}
