package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	assert := assert.New(t)

	list := Init(1)
	assert.Equal(1, list.data)
	// cannot delete the first node if there is only one item in the list.
	list.Pop()
	err := list.Delete(&list.singleNode)
	assert.Error(err)

	list.Append(2)
	err = list.Delete(&list.singleNode) // delete first node
	assert.NoError(err)
	assert.Equal(2, list.singleNode.data)

	list.Unshift(1)
	assert.Equal(1, list.singleNode.data)

	list.Append(3)
	last, _ := list.Find(3)
	err = list.InsertAfter(last, 4)
	assert.NoError(err)
	list.Append(5)

	list.Append(6)
	last, _ = list.Find(6)
	list.Append(8)
	list.InsertAfter(last, 7)
	list.Append(9)
	last, _ = list.Find(9)

	err = list.Delete(last)
	assert.NoError(err)

	i := 0
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	list.Pop()

	i = 0
	expected = []int{1, 2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	list.Shift()

	i = 0
	expected = []int{2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	err = list.Replace(20, 10)
	assert.Error(err)
	item, _ := list.Find(20)
	assert.Nil(item)

	err = list.Replace(7, 8)
	assert.NoError(err)
	item, _ = list.Find(8)
	assert.Equal(8, item.data)

	list.Unshift(1)
	list.Replace(8, 7)

	i = 0
	expected = []int{1, 2, 3, 4, 5, 6, 7}
	list.Each(func(val int) {
		assert.Equal(expected[i], val)
		i++
	})

	item, found := list.Find(7)
	assert.Equal(7, item.data)
	assert.True(found)

	item, found = list.Find(22)
	assert.Nil(item)
	assert.False(found)
}

func Example_singlyLinkedList() {
	list := Init(1)

	values := []int{2, 3, 4, 5, 6, 7, 8}
	for _, val := range values {
		list.Append(val)
	}
	sl := []int{}
	list.Each(func(val int) {
		sl = append(sl, val)
	})
	fmt.Println(sl)

	list.Pop()

	sl = nil
	list.Each(func(val int) {
		sl = append(sl, val)
	})
	fmt.Println(sl)

	list.Shift()

	sl = nil
	list.Each(func(val int) {
		sl = append(sl, val)
	})
	fmt.Println(sl)

	err := list.Replace(20, 10)
	fmt.Println(err)
	item, _ := list.Find(20)
	fmt.Println(item)

	list.Replace(7, 8)
	item, _ = list.Find(8)
	fmt.Println(item.data)

	item, _ = list.Find(8)
	fmt.Println(item.data)

	// Output:
	// [1 2 3 4 5 6 7 8]
	// [1 2 3 4 5 6 7]
	// [2 3 4 5 6 7]
	// requested node does not exists
	// <nil>
	// 8
	// 8
}
