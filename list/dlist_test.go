package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedList(t *testing.T) {
	assert := assert.New(t)

	list := InitDoubleList(1)
	assert.Equal(1, list.data)

	err := list.Delete(&list.doubleNode)
	assert.Error(err)

	node := list.Push(2)
	list.InsertBefore(node, 3)

	n := 3
	list.Each(func(i int) {
		assert.Equal(n, i)
		n--
	})

	node = list.Append(4)
	assert.Equal(4, node.data)

	n = 0
	expected := []int{3, 2, 1, 4}
	list.Each(func(i int) {
		assert.Equal(expected[n], i)
		n++
	})

	n1, found := list.Find(4)
	assert.Equal(4, n1.data)
	assert.True(found)

	n2, found := list.Find(10)
	assert.Nil(n2)
	assert.False(found)

	err = list.InsertAfter(node, 6)
	assert.NoError(err)

	n = 0
	expected = []int{3, 2, 1, 4, 6}
	list.Each(func(i int) {
		assert.Equal(expected[n], i)
		n++
	})

	n3 := list.Push(7)
	assert.Equal(7, n3.data)

	n = 0
	expected = []int{7, 3, 2, 1, 4, 6}
	list.Each(func(i int) {
		assert.Equal(expected[n], i)
		n++
	})

	list.Shift()
	list.Shift()
	n = 0
	expected = []int{2, 1, 4, 6}
	list.Each(func(i int) {
		assert.Equal(expected[n], i)
		n++
	})

	list.Pop()
	n = 0
	expected = []int{2, 1, 4}
	list.Each(func(i int) {
		assert.Equal(expected[n], i)
		n++
	})

	err = list.Delete(node)
	assert.NoError(err)

	// node with value 4 has been already removed from the list
	err = list.Delete(node)
	assert.Error(err)

	err = list.InsertBefore(node, 11111)
	assert.Error(err)

	_, err = list.Replace(10, 0)
	assert.Error(err)

	list.Replace(2, 0)
	n = 0
	expected = []int{0, 1}
	list.Each(func(i int) {
		assert.Equal(expected[n], i)
		n++
	})
}
