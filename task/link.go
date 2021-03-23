package task

import (
	"sync"
)

type node struct {
	next    *node
	content interface{}
}

type link struct {
	count int
	first *node
	end   *node
	lock  sync.Mutex
}

func newLink() *link {
	return &link{}
}

func (l *link) put(content interface{}) {
	l.lock.Lock()
	if l.count == 0 {
		l.end = &node{
			next:    nil,
			content: content,
		}
		l.first = l.end
	} else {
		l.end.next = &node{
			next:    nil,
			content: content,
		}
		l.end = l.end.next
	}
	l.count++
	l.lock.Unlock()
}

func (l *link) pop() interface{} {
	if l.count == 0 {
		return nil
	}
	l.lock.Lock()
	first := l.first
	l.first = l.first.next
	l.count--
	l.lock.Unlock()
	if first == nil {
		return nil
	}
	return first.content
}
