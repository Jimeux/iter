package maps

import "github.com/Jimeux/iter/iter"

// Keys returns an iterator over the keys in m.
func Keys[M ~map[K]V, K comparable, V any](m M) iter.Seq[K] {
	return func(yield func(K) bool) bool {
		for k := range m {
			if !yield(k) {
				return false
			}
		}
		return true
	}
}

// Values returns an iterator over the values in m.
func Values[M ~map[K]V, K comparable, V any](m M) iter.Seq[V] {
	return func(yield func(V) bool) bool {
		for _, v := range m {
			if !yield(v) {
				return false
			}
		}
		return true
	}
}
