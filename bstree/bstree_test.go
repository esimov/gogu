package bstree

import (
	//"fmt"
	"fmt"
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

	n := 10
	for i := 0; i < n; i++ {
		key := rand.Intn(n)
		val := rand.Int()
		bst.Insert(key, val)
		tmp[key] = val
	}
	fmt.Println("init size:", bst.Size())

	for key, val := range tmp {
		node, err := bst.Get(key)
		assert.NoError(err)
		assert.Equal(node.val, val)
	}

	bst.Traverse(func(item Item[int, int]) {
		assert.Equal(item.val, tmp[item.key])
	})
	for key := range tmp {
		fmt.Println(key)
		err := bst.Delete(key)
		assert.NoError(err)
	}

	// fmt.Println("after size:", bst.Size())

	// bst.Traverse(func(item Item[int, int]) {
	// 	fmt.Println(item)
	// 	assert.Equal(item.val, tmp[item.key])
	// })

	fmt.Println(bst.Size())
	assert.Equal(1, 1)
}
