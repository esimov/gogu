package btree

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree(t *testing.T) {
	assert := assert.New(t)

	btree := New[int, int]()
	assert.True(btree.IsEmpty())

	n := 10
	tmp := make(map[int]int, 0)

	for i := 0; i < n; i++ {
		key := i
		val := rand.Int()
		//fmt.Println(val)
		btree.Put(key, val)
		tmp[key] = val
	}
	fmt.Println(tmp)
	assert.False(btree.IsEmpty())
	//assert.Equal(10, btree.Size())

	btree.Traverse(func(key, val int) {
		fmt.Println(key, val)
		//assert.Equal(tmp[key], val)
	})

	// fmt.Println(tmp)
	// for key, val := range tmp {
	// 	//fmt.Println("Key:", key, "val:", val)
	// 	v, found := btree.Get(key)
	// 	assert.True(found)
	// 	assert.Equal(v, val)

	// 	btree.Remove(key)
	// 	delete(tmp, key)
	// }

	// assert.Equal(0, btree.Size())
	//assert.True(btree.IsEmpty())
}
