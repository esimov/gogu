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

func TestCache(t *testing.T) {
	assert := assert.New(t)

	c1 := NewCache[string, string](DefaultExpiration, 1*time.Minute)
	r1, err := c1.Get("foo")
	assert.Error(err)
	assert.Nil(r1)

	r2, err := c1.Get("bar")
	assert.Error(err)
	assert.Nil(r2)

	err = c1.Set("foo", "bar", DefaultExpiration)
	assert.NoError(err)
	err = c1.Set("foo", "", DefaultExpiration)
	assert.Error(err)

	r3, err := c1.Get("foo")
	assert.NoError(err)
	assert.NotEmpty(r3.Val())
	assert.Equal("bar", r3.Val())
	assert.NotEqual("baz", r3.Val())

	assert.False(c1.IsExpired("foo"))
	assert.False(c1.IsExpired("baz"))

	c2 := NewCache[string, int](DefaultExpiration, 1*time.Minute)
	err = c2.Set("foo", 1, DefaultExpiration)
	assert.NoError(err)

	r4, err := c2.Get("foo")
	assert.NoError(err)
	assert.NotEmpty(r4.Val())
	assert.Equal(1, r4.Val())
}

func TestCache_TestPointerStruct(t *testing.T) {
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
	fmt.Println()
	//assert := assert.New(t)

}
