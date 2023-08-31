package coro

import (
	"fmt"
	"testing"

	"github.com/Jimeux/iter/list"
)

func TestCoro(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		// since resume can only run when the calling goroutine is blocked,
		// and vice versa, sharing the running variable is not a race.
		// The two are synchronizing by taking turns executing.
		// If resume is called after the coroutine has exited,
		// resume returns a zero value and false.
		resume := New(func(_ int, yield func(string) int) string {
			// resume can be called once for each yield, plus once for the final return value
			yield("hello")
			yield("world")
			return "done"
		})
		for i := 0; i < 4; i++ {
			s, ok := resume(0)
			fmt.Printf("%q %v\n", s, ok)
		}
	})

	t.Run("iterator with New", func(t *testing.T) {
		l := list.New[int]()
		for i := range 2 {
			l.Push(i + 1)
		}

		push := func(more bool, yield func(int) bool) int {
			fmt.Println("call push (aka f)")
			if more {
				l.All()(yield)
			}
			var zero int
			return zero
		}
		resume := New(push)
		for {
			fmt.Println("call resume")
			v, ok := resume(true)
			fmt.Printf("ok: %t\n", ok)
			if !ok {
				fmt.Printf("ok: %t\n", ok)
				break
			}
			fmt.Printf("v=%d | ok=%t\n", v, ok)
		}
	})

	t.Run("Pull", func(t *testing.T) {
		l1 := list.New[int]()
		for i := range 2 {
			l1.Push(i)
		}
		l2 := list.New[int]()
		for i := range 5 {
			l2.Push(i * 10)
		}

		next1, stop1 := Pull(l1.All())
		next2, stop2 := Pull(l2.All())
		defer stop1()
		defer stop2()
		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 || !ok2 {
				fmt.Printf("ok1 || ok2: %t, %t\n", ok1, ok2)
				break
			}
			fmt.Printf("v1=%d, v2=%d | ok1=%t, ok2=%t\n", v1, v2, ok1, ok2)
		}
	})
}
