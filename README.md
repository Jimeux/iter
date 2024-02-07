# Go iterators

### Basic code examples based on the current proposals for iterators and the `iter` package.

The following packages are included, with tests demonstrating the iterator implementations:

* [bst](bst) - Binary search tree with traversal iterators
* [xiter](xiter) - Functions from the [iterator library proposal](https://github.com/golang/go/issues/61898)
* [list](list) - Linked list with traversal iterators
* [maps](maps) - Iterators for keys and values of maps
* [secret](secret) - Iterators for a custom slice type

See the related blog post on Medium:

* [A look at iterators in Go](https://medium.com/eureka-engineering/a-look-at-iterators-in-go-f8e86062937c)

## Run

Ensure Go 1.22 is installed.

Run tests using the `rangefunc` experiment flag.

```bash
$ GOEXPERIMENT=rangefunc go test ./...
```
