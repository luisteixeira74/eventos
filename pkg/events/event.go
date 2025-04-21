package events

import "time"

const (
	EventNameEventoTeste = "EventoTeste"
	EventNameSendEmail   = "SendEmail"
	EventNameCreateLog   = "CreateLog"
)

type Event struct {
	Name     string
	DateTime time.Time
	PayLoad  interface{}
}

func (e Event) GetName() string {
	return e.Name
}

func (e Event) GetDateTime() time.Time {
	return e.DateTime
}

func (e Event) GetPayLoad() interface{} {
	return e.PayLoad
}

func NewEvent(name string, payload interface{}) EventInterface {
	return Event{
		Name:     name,
		DateTime: time.Now(),
		PayLoad:  payload,
	}
}
