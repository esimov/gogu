package cache

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRUCache_Add(t *testing.T) {
	assert := assert.New(t)

	lruCache, _ := NewLRU[string, string](4)
	_, _, oldest := lruCache.Add("key1", "value1")
	assert.False(oldest)
	_, _, oldest = lruCache.Add("key2", "value2")
	assert.False(oldest)
	_, _, oldest = lruCache.Add("key3", "value3")
	assert.False(oldest)
	_, _, oldest = lruCache.Add("key4", "value4")
	assert.False(oldest)

	assert.Equal(4, lruCache.GetLength())

	oldestKey, oldestValue, removed := lruCache.Add("key5", "value5")
	assert.True(removed)
	assert.Equal("key1", oldestKey)
	assert.Equal("value1", oldestValue)

	oldestKey, oldestValue, removed = lruCache.Add("key6", "value6")
	assert.True(removed)
	assert.Equal("key2", oldestKey)
	assert.Equal("value2", oldestValue)

	assert.Equal(4, lruCache.GetLength())
}

func TestLRUCache_Get(t *testing.T) {
	assert := assert.New(t)

	lruCache, _ := NewLRU[string, string](4)
	lruCache.Add("key1", "value1")
	lruCache.Add("key2", "value2")
	lruCache.Add("key3", "value3")
	lruCache.Add("key4", "value4")

	assert.Equal(4, lruCache.GetLength())

	value, exists := lruCache.Get("key2")
	assert.True(exists)
	assert.Equal("value2", value)

	youngestKey, youngestValue, youngestExists := lruCache.GetYoungest()
	assert.True(youngestExists)
	assert.Equal("key2", youngestKey)
	assert.Equal("value2", youngestValue)

	oldestKey, oldestValue, oldestExists := lruCache.GetOldest()
	assert.True(oldestExists)
	assert.Equal("key1", oldestKey)
	assert.Equal("value1", oldestValue)

	oldestKeyUpdated, oldestValueUpdated, oldestExistsUpdated := lruCache.GetOldest()
	assert.True(oldestExistsUpdated)
	assert.Equal("key3", oldestKeyUpdated)
	assert.Equal("value3", oldestValueUpdated)
}

func TestLRUCache_Remove(t *testing.T) {
	assert := assert.New(t)

	lruCache, _ := NewLRU[string, string](4)
	lruCache.Add("key1", "value1")
	lruCache.Add("key2", "value2")
	lruCache.Add("key3", "value3")
	lruCache.Add("key4", "value4")

	assert.Equal(4, lruCache.GetLength())

	oldestKey, oldestValue, removed := lruCache.RemoveOldest()
	assert.True(removed)
	assert.Equal("key1", oldestKey)
	assert.Equal("value1", oldestValue)
	assert.Equal(3, lruCache.GetLength())

	oldestKey, oldestValue, removed = lruCache.RemoveOldest()
	assert.True(removed)
	assert.Equal("key2", oldestKey)
	assert.Equal("value2", oldestValue)
	assert.Equal(2, lruCache.GetLength())

	oldestKey, oldestValue, removed = lruCache.RemoveOldest()
	assert.True(removed)
	assert.Equal("key3", oldestKey)
	assert.Equal("value3", oldestValue)
	assert.Equal(1, lruCache.GetLength())

	oldestKey, oldestValue, removed = lruCache.RemoveOldest()
	assert.True(removed)
	assert.Equal("key4", oldestKey)
	assert.Equal("value4", oldestValue)
	assert.Equal(0, lruCache.GetLength())

	_, _, removed = lruCache.RemoveOldest()
	assert.False(removed)
	assert.Equal(0, lruCache.GetLength())

	lruCache.Add("key1", "value1")
	lruCache.Add("key2", "value2")
	lruCache.Add("key3", "value3")
	lruCache.Add("key4", "value4")

	value, r := lruCache.Remove("key2")
	assert.True(r)
	assert.Equal("value2", value)
	assert.Equal(3, lruCache.GetLength())

	value, r = lruCache.Remove("key1")
	assert.True(r)
	assert.Equal("value1", value)
	assert.Equal(2, lruCache.GetLength())

	value, r = lruCache.Remove("nonexistent")
	assert.False(r)
	assert.Equal(2, lruCache.GetLength())

	value, r = lruCache.Remove("key4")
	assert.True(r)
	assert.Equal("value4", value)
	assert.Equal(1, lruCache.GetLength())

	value, r = lruCache.Remove("key3")
	assert.True(r)
	assert.Equal("value3", value)
	assert.Equal(0, lruCache.GetLength())

	lruCache.Add("key1", "value1")
	lruCache.Add("key2", "value2")
	lruCache.Add("key3", "value3")

	oldestKey, oldestValue, removed = lruCache.RemoveOldest()
	assert.True(r)
	assert.Equal("key1", oldestKey)
	assert.Equal("value1", oldestValue)
	assert.Equal(2, lruCache.GetLength())

	youngestKey, youngestValue, r := lruCache.RemoveYoungest()
	assert.True(r)
	assert.Equal("key3", youngestKey)
	assert.Equal("value3", youngestValue)
	assert.Equal(1, lruCache.GetLength())
}

func TestLRUCache_GeneralUsage(t *testing.T) {
	assert := assert.New(t)

	lruCache, _ := NewLRU[string, string](5)
	lruCache.Add("key1", "value1")
	lruCache.Add("key2", "value2")
	lruCache.Add("key3", "value3")
	lruCache.Add("key4", "value4")
	lruCache.Add("key5", "value5")

	assert.Equal(5, lruCache.GetLength())

	lruCache.Add("key6", "value6")

	assert.Equal(5, lruCache.GetLength())

	value, exists := lruCache.Get("key2")
	assert.True(exists)
	assert.Equal("value2", value)

	youngestKey, youngestValue, youngestExists := lruCache.GetYoungest()
	assert.True(youngestExists)
	assert.Equal("key2", youngestKey)
	assert.Equal("value2", youngestValue)

	removedValue, removed := lruCache.Remove("key2")
	assert.True(removed)
	assert.Equal("value2", removedValue)
	assert.Equal(4, lruCache.GetLength())

	oldestKeyUpdated, oldestValueUpdated, oldestExistsUpdated := lruCache.GetOldest()
	assert.True(oldestExistsUpdated)
	assert.Equal("key3", oldestKeyUpdated)
	assert.Equal("value3", oldestValueUpdated)
}

func Example_lRUCache() {
	c, _ := NewLRU[string, string](3)
	item, available := c.Get("foo")
	fmt.Println(available)

	c.Add("foo", "bar")
	item, available = c.Get("foo")
	fmt.Println(available)
	fmt.Println(item)

	c.Add("foo2", "bar2")
	c.Add("foo3", "bar3")
	c.Add("foo4", "baz")

	fmt.Println(c.GetLength())
	fmt.Println()

	oldestKey, oldestValue, oldestAvailable := c.GetOldest()
	fmt.Println(oldestAvailable)
	fmt.Println(oldestKey)
	fmt.Println(oldestValue)
	fmt.Println()

	youngestKey, youngestValue, youngestAvailable := c.GetYoungest()
	fmt.Println(youngestAvailable)
	fmt.Println(youngestKey)
	fmt.Println(youngestValue)
	fmt.Println()

	oldestKey, oldestValue, oldestAvailable = c.RemoveOldest()
	fmt.Println(oldestAvailable)
	fmt.Println(oldestKey)
	fmt.Println(oldestValue)
	fmt.Println()

	// Output:
	// false
	// true
	// bar
	// 3
	//
	// true
	// foo2
	// bar2
	//
	// true
	// foo2
	// bar2
	//
	// true
	// foo3
	// bar3
}
