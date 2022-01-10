package v2

import "sync"

type Bus struct {
	rw     sync.RWMutex
	topics map[string]*Topic
}

func (b *Bus) Sub(topic string, receiver Receiver) {
	b.rw.Lock()
	if t, ok := b.topics[topic]; ok {
		b.rw.Unlock()
		t.rw.Lock()
		defer t.rw.Unlock()
		t.receivers = append(t.receivers, receiver)
	} else {
		b.topics[topic] = &Topic{
			receivers: []Receiver{{}},
		}
		b.rw.Unlock()
	}
}
