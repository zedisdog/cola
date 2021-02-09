package task

type node struct {
	next    *node
	content interface{}
}

type link struct {
	count int
	first *node
	end   *node
}

func newLink() *link {
	return &link{}
}

func (l *link) pushBack(content interface{}) {
	l.count = l.count + 1
	if l.count == 1 {
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
}

func (l *link) pop() interface{} {
	if l.count-1 < 0 {
		return nil
	}
	l.count = l.count - 1
	first := l.first
	if l.count < 1 {
		l.first = nil
	} else {
		l.first = l.first.next
	}
	return first.content
}
