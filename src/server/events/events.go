package events

type Events struct {
	Broadcast chan Message
}

func NewEvents() Events {
	return Events{
		make(chan Message),
	}
}
