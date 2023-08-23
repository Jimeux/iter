# Go iterators

### Some basic code examples based on the current proposals for iterators and the `iter` package.

## Setup

```
go install golang.org/dl/gotip@latest
gotip download 510541
```

## Run
```
gotip test ./...
```

## Proposal Summary

[proposal: spec: add range over int, range over func · Issue \#61405 · golang/go](https://github.com/golang/go/issues/61405#issuecomment-1638896606)

* Motivation
    * Allow custom generic containers, e.g. ordered maps, to work with `range` loops
    * Provide a better answer for the many functions in the standard library that collect a sequence of results and return the whole thing as a slice
        * If the results can be generated one at a time, then a representation that allows iterating over them scales better than returning an entire slice
        * Ex: `strings.Split`, `strings.Fields`
    * Could add iterators that wouldn’t be practical with slices
        * Ex: `strings.Lines(text)`
    * `bufio.Reader` and `bufio.Scanner` can be iterated over, but you have to know the pattern, which tends to be different for each type. 
      * Establishing a standard way to express iteration will help converge the many different approaches that exist today.
* Range over integers
    * `for i range N`
    * Incredibly common, so better as syntax sugar than as a library function
* What will idiomatic APIs with range functions look like?
    * Not decided
    * Depends on standard library proposal
    * Ex: `All() Seq[T]` for all elements in the container

## `iter` package

[proposal: x/exp/xiter: new package with iterator adapters](https://github.com/golang/go/issues/61898)

* See `iter` package in this repository
