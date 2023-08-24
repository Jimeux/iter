package bst

import (
	"slices"
	"testing"

	"github.com/Jimeux/iter/iter"
)

func TestBSTIterators(t *testing.T) {
	bst := New[int]()
	bst.Add(3) //         3
	bst.Add(5) //     2       5
	bst.Add(4) //  1        4
	bst.Add(2)
	bst.Add(1)

	tests := []struct {
		name string
		it   func() iter.Seq[int]
		want []int
	}{
		// TODO uncomment when Preorder is implemented
		// {"Preorder", bst.Preorder, []int{3, 2, 1, 5, 4}},
		// TODO uncomment when Postorder is implemented
		// {"Postorder", bst.Postorder, []int{4, 5, 1, 2, 3}},
		// TODO uncomment when Inorder is implemented
		// {"Inorder", bst.Inorder, []int{1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var got []int
			for s := range test.it() {
				got = append(got, s)
			}
			if !slices.Equal(got, test.want) {
				t.Fatalf("got %v but want %v", got, test.want)
			}
		})
	}
}

func TestBST(t *testing.T) {
	bst := New[int]()
	vals := []int{10, 5, 8, 3, 4, 6, 11, 1}
	for _, val := range vals {
		bst.Add(val)
	}
	for _, val := range vals {
		if !bst.Contains(val) {
			t.Fatalf("did not contain %d", val)
		}
	}
	for _, val := range vals {
		bst.Remove(val)
		if bst.Contains(val) {
			t.Fatalf("contained %d after call to Remove", val)
		}
	}
}
