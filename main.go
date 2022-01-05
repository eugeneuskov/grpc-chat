package main

import (
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfig, err := new(config.Config).Init()
	if err != nil {
		log.Fatalf("Failed initializing config: %s\n", err.Error())
		return
	}

	application := app.NewApplication(appConfig)

	println("App started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	println("\nApp shutting down...")
	application.Shutdown()
}
