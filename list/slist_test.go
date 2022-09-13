package gogu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	assert := assert.New(t)

	list := InitList(1)
	assert.Equal(1, list.data)
	list.Append(2)
	err := list.Delete(&list.node) // delete first node
	assert.NoError(err)
	assert.Equal(2, list.node.data)

	list.Unshift(1)
	assert.Equal(1, list.node.data)

	last := list.Append(3)
	err = list.InsertAfter(last, 4)
	assert.NoError(err)
	list.Append(5)

	last = list.Append(6)
	list.Append(8)
	list.InsertAfter(last, 7)
	last = list.Append(9)

	err = list.Delete(last)
	assert.NoError(err)

	i := 0
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item := list.DeleteLast()
	assert.Equal(6, item.data)

	i = 0
	expected = []int{1, 2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item = list.DeleteFirst()
	assert.Equal(1, item.data)

	i = 0
	expected = []int{2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})
}
