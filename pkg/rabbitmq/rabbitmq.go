package rabbitmq

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

// opens a channel to the RabbitMQ server
func OpenChannel() (*ampq.Channel, error) {
	conn, err := ampq.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *ampq.Channel, out chan<- ampq.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	for msgs := range msgs {
		out <- msgs
	}

	return nil
}

func Publish(ch *ampq.Channel, body string, exName string, routingKey string) error {
	err := ch.Publish(
		exName,
		routingKey,
		false,
		false,
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
