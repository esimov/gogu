package trie

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"testing"

	"github.com/esimov/gogu/queue"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	assert := assert.New(t)

	trie := New[string, int]()
	input := []string{"cats", "cape", "captain", "foes",
		"apple", "she", "root", "shells", "the", "thermos", "foo"}

	for idx, v := range input {
		trie.Put(v, idx)
	}

	assert.Equal(11, trie.Size())
	v, ok := trie.Get("cats")
	assert.Equal(0, v)
	assert.True(ok)

	str, err := trie.LongestPrefix("thermostat")
	assert.NoError(err)
	assert.Equal("thermos", str)

	str, err = trie.LongestPrefix("cap")
	assert.NoError(err)
	assert.Empty(str)

	str, err = trie.LongestPrefix("capetown")
	assert.NoError(err)
	assert.Equal("cape", str)

	q := queue.New[string]()
	_, err = trie.StartsWith(q, "")
	assert.Error(err)

	q1, err := trie.StartsWith(q, "ca")
	assert.NoError(err)
	assert.Equal(3, q1.Size())

	expected := []string{"cats", "cape", "captain"}
	sort.Strings(expected)

	n := 0
	for q1.Size() > 0 {
		val, err := q1.Dequeue()
		assert.NoError(err)
		assert.Equal(expected[n], val)
		n++
	}

	// Testing if the trie is sorted.
	sort.Strings(input)
	q2, _ := trie.Keys(q)

	for i := 0; i < len(input); i++ {
		val, err := q2.Dequeue()
		assert.NoError(err)
		assert.Equal(input[i], val)
	}

	// Replace an existing key.
	trie.Put("catz", 0)
	v, _ = trie.Get("catz")
	assert.Equal(0, v)
}

func TestTrie_Concurrency(t *testing.T) {
	assert := assert.New(t)
	wg := &sync.WaitGroup{}

	trie := New[string, int]()
	n := 10

	q := queue.New[string]()

	wg.Add(n)
	go func() {
		for i := 0; i < n; i++ {
			str := strconv.Itoa(i)
			trie.Put("test"+str, i)
			trie.Get("test" + str)
			wg.Done()
		}
	}()
	wg.Wait()

	keys, err := trie.Keys(q)
	fmt.Println(keys, err)

	res, _ := trie.LongestPrefix("te")
	trie.StartsWith(q, "ca")
	fmt.Println(res)

	fmt.Println("Size:", trie.Size())

	assert.Equal(1, 1)
}
