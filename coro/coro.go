package coro

import "fmt"

func New[In, Out any](f func(In, func(Out) In) Out) (resume func(In) (Out, bool)) {
	cin := make(chan In)
	cout := make(chan Out)
	running := true
	resume = func(in In) (out Out, ok bool) {
		if !running {
			return
		}
		cin <- in
		return <-cout, running
	}
	yield := func(out Out) In {
		cout <- out
		return <-cin
	}
	go func() {
		out := f(<-cin, yield)
		running = false
		cout <- out
	}()
	return resume
}

func Pull[V any](push func(yield func(V) bool) bool) (pull func() (V, bool), stop func()) {
	copush := func(more bool, yield func(V) bool) V {
		if more {
			push(yield)
		}
		var zero V
		return zero
	}
	resume := New(copush)
	pull = func() (V, bool) {
		return resume(true)
	}
	stop = func() {
		resume(false)
	}
	return pull, stop
}

func NewDebug[In, Out any](f func(In, func(Out) In) Out) (resume func(In) (Out, bool)) {
	cin := make(chan In)
	cout := make(chan Out)
	running := true
	resume = func(in In) (out Out, ok bool) {
		if !running {
			return
		}
		fmt.Printf("resume: send in=%v\n", in)
		cin <- in
		out = <-cout
		fmt.Printf("resume: receive out=%v\n", out)
		return out, running
	}
	yield := func(out Out) In {
		fmt.Printf("yield: send out=%v\n", out)
		cout <- out
		in := <-cin
		fmt.Printf("yield: receive in=%v\n", in)
		return in
	}
	go func() {
		fmt.Println("go func: wait for cin")
		in := <-cin
		fmt.Printf("go func: received in=%v\n", in)
		out := f(in, yield)
		running = false
		cout <- out
	}()
	return resume
}
