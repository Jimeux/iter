# coro

### Sample code and notes about https://research.swtch.com/coro 

* Coroutines provide __concurrency without parallelism__: when one coroutine is running, the one that resumed it or yielded to it is not.
* Because __coroutines run one at a time and only switch at specific points__ in the program, the coroutines can share data among themselves without races. 
  * The explicit switches serve as synchronization points, creating happens-before edges

## Converted Iterator Sequence

* The exact order of the initial events depends on goroutine scheduling
  * e.g. `resume` may be called before `goroutine1` blocks on `<-cin`
* `resume`
  * executes on `main`
  * returns control from `main` to `goroutine1`
* `f` function argument to `coro.New` (wrapped `iterator`)
  * executes on `goroutine1`
  * calls `yield`
    * `yield` returns control from `goroutine1` to `main`

```mermaid
sequenceDiagram
    actor m as main
    actor go as goroutine1
    participant in as cin
    participant out as cout

    m->>m: resume := coro.New(push)
    go->>in: in := <-cin
    note over m,in: coro.New starts goroutine1<br>which blocks on <-cin, waiting<br>for the initial value from resume
    
    m->>m: resume(true)
    m->>in: resume sends: cin <- true
    m->>out: resume blocks: out <- cout
    
    in-->>go: true is received
    go->>go: f(true, yield) is called
    note over m,in: f is the arg passed to coro.New (push)<br>yield is created inside coro.New, and connects cin and cout
    note over m,in: This begins the iterator inside push(),<br>which runs in goroutine1
    
    go->>go: iterator calls yield(1) with first element
    go->>out: yield sends first element (1) to cout

    go->>in: goroutine1 blocks on <-cin
    out-->>m: first element (1) is received by resume
    m->>m: v, ok := resume(true)<br>has returned
    note over m,go: Control has returned to main<br>from goroutine1

    alt
        m->>m: resume(true)
        note over m,go: Calling resume(true) again<br>returns control to goroutine1<br>and continues the iterator
    else
        m->>m: resume(false)
        note over m,go: Calling resume(false)<br>terminates the coroutine
        note over m,go: TODO sequence
    end
```

---

## Kotlin

>  a coroutine is not bound to any particular thread. It may suspend its execution in one thread and resume in another one.
