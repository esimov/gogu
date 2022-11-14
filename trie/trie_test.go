package trie

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"testing"

	"github.com/esimov/torx/queue"
	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	assert := assert.New(t)

	q := queue.New[string]()
	trie := New[string, int](q)
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

	_, err = trie.StartsWith("")
	assert.Error(err)

	q1, err := trie.StartsWith("ca")
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
	q2, _ := trie.Keys()

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

	q := queue.New[string]()
	trie := New[string, int](q)
	n := 100

	wg.Add(n)
	go func() {
		for i := 0; i < n; i++ {
			str := strconv.Itoa(i)
			trie.Put(str, i)
			wg.Done()
		}
	}()
	wg.Wait()

	assert.Equal(100, trie.Size())
	keys, err := trie.Keys()
	assert.NoError(err)
	assert.Equal(100, keys.Size())

	qs, err := trie.StartsWith("")
	assert.Error(err)
	assert.Equal(0, qs.Size())

	qs2, err := trie.StartsWith("2")
	assert.NoError(err)
	assert.Equal(11, qs2.Size())
}

func Example() {
	q := queue.New[string]()
	trie := New[string, int](q)
	input := []string{"cats", "cape", "captain", "foes",
		"apple", "she", "root", "shells", "the", "thermos", "foo"}

	for idx, v := range input {
		trie.Put(v, idx)
	}

	longestPref, _ := trie.LongestPrefix("capetown")
	q1, _ := trie.StartsWith("ca")

	result := []string{}
	for q1.Size() > 0 {
		val, _ := q1.Dequeue()
		result = append(result, val)
	}

	fmt.Println(trie.Size())
	fmt.Println(longestPref)
	fmt.Println(result)

	// Output:
	// 11
	// cape
	// [cape captain cats]
}
