package gogu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert := assert.New(t)

	str1 := Substr("abcdef", 0, 0)
	str2 := Substr("abcdef", -1, 0)
	str3 := Substr("abcdef", 7, 7)
	str4 := Substr("abcdef", 0, 20)
	str5 := Substr("abcdef", 5, 10)
	str6 := Substr("abcdef", 0, -1)
	str7 := Substr("abcdef", 2, -1)
	str8 := Substr("abcdef", 4, -4)
	str9 := Substr("abcdef", -3, -1)
	str10 := Substr("abcdef", 1, 3)
	str11 := Substr("abcdef", 0, 4)
	str12 := Substr("abcdef", 0, 8)
	str13 := Substr("abcdef", -1, 1)
	str14 := Substr("abcdef", -2, -8)
	str15 := Substr("abcdef", -4, 1)

	assert.Empty(str1)
	assert.Empty(str2)
	assert.Empty(str3)
	assert.Equal("abcdef", str4)
	assert.Equal("f", str5)
	assert.Equal("abcde", str6)
	assert.Equal("cde", str7)
	assert.Empty(str8)
	assert.Equal("de", str9)
	assert.Equal("bcd", str10)
	assert.Equal("abcd", str11)
	assert.Equal("abcdef", str12)
	assert.Equal("f", str13)
	assert.Equal("", str14)
	assert.Equal("c", str15)
}