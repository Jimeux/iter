// Package iter is based on the current proposal for x/exp/xiter,
// a possible new package with iterator adapters.
// See [https://github.com/golang/go/issues/61898].
package iter

// Seq is an iterator over a sequence of values.
// It can be used with for-range loops like: for v := range seq().
type Seq[T any] func(func(T) bool) bool

// Seq2 is an iterator over a sequence of pairs of values.
// It can be used with for-range loops like: for k, v := range seq().
type Seq2[T, U any] func(func(T, U) bool) bool

// Concat returns an iterator over the concatenation of the sequences.
func Concat[V any](seqs ...Seq[V]) Seq[V] {
	return func(yield func(V) bool) bool {
		for _, seq := range seqs {
			if !seq(yield) {
				return false
			}
		}
		return true
	}
}

// Concat2 returns an iterator over the concatenation of the sequences.
func Concat2[K, V any](seqs ...Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) bool {
		for _, seq := range seqs {
			if !seq(yield) {
				return false
			}
		}
		return true
	}
}

// Filter returns an iterator over seq that only includes
// the values v for which f(v) is true.
func Filter[V any](f func(V) bool, seq Seq[V]) Seq[V] {
	return func(yield func(V) bool) bool {
		for v := range seq {
			if f(v) && !yield(v) {
				return false
			}
		}
		return true
	}
}

// Filter2 returns an iterator over seq that only includes
// the pairs k, v for which f(k, v) is true.
func Filter2[K, V any](f func(K, V) bool, seq Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) bool {
		for k, v := range seq {
			if f(k, v) && !yield(k, v) {
				return false
			}
		}
		return true
	}
}

// Map returns an iterator over f applied to seq.
func Map[T, U any](f func(T) U, seq Seq[T]) Seq[U] {
	return func(yield func(U) bool) bool {
		for in := range seq {
			if !yield(f(in)) {
				return false
			}
		}
		return true
	}
}

// Map2 returns an iterator over f applied to seq.
func Map2[K1, V1, K2, V2 any](f func(K1, V1) (K2, V2), seq Seq2[K1, V1]) Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) bool {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return false
			}
		}
		return true
	}
}

// Reduce combines the values in seq using f.
// For each value v in seq, it updates sum = f(sum, v)
// and then returns the final sum.
// For example, if iterating over seq yields v1, v2, v3,
// Reduce returns f(f(f(sum, v1), v2), v3).
func Reduce[Sum, V any](sum Sum, f func(Sum, V) Sum, seq Seq[V]) Sum {
	for v := range seq {
		sum = f(sum, v)
	}
	return sum
}

// Reduce2 combines the values in seq using f.
// For each pair k, v in seq, it updates sum = f(sum, k, v)
// and then returns the final sum.
// For example, if iterating over seq yields (k1, v1), (k2, v2), (k3, v3)
// Reduce returns f(f(f(sum, k1, v1), k2, v2), k3, v3).
func Reduce2[Sum, K, V any](sum Sum, f func(Sum, K, V) Sum, seq Seq2[K, V]) Sum {
	for k, v := range seq {
		sum = f(sum, k, v)
	}
	return sum
}
