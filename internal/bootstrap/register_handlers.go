package bootstrap

import (
	"github.com/luisteixeira74/go-expert-eventos/pkg/events"
	"github.com/luisteixeira74/go-expert-eventos/pkg/handlers/email"
	"github.com/luisteixeira74/go-expert-eventos/pkg/handlers/log"
	"github.com/luisteixeira74/go-expert-eventos/pkg/handlers/user"
)

func RegisterEventHandlers() *events.EventDispatcher {
	dispatcher := events.NewEventDispatcher()

	// Este é um exemplo de simulação didática de múltiplos handlers para o mesmo evento ("UserCreate").
	// Em sistemas reais baseados em microservices, cada handler geralmente estaria em um serviço separado,
	// cada um escutando a fila e tratando o evento de forma independente.

	dispatcher.Register("UserCreate", user.CreateUserHandler{})
	dispatcher.Register("UserCreate", email.SendEmailHandler{})
	dispatcher.Register("UserCreate", log.CreateLogHandler{})

	return dispatcher
}
