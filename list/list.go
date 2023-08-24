package list

import (
	"fmt"
	"strings"

	"github.com/Jimeux/iter/iter"
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

// All returns an iterator over all elements starting from l.head.
func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) bool {
		for h := l.head; h != nil; h = h.next {
			if !yield(h.val) {
				return false
			}
		}
		return true
	}
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
