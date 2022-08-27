package gogu

func Empty[T any]() T {
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
			return Empty[T]()
		}
	}
	if length < 0 {
		newLength := len(str) + length
		if Abs(newLength) > len(str) || newLength < offset {
			return Empty[T]()
		}
		end = newLength
	} else {
		end = offset + length
	}

	if end > len(str) {
		end = len(str)
	}

	if !InRange(offset, 0, len(str)) || !InRange(end, 0, len(str)) {
		return Empty[T]()
	}

	return str[offset:end]
}
