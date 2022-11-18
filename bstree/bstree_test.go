package bstree

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBSTree(t *testing.T) {
	assert := assert.New(t)

	bst := New[int, int](func(a, b int) bool {
		return a < b
	})
	tmp := make(map[int]int, 0)

	n := 100
	for i := 0; i < n; i++ {
		key := rand.Intn(n)
		val := rand.Int()
		bst.Upsert(key, val)
		tmp[key] = val
	}

	assert.Equal(len(tmp), bst.Size())

	for key, val := range tmp {
		node, err := bst.Get(key)
		assert.NoError(err)
		assert.Equal(node.Val, val)
	}

	bst.Traverse(func(item Item[int, int]) {
		assert.Equal(item.Val, tmp[item.Key])
	})

	for key := range tmp {
		err := bst.Delete(key)
		assert.NoError(err)
	}

	assert.Equal(0, bst.Size())
}

func TestBSTree_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	bst := New[int, int](func(a, b int) bool {
		return a < b
	})

	tmp := make(map[int]int, 0)
	n := 10

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			key := rand.Intn(n)
			val := rand.Int()
			bst.Upsert(key, val)
			mu.Lock()
			tmp[key] = val
			mu.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()

	for key, val := range tmp {
		wg.Add(1)
		go func(key, val int) {
			node, err := bst.Get(key)
			assert.NoError(err)
			assert.Equal(node.Val, val)

			wg.Done()
		}(key, val)
	}
	wg.Wait()

	bst.Traverse(func(item Item[int, int]) {
		assert.Equal(item.Val, tmp[item.Key])
	})

	for key := range tmp {
		err := bst.Delete(key)
		assert.NoError(err)
		delete(tmp, key)
	}

	assert.Empty(bst.Size())
}

func Example_BSTree() {
	bst := New[int, string](func(a, b int) bool {
		return a < b
	})

	bst.Upsert(10, "foo")
	bst.Upsert(-1, "baz")
	bst.Upsert(2, "bar")
	bst.Upsert(-4, "qux")

	fmt.Println(bst.Size())

	tree := []string{}
	bst.Traverse(func(item Item[int, string]) {
		node, _ := bst.Get(item.Key)
		tree = append(tree, node.Val)
	})
	fmt.Println(tree)

	for key := range tree {
		bst.Delete(key)
	}

	fmt.Println(bst.Size())

	// Output:
	// 4
	// [qux baz bar foo]
	// 0
}
