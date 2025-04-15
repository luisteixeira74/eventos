package main

import (
	"github.com/luisteixeira74/go-expert-eventos/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		// Process the message
		// For example, print the message body
		println(string(msg.Body))

		// Acknowledge the message
		msg.Ack(false)
	}

}
