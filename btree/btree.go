package btree

import (
	"github.com/esimov/gogu"
	"golang.org/x/exp/constraints"
)

// Max children per binary tree must be even or greater than 2.
const maxChildren = 4

type entry[T any] struct {
	data T
	next *node[T]
}

type node[T any] struct {
	m        int // number of children
	children [maxChildren]entry[T]
}

func newNode[T any](m int) *node[T] {
	return &node[T]{
		m: m,
	}
}

type BTree[T constraints.Ordered] struct {
	n      int // the size of the binary tree (the number of nodes)
	heigth int // the height of the binary tree
	root   *node[T]
}

func New[T constraints.Ordered]() *BTree[T] {
	return &BTree[T]{
		root: newNode[T](0),
	}
}

func (t *BTree[T]) Size() int {
	return t.n
}

func (t *BTree[T]) IsEmpty() bool {
	return t.Size() == 0
}

func (t *BTree[T]) Height() int {
	return t.heigth
}

func (t *BTree[T]) Get(val T) (T, bool) {
	return t.search(t.root, val, t.heigth)
}

func (t *BTree[T]) search(n *node[T], data T, height int) (T, bool) {
	// extenernal node
	var val T
	if height == 0 {
		for i := 0; i < n.m; i++ {
			if gogu.Equal(data, n.children[i].data) {
				return n.children[i].data, true
			}
		}
	} else {
		// internal node
		for i := 0; i < n.m; i++ {
			if i+1 == n.m || gogu.Less(data, n.children[i+1].data) {
				return t.search(n.children[i].next, data, height-1)
			}
		}
	}

	return val, false
}

func (t *BTree[T]) Put(val T) {
	u := t.insert(t.root, val, t.heigth)
	t.n++
	if u == nil {
		return
	}
	// split the root
	n := newNode[T](2)
	n.children[0] = entry[T]{
		next: t.root,
	}
	n.children[1] = entry[T]{
		next: u,
	}

	t.root = n
	t.heigth++
}

func (t *BTree[T]) insert(n *node[T], data T, height int) *node[T] {
	entry := entry[T]{
		data: data,
		next: nil,
	}

	var j int
	// external node
	if height == 0 {
		for j = 0; j < n.m; j++ {
			if gogu.Equal(data, n.children[j].data) {
				var val T
				n.children[j].data = val
				return nil
			} else if gogu.Less(data, n.children[j].data) {
				break
			}
		}
	} else {
		// internal node
		for j = 0; j < n.m; j++ {
			if j+1 == n.m || gogu.Less(data, n.children[j+1].data) {
				node := t.insert(n.children[j].next, data, height-1)
				if node == nil {
					return nil
				}
				j++
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

func (t *BTree[T]) split(n *node[T]) *node[T] {
	h := newNode[T](maxChildren / 2)
	n.m = maxChildren / 2

	for i := 0; i < n.m; i++ {
		h.children[i] = n.children[n.m+i]
	}
	return h
}

func (t *BTree[T]) Remove(val T) {
	_, ok := t.Get(val)
	if !ok {
		return
	}
	t.insert(t.root, val, t.heigth)
	t.n--
}

func (t *BTree[T]) Traverse(fn func(data T)) {
	t.traverse(t.root, t.heigth, fn)
}

func (t *BTree[T]) traverse(n *node[T], depth int, fn func(data T)) {
	// extenernal node
	var val T
	if depth == 0 {
		for i := 0; i < n.m; i++ {
			if gogu.Equal(val, n.children[i].data) {
				continue
			}
			fn(n.children[i].data)
		}
	} else {
		// internal node
		for i := 0; i < n.m; i++ {
			t.traverse(n.children[i].next, depth-1, fn)
		}
	}
}
