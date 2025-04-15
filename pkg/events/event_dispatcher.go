package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered for this event")

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (d *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// Check if the handler is already registered for the event
	for _, h := range d.handlers[eventName] {
		if h == handler {
			return ErrHandlerAlreadyRegistered
		}
	}

	// Register the handler for the event
	d.handlers[eventName] = append(d.handlers[eventName], handler)

	return nil
}

func (d *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// Check if the handler is registered for the event
	for _, h := range d.handlers[eventName] {
		if h == handler {
			return true
		}
	}
	return false
}

func (d *EventDispatcher) Dispatch(event EventInterface) error {
	// Check if there are handlers registered for the event
	if handlers, ok := d.handlers[event.GetName()]; ok {
		wg := sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, &wg)
		}
		wg.Wait()
	} else {
		return errors.New("no handlers registered for this event")
	}
	return nil
}

func (d *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	// Check if the handler is registered for the event
	for i, h := range d.handlers[eventName] {
		if h == handler {
			// ex: [1, 2, 3] -> [1, 3]
			// remove the handler from the slice
			d.handlers[eventName] = append(d.handlers[eventName][:i], d.handlers[eventName][i+1:]...)
			return nil
		}
	}
	return errors.New("handler not registered for this event")
}

func (d *EventDispatcher) Clear() {
	d.handlers = make(map[string][]EventHandlerInterface)
}
