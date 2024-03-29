package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	Items []*SampleStruct
	Id    int
}

func TestCache_Basic(t *testing.T) {
	assert := assert.New(t)

	c1 := New[string, string](DefaultExpiration, 1*time.Minute)
	res1, err := c1.Get("foo")
	assert.Error(err)
	assert.Nil(res1)

	res1, err = c1.Get("bar")
	assert.Error(err)
	assert.Nil(res1)

	err = c1.Set("foo", "bar", DefaultExpiration)
	assert.NoError(err)
	err = c1.Set("foo", "", DefaultExpiration)
	assert.Error(err)

	res1, err = c1.Get("foo")
	assert.NoError(err)
	assert.NotEmpty(res1.Val())
	assert.Equal("bar", res1.Val())
	assert.NotEqual("baz", res1.Val())

	assert.False(c1.IsExpired("foo"))
	assert.False(c1.IsExpired("baz"))

	c2 := New[string, int](DefaultExpiration, 1*time.Minute)
	err = c2.Set("foo", 1, DefaultExpiration)
	assert.NoError(err)

	res2, err := c2.Get("foo")
	assert.NoError(err)
	assert.NotEmpty(res2.Val())
	assert.Equal(1, res2.Val())

	err = c2.SetDefault("bar", 2)
	assert.NoError(err)
	res2, _ = c2.Get("bar")
	assert.Equal(int64(DefaultExpiration), res2.expiration)

	res2, _ = c2.Get("bar")
	list := c2.List()
	assert.Equal(res2.Val(), list["bar"].object)
	assert.Len(list, 2)

	c2.Flush()
	assert.Equal(0, c2.Count())
	assert.Len(c2.List(), 0)

	items := make(map[string]int)
	items["item1"] = 1
	items["item2"] = 2
	err = c2.MapToCache(items, DefaultExpiration)
	assert.Equal(2, c2.Count())
	assert.NoError(err)
	res3, _ := c2.Get("item1")
	res4, _ := c2.Get("item2")
	assert.Equal(1, res3.Val())
	assert.Equal(2, res4.Val())
}

func ExampleCache() {
	c := New[string, string](DefaultExpiration, 1*time.Minute)
	item, err := c.Get("foo")
	fmt.Println(err)
	fmt.Println(item)

	c.Set("foo", "bar", DefaultExpiration)
	item, _ = c.Get("foo")
	fmt.Println(item.Val())

	err = c.Set("foo", "", DefaultExpiration)
	fmt.Println(err)
	fmt.Println(c.IsExpired("foo"))

	c.Update("foo", "baz", DefaultExpiration)
	item, _ = c.Get("foo")
	fmt.Println(item.Val())

	list := c.List()
	fmt.Println(len(list))

	c.Flush()
	fmt.Println(c.Count())

	c.Set("foo", "bar", DefaultExpiration)
	item, _ = c.Get("foo")
	fmt.Println(item.Val())

	err = c.Delete("foo")
	fmt.Println(err)
	fmt.Println(c.Count())

	// Output:
	// item with key 'foo' not found
	// <nil>
	// bar
	// item with key 'foo' already exists. Use the Update method
	// false
	// baz
	// 1
	// 0
	// bar
	// <nil>
	// 0
}

func TestCache_PointerStruct(t *testing.T) {
	assert := assert.New(t)

	st := &SampleStruct{Id: 1}

	c := New[string, SampleStruct](DefaultExpiration, 1*time.Minute)
	c.Set("foo", *st, DefaultExpiration)

	x, err := c.Get("foo")
	assert.NoError(err)
	assert.Equal(1, x.Val().Id)

	x.object.Id++

	y, err := c.Get("foo")
	assert.NoError(err)
	assert.Equal(2, y.Val().Id)
}

func TestCache_Update(t *testing.T) {
	assert := assert.New(t)

	c1 := New[string, string](DefaultExpiration, 1*time.Minute)
	err := c1.Set("item1", "a", DefaultExpiration)
	assert.NoError(err)

	res1, err := c1.Get("item1")
	assert.NotNil(res1)
	assert.Equal("a", res1.Val())
	assert.NoError(err)

	err = c1.Set("item1", "b", DefaultExpiration)
	assert.Error(err)
	err = c1.Update("item1", "", DefaultExpiration)
	assert.Error(err)
	err = c1.Update("item1", "c", DefaultExpiration)
	assert.NoError(err)
	res1, _ = c1.Get("item1")
	assert.Equal("c", res1.Val())
	c1.Update("item1", "d", NoExpiration)
}

func TestCache_Delete(t *testing.T) {
	assert := assert.New(t)

	c1 := New[string, string](DefaultExpiration, 0)
	c1.SetDefault("item1", "a")
	c1.SetDefault("item2", "a")
	c1.SetDefault("item3", "a")
	c1.SetDefault("item4", "a")

	assert.Len(c1.List(), 4)
	err := c1.Delete("item1")
	assert.NoError(err)
	assert.Len(c1.List(), 3)
	res, err := c1.Get("item1")
	assert.Nil(res)
	assert.Error(err)
	err = c1.Delete("item1")
	assert.Error(err)

	err = c1.Delete("item5")
	assert.Error(err)
	c1.Delete("item2")
	assert.Len(c1.List(), 2)

	c1.Delete("item4")
	assert.Len(c1.List(), 1)

	res, err = c1.Get("item3")
	assert.NoError(err)
	assert.Equal("a", res.Val())
	c1.Delete("item3")
	assert.Empty(c1.List())
}

func TestCache_ExpirationTime(t *testing.T) {
	assert := assert.New(t)

	c1 := New[string, string](NoExpiration, 0)
	c1.Set("item1", "a", DefaultExpiration)
	res, _ := c1.Get("item1")
	assert.Equal(int64(NoExpiration), res.expiration)

	c1.Update("item1", "b", NoExpiration)
	res, _ = c1.Get("item1")
	assert.Equal(int64(NoExpiration), res.expiration)

	c1.Set("item1", "a", 10*time.Millisecond)
	c1.Delete("item1")
	assert.Empty(c1.List())

	err := c1.DeleteExpired()
	assert.Nil(err)

	c1.Set("item1", "a", 1*time.Millisecond)
	c1.Set("item2", "a", 1*time.Millisecond)
	<-time.After(2 * time.Millisecond)
	err = c1.DeleteExpired()
	assert.NoError(err)
	assert.Len(c1.List(), 0)

	c1.Set("item1", "b", 10*time.Millisecond)
	c1.Set("item2", "b", 60*time.Millisecond)
	<-time.After(40 * time.Millisecond)
	c1.DeleteExpired()
	assert.Len(c1.List(), 1)
	<-time.After(100 * time.Millisecond)
	c1.DeleteExpired()
	assert.Empty(c1.List())

	c1.Set("item1", "c", 10*time.Millisecond)
	<-time.After(20 * time.Millisecond)
	res, err = c1.Get("item1")
	assert.Nil(res)
	assert.Error(err)

	c2 := New[string, int](5*time.Millisecond, 100*time.Millisecond)
	c2.Set("a", 1, DefaultExpiration)
	c2.Set("b", 2, NoExpiration)
	c2.Set("c", 3, 50*time.Millisecond)
	c2.Set("d", 4, 200*time.Millisecond)
	<-time.After(150 * time.Millisecond)
	assert.Equal(2, c2.Count())
	<-time.After(300 * time.Millisecond)
	assert.Equal(1, c2.Count())

	c2.Set("b", 1, 2*time.Millisecond)
	<-time.After(40 * time.Millisecond)
	assert.Equal(1, c2.Count())
}

func Example_expirationTime() {
	c1 := New[string, string](NoExpiration, 0)
	c1.Set("item1", "a", DefaultExpiration)
	item, _ := c1.Get("item1")
	fmt.Println(item.expiration)

	c1.Update("item1", "b", NoExpiration)
	item, _ = c1.Get("item1")
	fmt.Println(item.expiration)

	err := c1.DeleteExpired()
	fmt.Println(err)

	c1.Set("item1", "a", 10*time.Millisecond)
	<-time.After(20 * time.Millisecond)
	c1.DeleteExpired()
	fmt.Println(c1.Count())

	c1.Set("item1", "b", 1*time.Millisecond)
	c1.Set("item2", "b", 50*time.Millisecond)
	<-time.After(20 * time.Millisecond)
	c1.DeleteExpired()
	fmt.Println(c1.Count())

	<-time.After(70 * time.Millisecond)
	c1.DeleteExpired()
	fmt.Println(c1.Count())

	c2 := New[string, int](5*time.Millisecond, 100*time.Millisecond)
	c2.Set("a", 1, DefaultExpiration)
	c2.Set("b", 2, NoExpiration)
	c2.Set("c", 3, 50*time.Millisecond)
	c2.Set("d", 4, 200*time.Millisecond)
	<-time.After(150 * time.Millisecond)
	fmt.Println(c2.Count())
	<-time.After(300 * time.Millisecond)
	fmt.Println(c2.Count())

	// Output:
	// -1
	// -1
	// <nil>
	// 1
	// 2
	// 1
	// 2
	// 1
}
