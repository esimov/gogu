package btree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTree(t *testing.T) {
	assert := assert.New(t)

	btree := New[int]()
	// btree.Put(30)
	// btree.Put(50)
	// btree.Put(10)
	// btree.Put(20)
	// btree.Put(40)
	// btree.Put(2)
	// btree.Put(9)

	for i := 0; i < 100; i++ {
		btree.Put(i)
	}

	fmt.Println(btree.Size())
	fmt.Println(btree.Get(4))

	assert.Equal(1, 1)

	fmt.Println("===============")

	btree.Remove(40)
	btree.Remove(9)
	btree.Traverse(func(data int) {
		fmt.Println(data)
	})
	fmt.Println(btree.Height())
	fmt.Println(btree.Size())
}
