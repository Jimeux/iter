package secret

// Int is a secret int that may be hidden (not visible).
type Int struct {
	Val     int
	Visible bool
}

type Ints []Int

// All returns an iterator over all elements.
// TODO implement
// func (ss Ints) All() ...

// All returns an iterator over all elements with the index in the form (index, value).
// TODO implement
// func (ss Ints) AllWithIndex() ...
