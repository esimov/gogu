// Package btree provides an implementation of the B-tree data structure,
// which is a self-balancing tree data structure maintaining its values
// in sorted order and allowing each node to have more than two children,
// compared to the standard BST where each node has only two leaves.
// The implementation is an adapted version of https://algs4.cs.princeton.edu/62btree/BTree.java.
// This package is NOT thread-safe.
// For data consistency some sort of concurrency safe mechanism should be implemented on the client side.
package btree

import (
	"github.com/esimov/torx"
	"golang.org/x/exp/constraints"
)

// Max children per binary tree. Must be even or greater than 2.
const maxChildren = 4

// entry is the inner component of a node, which holds the node value and a pointer to the next node.
type entry[K constraints.Ordered, V any] struct {
	key       K
	value     V
	isRemoved bool
	next      *node[K, V]
}

// node is a data structure which defines how many children (leaves) each node has.
type node[K constraints.Ordered, V any] struct {
	m        int // number of children
	children [maxChildren]entry[K, V]
}

// newNode instantiates a new node with no leaves.
func newNode[K constraints.Ordered, V any](m int) *node[K, V] {
	return &node[K, V]{
		m: m,
	}
}

// BTree defines a data structure with one node, which is the root node.
type BTree[K constraints.Ordered, V any] struct {
	n      int // the size of the tree (the number of nodes)
	height int // the height of the tree
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
	return t.height
}

// Get searches for a key and in case it's found it returns the key's value
// together with a boolean flag signalig the key existence in the tree data structure.
func (t *BTree[K, V]) Get(key K) (V, bool) {
	return t.root.search(t, key, t.height)
}

// search is a private method which is invoked by the Get method.
func (n *node[K, V]) search(t *BTree[K, V], key K, height int) (V, bool) {
	// external node
	if height == 0 {
		for i := 0; i < n.m; i++ {
			if torx.Equal(key, n.children[i].key) {
				return n.children[i].value, true
			}
		}
	} else {
		// internal node
		for i := 0; i < n.m; i++ {
			if i+1 == n.m || torx.Less(key, n.children[i+1].key) {
				return n.children[i].next.search(t, key, height-1)
			}
		}
	}

	var v V
	return v, false
}

// Put inserts a new value into the B-tree.
func (t *BTree[K, V]) Put(key K, val V) {
	u := t.root.insert(t, key, val, t.height, false)
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
	t.height++
}

// insert is a private method which is invoked by the Put method.
func (n *node[K, V]) insert(t *BTree[K, V], key K, val V, height int, isRemoved bool) *node[K, V] {
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
			if torx.Equal(key, n.children[j].key) {
				n.children[j].value = val
				// This signals that we are invoking the Put or Remove method.
				n.children[j].isRemoved = isRemoved
				return nil
			} else if torx.Less(key, n.children[j].key) {
				break
			}
		}
	} else {
		// internal node
		for j = 0; j < n.m; j++ {
			if j+1 == n.m || torx.Less(key, n.children[j+1].key) {
				node := n.children[j].next.insert(t, key, val, height-1, isRemoved)
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

// Remove deletes a node from the B-tree.
func (t *BTree[K, V]) Remove(key K) {
	val, ok := t.Get(key)
	if !ok {
		return
	}
	t.n--
	t.root.insert(t, key, val, t.height, true)
}

// Traverse iterates over the tree nodes and invokes the callback function provided as argument.
func (t *BTree[K, V]) Traverse(fn func(key K, val V)) {
	t.traverse(t.root, t.height, fn)
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
