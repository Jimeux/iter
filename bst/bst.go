package bst

import (
	"cmp"

	"github.com/Jimeux/iter/iter"
)

type node[T any] struct {
	val   T
	left  *node[T]
	right *node[T]
}

// BST is a basic binary search tree implementation.
type BST[T cmp.Ordered] struct {
	root *node[T]
}

func New[T cmp.Ordered]() *BST[T] {
	return &BST[T]{}
}

// Preorder returns an iterator that performs a pre-order traversal over the tree.
func (t *BST[T]) Preorder() iter.Seq[T] {
	return func(yield func(T) bool) bool {
		return t.preorder(t.root, yield)
	}
}

func (t *BST[T]) preorder(root *node[T], yield func(T) bool) bool {
	return root == nil ||
		yield(root.val) && t.preorder(root.left, yield) && t.preorder(root.right, yield)
}

// Postorder returns an iterator that performs a post-order traversal over the tree.
func (t *BST[T]) Postorder() iter.Seq[T] {
	return func(yield func(T) bool) bool {
		return t.postorder(t.root, yield)
	}
}

func (t *BST[T]) postorder(root *node[T], yield func(T) bool) bool {
	return root == nil ||
		t.postorder(root.right, yield) && t.postorder(root.left, yield) && yield(root.val)
}

// Inorder returns an iterator that performs an in-order traversal over the tree.
func (t *BST[T]) Inorder() iter.Seq[T] {
	return func(yield func(T) bool) bool {
		return t.inorder(t.root, yield)
	}
}

func (t *BST[T]) inorder(root *node[T], yield func(T) bool) bool {
	return root == nil ||
		t.inorder(root.left, yield) && yield(root.val) && t.inorder(root.right, yield)
}

// Add adds val to the tree.
func (t *BST[T]) Add(val T) {
	t.root = t.add(val, t.root)
}

func (t *BST[T]) add(val T, n *node[T]) *node[T] {
	if n == nil {
		return &node[T]{val: val}
	}

	ord := cmp.Compare(val, n.val)
	if ord < 0 {
		n.left = t.add(val, n.left)
	} else if ord > 0 {
		n.right = t.add(val, n.right)
	} // don't add duplicates
	return n
}

// Remove removes val from the tree if it exists.
func (t *BST[T]) Remove(val T) {
	t.root = t.remove(val, t.root)
}

func (t *BST[T]) remove(val T, n *node[T]) *node[T] {
	if n == nil {
		return nil
	}

	ord := cmp.Compare(val, n.val)
	if ord < 0 {
		n.left = t.remove(val, n.left)
	} else if ord > 0 {
		n.right = t.remove(val, n.right)
	} else {
		if n.right != nil {
			suc := n.right
			for suc.left != nil {
				suc = suc.left
			}
			n.val = suc.val
			n.right = t.remove(suc.val, n.right)
		} else if n.left != nil {
			pred := n.left
			for pred.right != nil {
				pred = pred.right
			}
			n.val = pred.val
			n.left = t.remove(pred.val, n.left)
		} else {
			return nil
		}
	}
	return n
}

// Contains returns true if t contains val.
func (t *BST[T]) Contains(val T) bool {
	return t.contains(val, t.root)
}

func (t *BST[T]) contains(val T, n *node[T]) bool {
	if n == nil {
		return false
	}

	ord := cmp.Compare(val, n.val)
	if ord == 0 {
		return true
	} else if ord < 0 {
		return t.contains(val, n.left)
	}
	return t.contains(val, n.right)
}
