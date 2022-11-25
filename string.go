package torx

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

func Null[T any]() T {
	var t T
	return t
}

// Substr returns the portion of string specified by the offset and length.
//
// If offset is non-negative, the returned string will start
// at the offset'th position in string, counting from zero.
//
// If offset is negative, the returned string will start at
// the offset'th character from the end of string.
//
// If string is less than offset characters long, an empty string will be returned.
//
// If length is negative, then that many characters will be omitted
// from the end of string starting from the offset position.
func Substr[T ~string](str T, offset, length int) T {
	var end int

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

// ToLower converts a string to Lowercase.
func ToLower[T ~string](str T) T {
	result := make([]rune, 0, len(str))

	for _, val := range str {
		result = append(result, unicode.ToLower(rune(val)))
	}

	return T(result)
}

// ToUpper converts a string to Uppercase.
func ToUpper[T ~string](str T) T {
	result := make([]rune, 0, len(str))

	for _, val := range str {
		result = append(result, unicode.ToLower(rune(val)))
	}

	return T(result)
}

// Capitalize converts the first letter of the string
// to uppercase and the remaining letters to lowercase.
func Capitalize[T ~string](str T) T {
	result := make([]rune, 0, len(str))

	for i, val := range str {
		if i == 0 {
			result = append(result, unicode.ToUpper(rune(val)))
		} else {
			result = append(result, unicode.ToLower(rune(val)))
		}
	}

	return T(result)
}

// CamelCase converts a string to camelCase (https://en.wikipedia.org/wiki/CamelCase).
func CamelCase[T ~string](str T) T {
	newstr := strings.TrimSpace(string(str))

	r, _ := regexp.Compile("[-_&]+")
	newstr = r.ReplaceAllString(newstr, " ")

	var sb strings.Builder
	sb.Grow(len(newstr))

	var idx int
	for i, s := range strings.Split(newstr, " ") {
		r := []rune(s)

		if len(r) == 0 {
			idx++
			continue
		}

		if i == 0 || i == idx {
			frag := ToLower(s)
			sb.WriteString(frag)
			continue
		}
		sb.WriteString(Capitalize(s))
	}
	result := sb.String()

	return T(result)
}

// SnakeCase converts a string to snake_case (https://en.wikipedia.org/wiki/Snake_case).
func SnakeCase[T ~string](str T) T {
	return splitStringWithDelimiter(str, "_")
}

// KebabCase converts a string to kebab-case (https://en.wikipedia.org/wiki/Letter_case#Kebab_case).
func KebabCase[T ~string](str T) T {
	return splitStringWithDelimiter(str, "-")
}

// splitStringWithDelimiter splits a string to lower case with the provided delimiter.
func splitStringWithDelimiter[T ~string](str T, delimiter string) T {
	var sb strings.Builder
	newstr := strings.TrimSpace(string(str))

	rx, _ := regexp.Compile("[-_&]+")
	newstr = rx.ReplaceAllString(newstr, " ")

	var idx int
	chars := strings.Split(newstr, " ")
	for i, str := range chars {
		r := []rune(str)

		if len(r) == 0 {
			idx++
			continue
		}
		rx, _ = regexp.Compile("[a-zö][A-ZÖ]+")
		strIdx := rx.FindAllStringIndex(str, -1)

		if len(strIdx) > 0 {
			s := Substr(str, 0, strIdx[0][0]+1)
			sb.WriteString(ToLower(s))
			sb.WriteString(delimiter)

			for i := 0; i < len(strIdx); i++ {
				var s string
				if i < len(strIdx)-1 {
					subStrLen := strIdx[i+1][0] - strIdx[i][0]
					s = Substr(str, strIdx[i][0]+1, subStrLen)

					sb.WriteString(ToLower(s))
					sb.WriteString(delimiter)
				} else {
					s = Substr(str, strIdx[i][0]+1, len(str)-strIdx[i][0]+1)
					sb.WriteString(ToLower(s))
				}
			}

			if len(chars) > 1 && i != len(chars)-1 {
				sb.WriteString(delimiter)
			}
		} else {
			frag := ToLower(str)
			sb.WriteString(frag)

			if len(chars) > 1 && i != len(chars)-1 {
				sb.WriteString(delimiter)
			}
		}
	}
	result := sb.String()

	return T(result)
}

// PadLeft pads string on the left side if it's shorter than length.
// Padding characters are truncated if they exceed length.
func PadLeft[T ~string](str T, size int, token string) T {
	var tokenStr = token

	strLen := len(str)
	tokenLen := len(token)

	if size <= strLen {
		return T(str)
	}

	if tokenLen <= size-strLen {
		tokenStr = strings.Repeat(token, size-strLen)
	}

	tokenStr = tokenStr[:size-strLen]

	return T(tokenStr) + T(str)
}

// PadRight pads string on the right side if it's shorter than length.
// Padding characters are truncated if they exceed length.
func PadRight[T ~string](str T, size int, token string) T {
	var tokenStr = token

	strLen := len(str)
	tokenLen := len(token)

	if size <= strLen {
		return T(str)
	}

	if tokenLen <= size-strLen {
		tokenStr = strings.Repeat(token, size-strLen)
	}

	tokenStr = tokenStr[:size-strLen]

	return T(str) + T(tokenStr)
}

// Pads string on the left and right sides if it's shorter than length.
// Padding characters are truncated if they can't be evenly divided by length.
func Pad[T ~string](str T, size int, token string) T {
	var (
		leftTokenStr  = token
		rightTokenStr = token
	)

	strLen := len(str)
	tokenLen := len(token)

	if size <= strLen {
		return T(str)
	}
	split := float64(size-strLen) / 2
	left := int(math.Floor(split))
	right := int(math.Ceil(split))

	if tokenLen <= int(split) {
		leftTokenStr = strings.Repeat(token, left)
		rightTokenStr = strings.Repeat(token, right)
	}

	leftTokenStr = leftTokenStr[:left]
	rightTokenStr = rightTokenStr[:right]

	return T(leftTokenStr) + str + T(rightTokenStr)
}

// SplitAtIndex split the string at the specified index and
// returns a slice with the resulted two substrings.
func SplitAtIndex[T ~string](str T, index int) []T {
	result := make([]T, 0, 2)

	if index < 0 {
		return []T{"", str}
	}

	if index > len(str)-1 {
		return []T{str, ""}
	}

	for idx := range str {
		if idx == index {
			result = append(result, append(result, str[:idx+1], str[idx+1:])...)
		}
	}

	return result
}

// Wrap a string with the specified token.
func Wrap[T ~string](str T, token string) T {
	var s strings.Builder

	s.WriteString(token)
	s.WriteString(string(str))
	s.WriteString(token)

	return T(s.String())
}

// Unwrap a string with the specified token.
func Unwrap[T ~string](str T, token string) T {
	startToken := strings.Index(string(str), token)
	endToken := strings.LastIndex(string(str), token)

	if startToken == 0 && endToken <= len(str)-1 {
		str = str[len(token):endToken]
	}

	return str
}

// WrapAllRune is like Wrap, only that it's applied over runes instead of strings.
func WrapAllRune[T ~string](str T, token string) T {
	var s strings.Builder

	for _, st := range str {
		s.WriteString(token)
		s.WriteRune(st)
		s.WriteString(token)
	}

	return T(s.String())
}

// ReverseStr returns a new string with the characters in reverse order.
func ReverseStr[T ~string](str T) T {
	res := []rune(str)

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return T(res)
}
