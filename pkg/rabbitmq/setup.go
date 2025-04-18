package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	ExchangeName = "notificacao.exchange"
	ExchangeType = "direct"
	QueueName    = "notificacao.queue"
	RoutingKey   = "enviar_notificacao"
)

// SetupRabbitMQ cria a exchange, fila e binding necessários.
func SetupRabbitMQ(ch *amqp.Channel) error {
	// Criar Exchange
	if err := ch.ExchangeDeclare(
		ExchangeName,
		ExchangeType,
		true,  // durable
		false, // auto-delete
		false, // internal
		false, // no-wait
		nil,   // args
	); err != nil {
		return err
	}

	// Criar Fila
	_, err := ch.QueueDeclare(
		QueueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	// Fazer Binding
	if err := ch.QueueBind(
		QueueName,
		RoutingKey,
		ExchangeName,
		false, // no-wait
		nil,
	); err != nil {
		return err
	}

	log.Println("Exchange, fila e binding configurados com sucesso!")
	return nil
}

func GetRoutingKey() string {
	return RoutingKey
}

func GetExchangeName() string {
	return ExchangeName
}
