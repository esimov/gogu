package gogu

import (
	"unicode"
)

func Null[T any]() T {
	var t T
	return t
}

// Substr returns the portion of string specified by the offset and length.

// If offset is non-negative, the returned string will start
// at the offset'th position in string, counting from zero.

// If offset is negative, the returned string will start at
// the offset'th character from the end of string.

// If string is less than offset characters long, an empty string will be returned.

// If length is negative, then that many characters will be omitted
// from the end of string starting from the offset position.
func Substr[T ~string](str T, offset, length int) T {
	var end = length

	if offset < 0 {
		offset = len(str) + offset
		if Abs(offset) > len(str) {
			return Null[T]()
		}
	}
	if length < 0 {
		newLength := len(str) + length
		if Abs(newLength) > len(str) || newLength < offset {
			return Null[T]()
		}
		end = newLength
	} else {
		end = offset + length
	}

	if end > len(str) {
		end = len(str)
	}

	if !InRange(offset, 0, len(str)) || !InRange(end, 0, len(str)) {
		return Null[T]()
	}

	return str[offset:end]
}

// Capitalize converts the first letter of the string
// to uppercase and the remaining letters to lowercase.
func Capitalize[T ~string](str T) T {
	res := make([]rune, 0, len(str))

	for i, val := range str {
		if i == 0 {
			res = append(res, unicode.ToUpper(rune(val)))
		} else {
			res = append(res, unicode.ToLower(rune(val)))
		}
	}

	return T(res)
}
