package main

import (
	"github.com/luisteixeira74/go-expert-eventos/internal/bootstrap"
	"github.com/luisteixeira74/go-expert-eventos/pkg/events"
	"github.com/luisteixeira74/go-expert-eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Setup
	if err := rabbitmq.SetupRabbitMQ(ch); err != nil {
		panic(err)
	}

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, rabbitmq.QueueName)

	dispatcher := bootstrap.RegisterEventHandlers()

	for msg := range msgs {
		// Transformar o body em um Event real
		event := events.NewEvent("UserCreated", string(msg.Body))

		dispatcher.Dispatch(event) // handler executa a função com o payload

		// Acknowledge the message
		msg.Ack(false)
	}
}
