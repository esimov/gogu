// Package btree provides an implementation of the B-tree data structure,
// which is a self-balancing tree data structure maintaining its values
// in sorted order and allowing each node to have more than two children,
// compared to the standard BST where each node has only two leaves.
// The implementation is a simplified version of
// https://algs4.cs.princeton.edu/62btree/BTree.java.
package btree

import (
	"github.com/esimov/gogu"
	"golang.org/x/exp/constraints"
)

// Max children per binary tree must be even or greater than 2.
const maxChildren = 4

// entry is the inner component of a node, which holds the node value and a pointer to the next node.
type entry[K, V any] struct {
	key       K
	value     V
	isRemoved bool
	next      *node[K, V]
}

// node is a data structure which defines how many children (leaves) each node has.
type node[K, V any] struct {
	m        int // number of children
	children [maxChildren]entry[K, V]
}

// newNode instantiates a new node with no leaves.
func newNode[K, V any](m int) *node[K, V] {
	return &node[K, V]{
		m: m,
	}
}

// BTree is the main component of the B-tree which starts only with one node, which is the root.
type BTree[K constraints.Ordered, V any] struct {
	n      int // the size of the tree (the number of nodes)
	heigth int // the height of the tree
	root   *node[K, V]
}

// New creates a new B-tree.
func New[K constraints.Ordered, V any]() *BTree[K, V] {
	return &BTree[K, V]{
		root: newNode[K, V](0),
	}
}

// Size returns the B-tree size (the number of elements).
func (t *BTree[K, V]) Size() int {
	return t.n
}

// IsEmpty checks if a B-tree is empty or not.
func (t *BTree[K, V]) IsEmpty() bool {
	return t.Size() == 0
}

// Height returns the B-tree size (how many levels it has).
func (t *BTree[K, V]) Height() int {
	return t.heigth
}

// Get searches for a value in the tree and if it's found it returns the value
// together with a boolean value signalig if it's found or not.
func (t *BTree[K, V]) Get(key K) (V, bool) {
	return t.search(t.root, key, t.heigth)
}

// search is a private method which is invoked by the Get method.
func (t *BTree[K, V]) search(n *node[K, V], key K, height int) (V, bool) {
	// external node
	if height == 0 {
		for i := 0; i < n.m; i++ {
			if gogu.Equal(key, n.children[i].key) {
				return n.children[i].value, true
			}
		}
	} else {
		// internal node
		for i := 0; i < n.m; i++ {
			if i+1 == n.m || gogu.Less(key, n.children[i+1].key) {
				return t.search(n.children[i].next, key, height-1)
			}
		}
	}

	var v V
	return v, false
}

// Put inserts a new value into the B-tree.
// If val is nil, this effectively deletes the value from the tree.
func (t *BTree[K, V]) Put(key K, val V) {
	u := t.insert(t.root, key, val, t.heigth, false)
	t.n++
	if u == nil {
		return
	}
	// split the root
	n := newNode[K, V](2)
	n.children[0] = entry[K, V]{
		key:  t.root.children[0].key,
		next: t.root,
	}
	n.children[1] = entry[K, V]{
		key:  u.children[0].key,
		next: u,
	}

	t.root = n
	t.heigth++
}

// insert is a private method which is invoked by the Put method.
func (t *BTree[K, V]) insert(n *node[K, V], key K, val V, height int, isRemoved bool) *node[K, V] {
	entry := entry[K, V]{
		key:   key,
		value: val,
		next:  nil,
	}

	var j int
	// external node
	if height == 0 {
		for j = 0; j < n.m; j++ {
			// If the value already exists in the B-tree this will be overwritten.
			if gogu.Equal(key, n.children[j].key) {
				n.children[j].value = val
				n.children[j].isRemoved = isRemoved
				return nil
			} else if gogu.Less(key, n.children[j].key) {
				break
			}
		}
	} else {
		// internal node
		for j = 0; j < n.m; j++ {
			if j+1 == n.m || gogu.Less(key, n.children[j+1].key) {
				node := t.insert(n.children[j].next, key, val, height-1, isRemoved)
				if node == nil {
					return nil
				}
				j++
				entry.key = node.children[0].key
				entry.next = node
				break
			}
		}
	}
	for i := n.m; i > j; i-- {
		n.children[i] = n.children[i-1]
	}

	n.children[j] = entry
	n.m++
	if n.m < maxChildren {
		return nil
	} else {
		return t.split(n)
	}
}

func (t *BTree[K, V]) split(n *node[K, V]) *node[K, V] {
	h := newNode[K, V](maxChildren / 2)
	n.m = maxChildren / 2

	for i := 0; i < n.m; i++ {
		h.children[i] = n.children[n.m+i]
	}
	return h
}

// Remove deletes an element from the B-tree.
func (t *BTree[K, V]) Remove(key K) {
	val, ok := t.Get(key)
	if !ok {
		return
	}
	t.n--
	t.insert(t.root, key, val, t.heigth, true)
}

// Traverse iterates over the values of the tree and invokes
// the callback function provided as argument over the node elements.
func (t *BTree[K, V]) Traverse(fn func(key K, val V)) {
	t.traverse(t.root, t.heigth, fn)
}

func (t *BTree[K, V]) traverse(n *node[K, V], depth int, fn func(K, V)) {
	// extenernal node
	if depth == 0 {
		for i := 0; i < n.m; i++ {
			l := n.children[i]
			if l.isRemoved {
				continue
			}
			fn(l.key, l.value)
		}
	} else {
		// internal node
		for i := 0; i < n.m; i++ {
			t.traverse(n.children[i].next, depth-1, fn)
		}
	}
}
