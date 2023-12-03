# Go iterators

### Basic code examples based on the current proposals for iterators and the `iter` package.

The following packages are included, with tests demonstrating the iterator implementations.

* [bst](bst) - Binary search tree with traversal iterators
* [iter](iter) - Core types and functions from the proposals
* [list](list) - Linked list with traversal iterators
* [maps](maps) - Iterators for keys and values of maps
* [secret](secret) - Iterators for a custom slice type

## Setup

Follow instructions for installing [gotip](https://pkg.go.dev/golang.org/dl/gotip). (Should not be necessary with Go 1.22.)

```bash
$ go install golang.org/dl/gotip@latest
$ gotip download
```

## Run

Run tests with `gotip` using the `rangefunc` experiment flag.

```bash
$ GOEXPERIMENT=rangefunc gotip test ./...
```
