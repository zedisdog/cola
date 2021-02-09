package task

import "testing"

func TestLink_PushBack(t *testing.T) {
	l := newLink()
	l.pushBack(1)
	if l.count != 1 {
		t.Fatal("count should be 1")
	}
	if l.first.content != 1 {
		t.Fatal("first node's content should be 1")
	}
	if l.end.content != 1 {
		t.Fatal("first node's content should be 1")
	}

	l.pushBack(2)
	if l.count != 2 {
		t.Fatal("count should be 2")
	}
	if l.first.content != 1 {
		t.Fatal("first node's content should be 1")
	}
	if l.end.content != 2 {
		t.Fatal("first node's content should be 2")
	}
}

func TestLink_pop(t *testing.T) {
	l := newLink()
	l.pushBack(1)
	l.pushBack(2)

	content := l.pop()
	if content != 1 {
		t.Fatal("node's content should be 1")
	}
	content = l.pop()
	if content != 2 {
		t.Fatal("node's content should be 2")
	}
	content = l.pop()
	if content != nil {
		t.Fatal("node's content should be nil")
	}
}
