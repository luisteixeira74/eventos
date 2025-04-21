package main

import (
	"encoding/json"
	"log"

	"github.com/luisteixeira74/go-expert-eventos/pkg/rabbitmq"
)

type UserCreatePayload struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

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

	// 1. Criar payload
	payload := UserCreatePayload{
		Name:      "João da Silva",
		Email:     "joao@example.com",
		CreatedAt: "2025-04-21T15:04:05Z",
	}

	// 2. Transformar para JSON string
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Erro ao serializar JSON: %v", err)
	}

	body := string(jsonBytes)

	err = rabbitmq.Publish(ch, body, rabbitmq.ExchangeName, rabbitmq.RoutingKey)
	if err != nil {
		log.Fatalf("Erro ao publicar evento: %v", err)
	}

	log.Println("Evento publicado com sucesso!")
}
