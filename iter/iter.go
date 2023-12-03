// Package iter is based on the current proposal for x/exp/xiter,
// a possible new package with iterator adapters.
// See [https://github.com/golang/go/issues/61898].
package iter

// Seq is an iterator over sequences of individual values.
// When called as seq(yield), seq calls yield(v) for each value v in the sequence,
// stopping early if yield returns false.
type Seq[V any] func(yield func(V) bool)

// Seq2 is an iterator over sequences of pairs of values, most commonly key-value pairs.
// When called as seq(yield), seq calls yield(k, v) for each pair (k, v) in the sequence,
// stopping early if yield returns false.
type Seq2[K, V any] func(yield func(K, V) bool)

// Concat returns an iterator over the concatenation of the sequences.
func Concat[V any](seqs ...Seq[V]) Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			seq(yield)
		}
	}
}

// Concat2 returns an iterator over the concatenation of the sequences.
func Concat2[K, V any](seqs ...Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			seq(yield)
		}
	}
}

// Filter returns an iterator over seq that only includes
// the values v for which f(v) is true.
func Filter[V any](f func(V) bool, seq Seq[V]) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if f(v) && !yield(v) {
				return
			}
		}
	}
}

// Filter2 returns an iterator over seq that only includes
// the pairs k, v for which f(k, v) is true.
func Filter2[K, V any](f func(K, V) bool, seq Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if f(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// Map returns an iterator over f applied to seq.
func Map[T, U any](f func(T) U, seq Seq[T]) Seq[U] {
	return func(yield func(U) bool) {
		for in := range seq {
			if !yield(f(in)) {
				return
			}
		}
	}
}

// Map2 returns an iterator over f applied to seq.
func Map2[K1, V1, K2, V2 any](f func(K1, V1) (K2, V2), seq Seq2[K1, V1]) Seq2[K2, V2] {
	return func(yield func(K2, V2) bool) {
		for k, v := range seq {
			if !yield(f(k, v)) {
				return
			}
		}
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
