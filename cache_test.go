package gogu

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	Id    int
	Items []*SampleStruct
}

func TestCache_Basic(t *testing.T) {
	assert := assert.New(t)

	c1 := NewCache[string, string](DefaultExpiration, 1*time.Minute)
	res1, err := c1.Get("foo")
	assert.Error(err)
	assert.Nil(res1)

	res2, err := c1.Get("bar")
	assert.Error(err)
	assert.Nil(res2)

	err = c1.Set("foo", "bar", DefaultExpiration)
	assert.NoError(err)
	err = c1.Set("foo", "", DefaultExpiration)
	assert.Error(err)

	res3, err := c1.Get("foo")
	assert.NoError(err)
	assert.NotEmpty(res3.Val())
	assert.Equal("bar", res3.Val())
	assert.NotEqual("baz", res3.Val())

	assert.False(c1.IsExpired("foo"))
	assert.False(c1.IsExpired("baz"))

	c2 := NewCache[string, int](DefaultExpiration, 1*time.Minute)
	err = c2.Set("foo", 1, DefaultExpiration)
	assert.NoError(err)

	r4, err := c2.Get("foo")
	assert.NoError(err)
	assert.NotEmpty(r4.Val())
	assert.Equal(1, r4.Val())

	err = c2.SetDefault("bar", 2)
	assert.NoError(err)
	res4, _ := c2.Get("bar")
	assert.Equal(int64(DefaultExpiration), res4.expiration)
}

func TestCache_PointerStruct(t *testing.T) {
	assert := assert.New(t)
	st := &SampleStruct{Id: 1}

	c := NewCache[string, SampleStruct](DefaultExpiration, 1*time.Minute)
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

	c1 := NewCache[string, string](DefaultExpiration, 1*time.Minute)
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

	c1 := NewCache[string, string](DefaultExpiration, 0)
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

	c1 := NewCache[string, string](NoExpiration, 0)
	c1.Set("item1", "a", DefaultExpiration)
	res, _ := c1.Get("item1")
	assert.Equal(int64(DefaultExpiration), res.expiration)

	c1.Update("item1", "b", NoExpiration)
	res, _ = c1.Get("item1")
	assert.Equal(int64(DefaultExpiration), res.expiration)

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

	c1.Set("item1", "b", 1*time.Millisecond)
	c1.Set("item2", "b", 3*time.Millisecond)
	<-time.After(2 * time.Millisecond)
	c1.DeleteExpired()
	assert.Len(c1.List(), 1)
	<-time.After(1 * time.Millisecond)
	c1.DeleteExpired()
	assert.Empty(c1.List())

	fmt.Println()
}
