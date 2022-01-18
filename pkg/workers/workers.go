package workers

import (
	"github.com/eugeneuskov/grpc-chat/config"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"time"
)

type Workers struct {
	rabbitConfig *config.Rabbit
	services     *services.Services
	CreateUser
}

func NewWorkers(rabbitConfig *config.Rabbit, services *services.Services) *Workers {
	return &Workers{rabbitConfig: rabbitConfig, services: services}
}

func (w *Workers) Run() {
	w.initWorkers()
	time.Sleep(time.Duration(w.rabbitConfig.DelayWorkersRun) * time.Second)

	go w.CreateUser.Start()
}

func (w *Workers) initWorkers() {
	w.CreateUser = newCreateUserWorker(
		w.rabbitConfig.AmqpServerUrl,
		w.rabbitConfig.Queues.CreateUser,
		w.services.ExternalAuth,
	)
}
