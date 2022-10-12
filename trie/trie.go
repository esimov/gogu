// Package trie provides a concurrent safe implementation of the ternary search tree data structure.
// Trie is similar to binary search tree, but it has up to three children rather than two as of BST.
// Tries are used for locating specific keys from within a set or
// for quick lookup searches within a text like auto completion or spell checking.
package trie

import (
	"fmt"
	"sync"
)

var ErrorNotFound = fmt.Errorf("trie node not found")

// Queuer exposes the basic interface methods for querying the trie data structure
// both for searching and for retrieving the existing keys. These are generic methods
// having the same signature as the correspondig concrete methods from the queue package.
// Because both the plain array and the linked listed version of the queue package
// has the same method signature, each of them could be plugged in.
type Queuer[K ~string] interface {
	Enqueue(K)
	Dequeue() (K, error)
	Size() int
	Clear()
}

type node[K ~string, V any] struct {
	c       byte
	left    *node[K, V]
	mid     *node[K, V]
	right   *node[K, V]
	isValid bool
	Item[K, V]
}

// Item is a key-value struct pair used for storing the node values.
type Item[K ~string, V any] struct {
	key K
	val V
}

// newNode creates a new node.
func newNode[K ~string, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		Item: Item[K, V]{
			key: key,
			val: val,
		},
	}
}

// Trie is a lock-free tree data structure having the root as the first node.
// It's guarded with a mutex for concurrent data access.
type Trie[K ~string, V any] struct {
	n    int
	root *node[K, V]
	mu   *sync.RWMutex
	q    Queuer[K]
}

// New initializes a new Trie data structure.
func New[K ~string, V any](q Queuer[K]) *Trie[K, V] {
	return &Trie[K, V]{
		mu: &sync.RWMutex{},
		q:  q,
	}
}

// Size returns the trie size.
func (t *Trie[K, V]) Size() int {
	t.mu.RLock()
	defer t.mu.RUnlock()

	return t.n
}

// Contains checks if a key exists in the symbol table.
func (t *Trie[K, V]) Contains(key K) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if len(key) == 0 {
		return false
	}
	_, ok := t.Get(key)
	return ok
}

// Put inserts a new node into the symbol table, overwriting the old value
// with the new value if the key is already in the symbol table.
func (t *Trie[K, V]) Put(key K, val V) {
	if !t.Contains(key) {
		t.mu.Lock()
		t.n++
		t.mu.Unlock()
	}
	t.mu.Lock()
	t.root = t.root.put(t, key, val, 0, true)
	t.mu.Unlock()
}

func (n *node[K, V]) put(t *Trie[K, V], key K, val V, d int, isValid bool) *node[K, V] {
	c := key[d]
	if n == nil {
		n = newNode(key, val)
		n.c = c
	}

	if c < n.c {
		n.left = n.left.put(t, key, val, d, isValid)
	} else if c > n.c {
		n.right = n.right.put(t, key, val, d, isValid)
	} else if d < len(key)-1 {
		n.mid = n.mid.put(t, key, val, d+1, isValid)
	} else {
		n.isValid = isValid
		n.val = val
	}
	return n
}

// Get retrieves a node's value based on the key.
// If the key does not exists it returns false.
func (t *Trie[K, V]) Get(key K) (v V, ok bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if len(key) == 0 {
		return v, false
	}
	x, err := t.root.get(key, 0)
	if x == nil || err != nil {
		return v, false
	}

	return x.val, true
}

func (n *node[K, V]) get(key K, d int) (*node[K, V], error) {
	if n == nil {
		return nil, ErrorNotFound
	}
	if len(key) == 0 {
		return nil, fmt.Errorf("key for the get() method should not be empty")
	}
	c := key[d]

	if c < n.c {
		return n.left.get(key, d)
	} else if c > n.c {
		return n.right.get(key, d)
	} else if d < len(key)-1 {
		return n.mid.get(key, d+1)
	}
	return n, nil
}

// LongestPrefix returns the string in the symbol table that is the
// longest prefix of query, or empty if such string does not exists.
func (t *Trie[K, V]) LongestPrefix(query K) (K, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if len(query) == 0 {
		var k K
		return k, fmt.Errorf("query for the LongestPrefix() method should not be empty")
	}

	length := 0
	x := t.root
	i := 0
	for x != nil && i < len(query) {
		c := query[i]
		if c < x.c {
			x = x.left
		} else if c > x.c {
			x = x.right
		} else {
			i++
			if x.isValid {
				length = i
			}
			x = x.mid
		}
	}
	return query[:length], nil
}

// StartsWith returns all of the keys in the set that start with prefix.
func (t *Trie[K, V]) StartsWith(prefix K) (Queuer[K], error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	t.q.Clear()

	if len(prefix) == 0 {
		return t.q, fmt.Errorf("prefix for the StartsWith() method should not be empty")
	}

	x, err := t.root.get(prefix, 0)
	if x == nil || err != nil {
		return t.q, nil
	}
	if x.isValid {
		t.q.Enqueue(prefix)
	}
	x.mid.collect(t, prefix)

	return t.q, nil
}

// Keys collects all the existing keys in the set.
func (t *Trie[K, V]) Keys() (Queuer[K], error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	t.q.Clear()

	var err error
	t.root.collect(t, "")
	return t.q, err
}

func (n *node[K, V]) collect(t *Trie[K, V], prefix K) (Queuer[K], error) {
	if n == nil {
		return t.q, ErrorNotFound
	}

	n.left.collect(t, prefix)
	if n.isValid {
		t.q.Enqueue(prefix + K(n.c))
	}
	n.mid.collect(t, prefix+K(n.c))

	return n.right.collect(t, prefix)
}
