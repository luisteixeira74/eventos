package user

import (
	"fmt"
	"sync"

	"github.com/luisteixeira74/go-expert-eventos/pkg/events"
)

type CreateUserHandler struct{}

func (h CreateUserHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Create user on DB with payload:", event.GetPayLoad())
}
