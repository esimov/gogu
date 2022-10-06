package trie

import (
	"fmt"
	"testing"

	"github.com/esimov/gogu/queue"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	assert := assert.New(t)

	trie := New[string, int]()
	trie.Put("cats", 1)
	trie.Put("cape", 2)
	trie.Put("captain", 3)
	trie.Put("foes", 4)
	trie.Put("she", 5)
	trie.Put("shells", 6)
	trie.Put("the", 7)
	trie.Put("thermos", 8)
	trie.Put("fo", 9)
	trie.Put("foo", 10)
	fmt.Println("Size:", trie.Size())
	v, ok := trie.Get("cats")
	fmt.Println(v, ok)
	res, _ := trie.LongestPrefix("thermostat")
	fmt.Println(res)

	q := queue.New[string]()

	q1, err := trie.StartsWith(q, "ca")
	if err == nil {
		fmt.Println("New Size:", q1.Size())
		val, _ := q1.Dequeue()
		fmt.Println("value:", val)
		for i := 0; i < q1.Size(); i++ {
			val, _ := q1.Dequeue()
			fmt.Println("value:", val)
		}
	}

	q2, err := trie.Keys(q)
	fmt.Println(q2)
	if err == nil {
		fmt.Println(q2.Size())
	}
	q2.Dequeue()
	fmt.Println("q:", q2)

	assert.Equal(1, 1)
}
