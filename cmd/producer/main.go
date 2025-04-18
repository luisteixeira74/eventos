package main

import "github.com/luisteixeira74/go-expert-eventos/pkg/rabbitmq"

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

	rabbitmq.Publish(ch, "Hello World", rabbitmq.ExchangeName, rabbitmq.RoutingKey)
}
