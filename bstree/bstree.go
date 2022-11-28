// Package bstree provides an implementation of the Binary Search Tree (BST)
// data structure algorithm, where each node has at most two child nodes and
// the key of its internal node is greater than all the keys in the respective
// node's left subtree and less than the ones in the right subtree.
package bstree

import (
	"fmt"
	"sync"

	"github.com/esimov/torx"
	"golang.org/x/exp/constraints"
)

var ErrorNotFound = fmt.Errorf("BST node not found")

// Item contains the node's data as a key-value pair data structure.
type Item[K constraints.Ordered, V any] struct {
	Key K
	Val V
}

// Node represents the BST internal Node, having as components the Node item defined
// as a key-value pair and two separate pointers to the left and right child nodes.
type Node[K constraints.Ordered, V any] struct {
	Left  *Node[K, V]
	Right *Node[K, V]
	Item[K, V]
}

// NewNode creates a new node.
func NewNode[K constraints.Ordered, V any](key K, val V) *Node[K, V] {
	return &Node[K, V]{
		Item: Item[K, V]{
			Key: key,
			Val: val,
		},
	}
}

// BsTree is the basic component for the BST data structure initialization.
// It incorporates a thread safe mechanism using `sync.Mutex` to guarantee
// the data consistency on concurrent read and write operation.
type BsTree[K constraints.Ordered, V any] struct {
	mu   sync.RWMutex
	comp torx.CompFn[K]
	root *Node[K, V]
	size int
}

// New initializes a new BST data structure together with a comparison operator.
// Depending on the comparator it sorts the tree in ascending or descending order.
func New[K constraints.Ordered, V any](comp torx.CompFn[K]) *BsTree[K, V] {
	return &BsTree[K, V]{
		mu:   sync.RWMutex{},
		comp: comp,
	}
}

// Size returns the size of the tree.
func (b *BsTree[K, V]) Size() int {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.size
}

// Get retrieves the node item and an error in case the requested node does not exists.
func (b *BsTree[K, V]) Get(key K) (Item[K, V], error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.root.get(b, key)
}

func (n *Node[K, V]) get(b *BsTree[K, V], key K) (Item[K, V], error) {
	if n == nil {
		var it Item[K, V]
		return it, ErrorNotFound
	}

	if torx.Compare(key, n.Key, b.comp) == 1 {
		return n.Left.get(b, key)
	} else if torx.Compare(key, n.Key, b.comp) == -1 {
		return n.Right.get(b, key)
	}

	return n.Item, nil
}

// Upsert insert a new node or update an existing node in case the key is found in the tree list.
func (b *BsTree[K, V]) Upsert(key K, val V) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.root == nil {
		b.root = NewNode(key, val)
		b.size++
	} else {
		b.root.upsert(b, key, val)
	}

}

func (n *Node[K, V]) upsert(b *BsTree[K, V], key K, val V) {
	if torx.Compare(key, n.Key, b.comp) == 1 {
		if n.Left == nil {
			n.Left = NewNode(key, val)
			b.size++
		} else {
			n.Left.upsert(b, key, val)
		}
	} else if torx.Compare(key, n.Key, b.comp) == -1 {
		if n.Right == nil {
			n.Right = NewNode(key, val)
			b.size++
		} else {
			n.Right.upsert(b, key, val)
		}
	} else {
		n.Val = val
	}
}

// min searches for the latest node on the left branch, but considering that BST
// is an ordered tree data structure it happens that it holds also the smallest value.
func (n *Node[K, V]) min() *Node[K, V] {
	for ; n.Left != nil; n = n.Left {
	}
	return n
}

// Delete removes a node defined by its key from the tree structure.
func (b *BsTree[K, V]) Delete(key K) error {
	var err error
	b.mu.RLock()
	b.root, err = b.root.delete(b, key)
	b.size--
	b.mu.RUnlock()

	return err
}

func (n *Node[K, V]) delete(b *BsTree[K, V], key K) (*Node[K, V], error) {
	var err error
	if n == nil {
		return nil, ErrorNotFound
	}

	if torx.Compare(key, n.Key, b.comp) == 1 {
		n.Left, err = n.Left.delete(b, key)
		return n, err
	} else if torx.Compare(key, n.Key, b.comp) == -1 {
		n.Right, err = n.Right.delete(b, key)
		return n, err
	} else {
		// case 1: node has no child
		if n.Left == nil && n.Right == nil {
			return nil, nil
		}
		// case 2a: node has left child only
		if n.Left != nil && n.Right == nil {
			return n.Left, nil
		}
		// case 2b: node has right child only
		if n.Left == nil && n.Right != nil {
			return n.Right, nil
		}
		// case 3: node with two children
		// Get the latest value on the left branch, which,
		// following the BST rules, should have the smallest value.
		min := n.Right.min()
		n.Key = min.Key
		n.Val = min.Val
		// Delete the inorder successor.
		n.Right, err = n.Right.delete(b, min.Key)

		return n, err
	}
}

// Traverse iterates over the tree structure and invokes the callback function provided as a parameter.
func (b *BsTree[K, V]) Traverse(fn func(Item[K, V])) {
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

func (n *Node[K, V]) traverse(b *BsTree[K, V], ch chan<- Item[K, V]) {
	if n == nil {
		return
	}
	n.Left.traverse(b, ch)
	ch <- Item[K, V]{
		Key: n.Key,
		Val: n.Val,
	}
	n.Right.traverse(b, ch)
}
