package list

import (
	"slices"
	"strconv"
	"testing"

	"github.com/Jimeux/iter/iter"
)

func TestListIterators(t *testing.T) {
	l := New[int]()
	for i := range 5 {
		l.Push(i)
	}

	t.Run("All", func(t *testing.T) {
		got := make([]int, 0, l.Len())
		for s := range l.All() {
			got = append(got, s)
		}

		want := []int{4, 3, 2, 1, 0}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
	t.Run("FilterMapReduce", func(t *testing.T) {
		filter := iter.Filter(func(i int) bool {
			return i > 1
		}, l.All())
		mapped := iter.Map(func(i int) string {
			return strconv.Itoa(i) + "!"
		}, filter)
		got := iter.Reduce(nil, func(sum []string, s string) []string {
			return append(sum, s)
		}, mapped)

		want := []string{"4!", "3!", "2!"}
		if !slices.Equal(got, want) {
			t.Fatalf("got %v but want %v", got, want)
		}
	})
}

func TestList(t *testing.T) {
	l := New[int]()
	if !l.Empty() {
		t.Fatal("new list was not empty")
	}

	popped, ok := l.Pop()
	if popped != 0 || ok {
		t.Fatalf("pop on an empty list returned %d, %t", popped, ok)
	}

	l.Push(1)
	popped, ok = l.Pop()
	if popped != 1 || !ok {
		t.Fatalf("expected 1 for Pop but got %d, %t", popped, ok)
	}
	if !l.Empty() {
		t.Fatal("Empty returned false for empty list")
	}

	l.Push(3)
	l.Push(2)
	l.Push(1)

	serial := l.String()
	if serial != "[1,2,3]" {
		t.Fatalf("expected [1,2,3] from String, but got %s", serial)
	}

	ok = l.Remove(1)
	if !ok {
		t.Fatal("expected Remove to return true")
	}
	serial = l.String()
	if serial != "[2,3]" {
		t.Fatalf("expected [2,3] from String, but got %s", serial)
	}
}
