package evbus

import (
	"sync"
)

var EvBus *Bus

func init() {
	EvBus = &Bus{}
}

type Handler func(Event)

type Bus struct {
	m sync.Map
}

func Sub(eventName interface{}, handler Handler) {
	EvBus.Sub(eventName, handler)
}
func (b *Bus) Sub(eventName interface{}, handler Handler) {
	var handlers []Handler
	if m, ok := b.m.Load(eventName); !ok {
		handlers = make([]Handler, 0, 1)
	} else {
		handlers = m.([]Handler)
	}

	handlers = append(handlers, handler)

	b.m.Store(eventName, handlers)
}

func Fire(event Event) {
	EvBus.Fire(event)
}
func (b *Bus) Fire(event Event) {
	if handlers, ok := b.m.Load(event.Name); !ok {
		return
	} else {
		for _, handler := range handlers.([]Handler) {
			handler(event)
		}
	}
}
