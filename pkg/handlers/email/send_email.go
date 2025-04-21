package email

import (
	"fmt"
	"sync"

	"github.com/luisteixeira74/go-expert-eventos/pkg/events"
)

type SendEmailHandler struct{}

func (h SendEmailHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Send email to user with payload:", event.GetPayLoad())
}
