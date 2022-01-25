package workers

import (
	"fmt"
	"github.com/eugeneuskov/grpc-chat/pkg/services"
	"github.com/eugeneuskov/grpc-chat/pkg/structs"
	"github.com/mailru/easyjson"
	"github.com/streadway/amqp"
	"log"
)

type createUserWorker struct {
	amqpServerUrl string
	queueName     string
	service       services.External
}

func (c *createUserWorker) Start() {
	connectionRabbitMQ, err := connectToRabbit(c.amqpServerUrl)
	if err != nil {
		log.Printf("connect to rabbit failed: %s\n", err.Error())
		return
	}
	defer func(connectionRabbitMQ *amqp.Connection) {
		_ = connectionRabbitMQ.Close()
	}(connectionRabbitMQ)

	channelRabbitMQ, err := connectToChannel(connectionRabbitMQ)
	if err != nil {
		log.Printf("connect to channel failed: %s\n", err.Error())
		return
	}
	defer func(channelRabbitMQ *amqp.Channel) {
		_ = channelRabbitMQ.Close()
	}(channelRabbitMQ)

	if err = channelQueueDeclare(channelRabbitMQ, c.queueName); err != nil {
		log.Printf("channel queue declare failed: %s\n", err.Error())
		return
	}

	messages, err := channelConsume(channelRabbitMQ, c.queueName)
	if err != nil {
		log.Printf("channel consume failed: %s\n", err.Error())
		return
	}
	println("create user worker started...")

	forever := make(chan bool)
	go func() {
		for message := range messages {
			user, err := c.validate(message.Body)
			if err != nil {
				println(err.Error())
				continue
			}

			if err := c.service.CreateUser(user); err != nil {
				log.Printf("create user error: %s\n", err.Error())
				continue
			}
		}
	}()
	<-forever
}

func (c *createUserWorker) validate(message []byte) (*structs.User, error) {
	var user structs.User
	if err := easyjson.Unmarshal(message, &user); err != nil {
		return nil, fmt.Errorf("parse error: %s", err.Error())
	}

	if user.IsEmpty() {
		return nil, fmt.Errorf("not fulled data")
	}

	return &user, nil
}

func newCreateUserWorker(amqpServerUrl, queueName string, service services.External) *createUserWorker {
	return &createUserWorker{amqpServerUrl, queueName, service}
}
