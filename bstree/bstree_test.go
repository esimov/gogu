package bstree

import (
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
		assert.Equal(node.val, val)
	}

	bst.Traverse(func(item Item[int, int]) {
		assert.Equal(item.val, tmp[item.key])
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
			assert.Equal(node.val, val)

			wg.Done()
		}(key, val)
	}
	wg.Wait()

	bst.Traverse(func(item Item[int, int]) {
		assert.Equal(item.val, tmp[item.key])
	})

	for key := range tmp {
		err := bst.Delete(key)
		assert.NoError(err)
		delete(tmp, key)
	}

	assert.Empty(bst.Size())
}
