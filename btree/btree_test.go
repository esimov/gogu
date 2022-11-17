package btree

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree(t *testing.T) {
	assert := assert.New(t)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	btree := New[int, int]()
	assert.True(btree.IsEmpty())

	n := 100
	tmp := make(map[int]int, 0)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			key := i
			val := rand.Int()

			mu.Lock()
			btree.Put(key, val)
			tmp[key] = val
			mu.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()

	assert.Equal(n, btree.Size())

	btree.Traverse(func(key, val int) {
		v, found := btree.Get(key)
		assert.True(found)
		assert.Equal(v, val)

		btree.Remove(key)
		delete(tmp, key)
	})

	assert.Empty(btree.Size())
	assert.True(btree.IsEmpty())
}

func ExampleBTree() {
	btree := New[int, string]()
	fmt.Println(btree.IsEmpty())

	btree.Put(10, "foo")
	btree.Put(-1, "baz")
	btree.Put(2, "bar")
	btree.Put(-4, "qux")

	fmt.Println(btree.Size())

	tree := []string{}
	btree.Traverse(func(key int, val string) {
		item, _ := btree.Get(key)
		tree = append(tree, item)
	})
	fmt.Println(tree)

	// Output:
	// true
	// 4
	// [qux baz bar foo]
}
