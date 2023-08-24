package secret

import "github.com/Jimeux/iter/iter"

// Int is a secret int that may be hidden (not visible).
type Int struct {
	Val     int
	Visible bool
}

type Ints []Int

// All returns an iterator over all elements.
func (ss Ints) All() iter.Seq[Int] {
	return func(yield func(Int) bool) bool {
		for _, s := range ss {
			if !yield(s) {
				return false
			}
		}
		return true
	}
}

// AllWithIndex returns an iterator over all elements with the index in the form (index, value).
func (ss Ints) AllWithIndex() iter.Seq2[int, Int] {
	return func(yield func(int, Int) bool) bool {
		for i, s := range ss {
			if !yield(i, s) {
				return false
			}
		}
		return true
	}
}
