package workers

import (
	"errors"
	"github.com/streadway/amqp"
)

func connectToRabbit(amqpServerUrl string) (*amqp.Connection, error) {
	connectionRabbitMq, err := amqp.Dial(amqpServerUrl)
	if err != nil {
		return nil, err
	}

	if connectionRabbitMq == nil {
		return nil, errors.New("amqp connection is nil")
	}

	return connectionRabbitMq, nil
}

func connectToChannel(connectionRabbitMq *amqp.Connection) (*amqp.Channel, error) {
	channelRabbitMq, err := connectionRabbitMq.Channel()
	if err != nil {
		return nil, err
	}

	if channelRabbitMq == nil {
		return nil, errors.New("amqp channel is nil")
	}

	return channelRabbitMq, nil
}

func channelQueueDeclare(channelRabbitMq *amqp.Channel, queueName string) error {
	_, err := channelRabbitMq.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)

	return err
}

func channelConsume(channelRabbitMq *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	return channelRabbitMq.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
}
