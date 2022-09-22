package btree

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree(t *testing.T) {
	assert := assert.New(t)

	btree := New[int, int]()
	assert.True(btree.IsEmpty())

	n := 100
	tmp := make(map[int]int, 0)

	for i := 0; i < n; i++ {
		key := i
		val := rand.Int()
		btree.Put(key, val)
		tmp[key] = val
	}
	assert.False(btree.IsEmpty())
	assert.Equal(100, btree.Size())

	btree.Traverse(func(key, val int) {
		assert.Equal(tmp[key], val)
	})

	for key, val := range tmp {
		v, found := btree.Get(key)
		assert.True(found)
		assert.Equal(v, val)

		btree.Remove(key)
		delete(tmp, key)
	}

	assert.Equal(0, btree.Size())
	assert.True(btree.IsEmpty())
}
