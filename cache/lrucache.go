package cache

import (
	"errors"
)

// node is the entity used for storing elements in the cache. It acts both as a wrapper in for elements in a map
// as well is part of an implementation of a double linked list
type node[K comparable, V any] struct {
	next, prev *node[K, V]
	list       *lruList[K, V]

	key   K
	value V
}

// lruList is an implementation for inner doubly linked list
type lruList[K comparable, V any] struct {
	root node[K, V] // root element in the list. It should not be removed or changed
	len  int
}

// newLRUList initializes a new double linked list with the root element and with an initial size of 0
func newLRUList[K comparable, V any]() *lruList[K, V] {
	// Initiate the root element
	lst := lruList[K, V]{root: node[K, V]{}, len: 0}
	lst.root.prev = &lst.root
	lst.root.next = &lst.root

	return &lst
}

// moveAfter moves nd node after the current node
func (l *lruList[K, V]) moveAfter(current *node[K, V], nd *node[K, V]) {
	if current == nd {
		return
	}

	nd.prev.next = nd.next
	nd.next.prev = nd.prev

	nd.prev = current
	nd.next = current.next
	nd.prev.next = nd
	nd.next.prev = nd
}

// moveFront moves nd at the front of the list (after the root element)
func (l *lruList[K, V]) moveFront(nd *node[K, V]) {
	l.moveAfter(&l.root, nd)
}

// length return the number of elements in the list
func (l *lruList[K, V]) length() int {
	return l.len
}

// addAfter adds a new element after the current node
func (l *lruList[K, V]) addAfter(current *node[K, V], key K, value V) *node[K, V] {
	newNode := node[K, V]{
		prev:  current,
		next:  current.next,
		list:  l,
		key:   key,
		value: value,
	}
	current.next.prev = &newNode
	current.next = &newNode
	l.len++
	return &newNode
}

// addFront adds a new element to the front of the list (after the root node)
func (l *lruList[K, V]) addFront(key K, value V) *node[K, V] {
	x := l.addAfter(&l.root, key, value)
	return x
}

// last returns the last node from the list
func (l *lruList[K, V]) last() *node[K, V] {
	return l.root.prev
}

// first returns the first node from the list
func (l *lruList[K, V]) first() *node[K, V] {
	return l.root.next
}

// remove removes the nd node form the list
func (l *lruList[K, V]) remove(nd *node[K, V]) bool {
	if nd != &l.root {
		nd.prev.next = nd.next
		nd.next.prev = nd.prev
		l.len--
		return true
	}
	return false
}

// removeLast removes the last node from the list
func (l *lruList[K, V]) removeLast() bool {
	return l.remove(l.last())
}

// LRUCache implements a fixed size LRU cache using a map and a double linked list
type LRUCache[K comparable, V any] struct {
	items     map[K]*node[K, V]
	evictList *lruList[K, V]
	size      int
}

// NewLRU initializes a new LRU cache
func NewLRU[K comparable, V any](size int) (*LRUCache[K, V], error) {
	if size <= 0 {
		return nil, errors.New("size must be a positive value")
	}

	lru := &LRUCache[K, V]{
		items:     make(map[K]*node[K, V]),
		evictList: newLRUList[K, V](),
		size:      size,
	}

	return lru, nil
}

// Add adds a value to the cache. If the oldest value is evicted, this value and they key for it is returned.
func (c *LRUCache[K, V]) Add(key K, value V) (oldestKey K, oldestValue V, removed bool) {
	// If the element is in the cache, move it to the front and return
	if item, ok := c.items[key]; ok {
		c.evictList.moveFront(item)
		item.value = value
		return
	}

	// Since this is a new element, put this to the front
	item := c.evictList.addFront(key, value)
	c.items[key] = item

	// Remove the oldest element if the cache is full
	if c.Count() > c.size {
		return c.RemoveOldest()
	}
	return
}

// Count return the number of the current values from the cache. It should be LE then the initial size of the cache
func (c *LRUCache[K, V]) Count() int {
	return c.evictList.len
}

// GetOldest returns the oldest key/value pair from the cache if the cache has any values
func (c *LRUCache[K, V]) GetOldest() (key K, value V, available bool) {
	if item := c.evictList.last(); item != &c.evictList.root {
		// Since the oldest was touched, it is not the oldest anymore so move it to the front
		c.evictList.moveFront(item)
		return item.key, item.value, true
	}
	return
}

// Get return the element for the key if the element is present in the cache
func (c *LRUCache[K, V]) Get(key K) (value V, available bool) {
	if item, ok := c.items[key]; ok {
		// The item was touched, move it to the front in the list
		c.evictList.moveFront(item)
		return item.value, true
	}
	return
}

// GetYoungest returns the youngest key/value pair from the cache if the cache has any values
func (c *LRUCache[K, V]) GetYoungest() (key K, value V, available bool) {
	if item := c.evictList.first(); item != &c.evictList.root {
		return item.key, item.value, true
	}
	return
}

// RemoveOldest removes the oldest value from the cache. It returns he key/value pair which was removed
func (c *LRUCache[K, V]) RemoveOldest() (key K, value V, removed bool) {
	if item := c.evictList.last(); item != &c.evictList.root {
		delete(c.items, item.key)
		return item.key, item.value, c.evictList.removeLast()
	}
	return
}

// Remove removes an element form the cache denoted by the key. The value removed is returned
func (c *LRUCache[K, V]) Remove(key K) (value V, removed bool) {
	if item, ok := c.items[key]; ok {
		// The item was touched, move it to the front in the list
		delete(c.items, item.key)
		c.evictList.remove(item)
		return item.value, true
	}
	return
}

// RemoveYoungest removes the youngest value from the cache. The key/value pair removed is returned
func (c *LRUCache[K, V]) RemoveYoungest() (key K, value V, removed bool) {
	if item := c.evictList.first(); item != &c.evictList.root {
		delete(c.items, item.key)
		return item.key, item.value, c.evictList.removeLast()
	}
	return
}

// Flush clears all values from the cache
func (c *LRUCache[K, V]) Flush() {
	c.items = make(map[K]*node[K, V])
	c.evictList = newLRUList[K, V]()
}
