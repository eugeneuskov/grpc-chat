package app

import (
	"database/sql"
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/repositories"
	"github.com/eugeneuskov/grpc-chat/pkg/server"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"gorm.io/gorm"
	"log"
	"time"
)

type Application struct {
	config *config.Config
	server *server.Server
	db     *gorm.DB
	sqlDb  *sql.DB
}

func NewApplication(config *config.Config) *Application {
	return &Application{
		config: config,
	}
}

func (a *Application) Run() {
	println("App starting...")
	a.dbConnect()
	a.server = server.NewServer(&a.config.Tls, &a.config.App)
	a.server.InitGrpcServer()

	serviceList := services.NewServices(repositories.NewRepositories(a.db))
	services.NewGrpcServices(a.server.Grpc).InitServices(serviceList)

	go a.server.Run()
}

func (a *Application) dbConnect() {
	db, err := repositories.NewPostgresConnection(&a.config.Database)
	if err != nil {
		log.Fatalf("Failed to initialized DB: %s\n", err.Error())
		return
	}

	a.sqlDb, err = db.DB()
	if err != nil {
		log.Fatalf("Failed to start DB: %s\n", err.Error())
		return
	}
	a.sqlDb.SetMaxIdleConns(10)
	a.sqlDb.SetMaxOpenConns(100)
	a.sqlDb.SetConnMaxLifetime(time.Hour)

	a.db = db
	println("DB connected\nRun migrations...")
	repositories.Migrate(a.db)
	println("Migrations completed")
}

func (a *Application) Shutdown() {
	println("\nApp shutting down...")

	if err := a.sqlDb.Close(); err != nil {
		println("Error while closing DB: %s\n", err.Error())
	} else {
		println("DB closed")
	}
	// _ = a.server.Shutdown(context.Background())

	println("OFF")
}
