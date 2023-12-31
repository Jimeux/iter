package list

import (
	"fmt"
	"strings"
)

type node[T any] struct {
	val  T
	next *node[T]
}

// List is a basic singly-linked list implementation.
type List[T comparable] struct {
	head *node[T]
	size int
}

func New[T comparable]() *List[T] {
	return &List[T]{}
}

// All is an iterator over all elements starting from the head of l.
func (l *List[T]) All(yield func(T) bool) {
	for h := l.head; h != nil; h = h.next {
		if !yield(h.val) {
			return
		}
	}
}

// Backward is a reverse iterator starting from the tail of l.
func (l *List[T]) Backward(yield func(T) bool) {
	l.backward(l.head, yield)
}

func (l *List[T]) backward(n *node[T], yield func(T) bool) bool {
	return n == nil || l.backward(n.next, yield) && yield(n.val)
}

func (l *List[T]) Empty() bool {
	return l.head == nil
}

func (l *List[T]) Len() int {
	return l.size
}

// Push adds the given value to the head of the list.
func (l *List[T]) Push(val T) {
	l.head = &node[T]{val: val, next: l.head}
	l.size++
}

// Pop removes and returns the head of the list.
func (l *List[T]) Pop() (T, bool) {
	if l.Empty() {
		var t T
		return t, false
	}
	h := l.head
	l.head = l.head.next
	l.size--
	return h.val, true
}

// Remove returns the first item matching the given val based on == comparison.
func (l *List[T]) Remove(val T) bool {
	if l.Empty() {
		return false
	}
	dummy := &node[T]{next: l.head}
	head, prev := l.head, dummy
	for head != nil {
		if head.val == val {
			prev.next = head.next
			break
		}
		prev = head
		head = head.next
	}
	l.head = dummy.next
	l.size--
	return true
}

// String implements Stringer for the list.
func (l *List[T]) String() string {
	sb := strings.Builder{}
	h := l.head
	for h != nil {
		sb.WriteString(fmt.Sprint(h.val))
		if h.next != nil {
			sb.WriteString(",")
		}
		h = h.next
	}
	return "[" + sb.String() + "]"
}
