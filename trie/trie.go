package trie

import (
	"fmt"
)

var ErrorNotFound = fmt.Errorf("trie node not found")

type Queuer[K ~string] interface {
	Enqueue(K)
	Dequeue() K
}

type node[K ~string, V any] struct {
	c       byte
	left    *node[K, V]
	mid     *node[K, V]
	right   *node[K, V]
	isValid bool
	Item[K, V]
}

type Item[K ~string, V any] struct {
	key K
	val V
}

func newNode[K ~string, V any](key K, val V) *node[K, V] {
	return &node[K, V]{
		Item: Item[K, V]{
			key: key,
			val: val,
		},
	}
}

type Trie[K ~string, V any] struct {
	n    int
	root *node[K, V]
}

func New[K ~string, V any]() *Trie[K, V] {
	return &Trie[K, V]{}
}

func (t *Trie[K, V]) Size() int {
	return t.n
}

func (t *Trie[K, V]) Contains(key K) bool {
	if len(key) == 0 {
		return false
	}
	_, ok := t.Get(key)
	return ok
}

func (t *Trie[K, V]) Put(key K, val V) {
	if !t.Contains(key) {
		t.n++
	}
	t.root = t.root.put(key, val, 0, true)
}

func (n *node[K, V]) put(key K, val V, d int, isValid bool) *node[K, V] {
	c := key[d]
	if n == nil {
		n = newNode(key, val)
		n.c = c
	}

	if c < n.c {
		n.left = n.left.put(key, val, d, isValid)
	} else if c > n.c {
		n.right = n.right.put(key, val, d, isValid)
	} else if d < len(key)-1 {
		n.mid = n.mid.put(key, val, d+1, isValid)
	} else {
		n.isValid = isValid
		n.val = val
	}
	return n
}

func (t *Trie[K, V]) Get(key K) (v V, ok bool) {
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

func (t *Trie[K, V]) LongestPrefix(query K) (K, error) {
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

func (t *Trie[K, V]) Keys(q Queuer[K]) (Queuer[K], error) {
	return t.root.collect(q, "")
}

// Returns all of the keys in the set that start with prefix.
func (t *Trie[K, V]) StartsWith(q Queuer[K], prefix K) (Queuer[K], error) {
	if len(prefix) == 0 {
		var q Queuer[K]
		return q, fmt.Errorf("prefix for the StartsWith() method should not be empty")
	}

	x, err := t.root.get(prefix, 0)
	if x == nil || err != nil {
		return q, nil
	}
	if x.isValid {
		q.Enqueue(prefix)
	}
	x.mid.collect(q, prefix)

	return q, nil
}

func (n *node[K, V]) collect(q Queuer[K], prefix K) (Queuer[K], error) {
	if n == nil {
		var q Queuer[K]
		return q, ErrorNotFound
	}

	n.left.collect(q, prefix)
	if n.isValid {
		q.Enqueue(prefix + K(n.c))
	}
	n.mid.collect(q, prefix+K(n.c))

	return n.right.collect(q, prefix)
}
