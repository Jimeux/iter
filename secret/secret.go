package secret

// Int is a secret int that may be hidden (not visible).
type Int struct {
	Val     int
	Visible bool
}

type Ints []Int

// All is an iterator over all elements.
func (ii Ints) All(yield func(Int) bool) {
	for _, s := range ii {
		if !yield(s) {
			return
		}
	}
}

// AllWithIndex is an iterator over all elements with the index in the form (index, value).
func (ii Ints) AllWithIndex(yield func(int, Int) bool) {
	for i, s := range ii {
		if !yield(i, s) {
			return
		}
	}
}
