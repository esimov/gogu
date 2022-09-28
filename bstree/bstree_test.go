package bstree

import (
	"math/rand"
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
