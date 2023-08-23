package bst

type node[T any] struct {
	val   T
	left  *node[T]
	right *node[T]
}

// BST is a basic binary search tree implementation.
type BST[T any] struct {
	root *node[T]
	// cmp compares the vals of two items in the tree.
	cmp func(T, T) int
}

func New[T any](cmp func(T, T) int) *BST[T] {
	return &BST[T]{cmp: cmp}
}

// Preorder returns an iterator that performs a pre-order traversal over the tree.
// TODO implement
// func (t *BST[T]) Preorder() ...

// Postorder returns an iterator that performs a pre-order traversal over the tree.
// TODO implement
// func (t *BST[T]) Postorder() ...

// Inorder returns an iterator that performs a pre-order traversal over the tree.
// TODO implement
// func (t *BST[T]) Inorder() ...

// Add adds val to the tree.
func (t *BST[T]) Add(val T) {
	t.root = t.add(val, t.root)
}

func (t *BST[T]) add(val T, n *node[T]) *node[T] {
	if n == nil {
		return &node[T]{val: val}
	}

	cmp := t.cmp(val, n.val)
	if cmp < 0 {
		n.left = t.add(val, n.left)
	} else if cmp > 0 {
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

	cmp := t.cmp(val, n.val)
	if cmp < 0 {
		n.left = t.remove(val, n.left)
	} else if cmp > 0 {
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

	cmp := t.cmp(val, n.val)
	if cmp == 0 {
		return true
	} else if cmp < 0 {
		return t.contains(val, n.left)
	}
	return t.contains(val, n.right)
}
