package gogu

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_Get(t *testing.T) {
	assert := assert.New(t)
	c := NewCache[string, string](DefaultExpiration, 1*time.Minute)
	assert.Nil(c.Get("foo"))
}
