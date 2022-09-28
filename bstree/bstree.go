// Package bstree provides an implementation of the Binary Search Tree (BST)
// data structure algorithm, where each node has at most two child nodes and
// the key of its internal node is greater than all the keys in the respective
// node's left subtree and less than the ones in the right subtree.
package bstree

import (
	"fmt"
	"sync"

	"github.com/esimov/gogu"
	"golang.org/x/exp/constraints"
)

var ErrorNotFound = fmt.Errorf("node not found")

// Item contains the node data as a key-value pair.
type Item[K constraints.Ordered, V any] struct {
	key K
	val V
}

// node represents the BST internal node, having as components the node item defined
// as a key-value pair and two separate pointers to the left and right child nodes.
type node[K constraints.Ordered, V any] struct {
	left  *node[K, V]
	right *node[K, V]
	Item[K, V]
}

// newNode creates a new node.
func newNode[K constraints.Ordered, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		Item: Item[K, V]{
			key: key,
			val: val,
		},
	}
}

// bsTree is the basic component for the BST data structure initialization.
// It incorporates a concurrent safe mechanism using sync.Mutex to guarantee
// the data consistency on concurrent read and write access.
type bsTree[K constraints.Ordered, V any] struct {
	mu   *sync.RWMutex
	comp gogu.CompFn[K]
	root *node[K, V]
	size int
}

// New initializes a new BST data structure together with a comparison operator.
// Depending on the comparator it sorts the tree in ascending or descending order.
func New[K constraints.Ordered, V any](comp gogu.CompFn[K]) *bsTree[K, V] {
	return &bsTree[K, V]{
		mu:   &sync.RWMutex{},
		comp: comp,
	}
}

// Size returns the size of the tree.
func (b *bsTree[K, V]) Size() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.size
}

// Get retrieves the node item and an error in case the requested node does not exists.
func (b *bsTree[K, V]) Get(key K) (Item[K, V], error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.root.get(b, key)
}

func (n *node[K, V]) get(b *bsTree[K, V], key K) (Item[K, V], error) {
	if n == nil {
		var it Item[K, V]
		return it, ErrorNotFound
	}

	if gogu.Compare(key, n.key, b.comp) == 1 {
		return n.left.get(b, key)
	} else if gogu.Compare(key, n.key, b.comp) == -1 {
		return n.right.get(b, key)
	}

	return n.Item, nil
}

// Upsert insert a new node, or update an existing node in case the key is found in the tree list.
func (b *bsTree[K, V]) Upsert(key K, val V) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.root == nil {
		b.root = newNode(key, val)
		b.size++
	} else {
		b.root.upsert(b, key, val)
	}

}

func (n *node[K, V]) upsert(b *bsTree[K, V], key K, val V) {
	if gogu.Compare(key, n.key, b.comp) == 1 {
		if n.left == nil {
			n.left = newNode(key, val)
			b.size++
		} else {
			n.left.upsert(b, key, val)
		}
	} else if gogu.Compare(key, n.key, b.comp) == -1 {
		if n.right == nil {
			n.right = newNode(key, val)
			b.size++
		} else {
			n.right.upsert(b, key, val)
		}
	} else {
		n.val = val
	}
}

// isLeaf checks if a node is a leaf node.
func (n *node[K, V]) isLeaf() bool {
	return !n.hasLeft() && !n.hasRight()
}

// hasLeft checks if a node has children on left branch.
func (n *node[K, V]) hasLeft() bool {
	return n.left != nil
}

// hasRight checks if a node has children on right branch.
func (n *node[K, V]) hasRight() bool {
	return n.right != nil
}

// min searches for the latest node on the left branch, but considering that BST
// is an ordered tree structure it happens that it contains also the smallest value.
func (n *node[K, V]) min() *node[K, V] {
	for ; n.left != nil; n = n.left {
	}
	return n
}

// Delete removes a node defined by the key.
func (b *bsTree[K, V]) Delete(key K) error {
	var err error
	b.mu.RLock()
	b.root, err = b.root.delete(b, key)
	b.size--
	b.mu.RUnlock()

	return err
}

func (n *node[K, V]) delete(b *bsTree[K, V], key K) (*node[K, V], error) {
	var err error
	if n == nil {
		return nil, ErrorNotFound
	}

	if gogu.Compare(key, n.key, b.comp) == 1 {
		n.left, err = n.left.delete(b, key)
		return n, err
	} else if gogu.Compare(key, n.key, b.comp) == -1 {
		n.right, err = n.right.delete(b, key)
		return n, err
	} else {

		// case 1: node has no child
		if n.isLeaf() {
			return nil, nil
		}
		// case 2a: node has left child only
		if n.hasLeft() && !n.hasRight() {
			return n.left, nil
		}
		// case 2b: node has right child only
		if !n.hasLeft() && n.hasRight() {
			return n.right, nil
		}
		// case 3: node with two children
		// Get the latest value on the left branch, which,
		// following the BST rules, should have the smallest value.
		min := n.right.min()
		n.key = min.key
		n.val = min.val
		// Delete the inorder successor.
		n.right, err = n.right.delete(b, min.key)

		return n, err
	}
}

// Traverse iterates over the tree structure and invokes
// the callback function provided as a parameter.
func (b *bsTree[K, V]) Traverse(fn func(Item[K, V])) {
	ch := make(chan Item[K, V])
	n := b.root
	go func() {
		b.mu.RLock()
		n.traverse(b, ch)
		b.mu.RUnlock()

		close(ch)
	}()

	for item := range ch {
		fn(item)
	}
}

func (n *node[K, V]) traverse(b *bsTree[K, V], ch chan<- Item[K, V]) {
	if n == nil {
		return
	}
	n.left.traverse(b, ch)
	ch <- Item[K, V]{
		key: n.key,
		val: n.val,
	}
	n.right.traverse(b, ch)
}
