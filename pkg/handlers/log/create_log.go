package log

import (
	"fmt"
	"sync"

	"github.com/luisteixeira74/go-expert-eventos/pkg/events"
)

type CreateLogHandler struct{}

func (h CreateLogHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Create a Log with payload:", event.GetPayLoad())
}
