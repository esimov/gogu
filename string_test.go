package torx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Substr(t *testing.T) {
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
	str16 := Substr("abcdef", -10, -10)

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
	assert.Equal("", str16)
}

func Example_substr() {
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

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)
	fmt.Println(str6)
	fmt.Println(str7)
	fmt.Println(str8)
	fmt.Println(str9)
	fmt.Println(str10)

	// Output:
	// abcdef
	// f
	// abcde
	// cde
	//
	// de
	// bcd
}

func TestString_Capitalize(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Abc", Capitalize("Abc"))
	assert.Equal("Abc", Capitalize("abc"))
	assert.Equal("Abc", Capitalize("abC"))
	assert.Equal("Abc", Capitalize("aBC"))
	assert.Equal("Abc", Capitalize("ABC"))
	assert.Equal("Abø", Capitalize("aBø"))
	assert.Equal("", Capitalize(""))
}

func TestString_CamelCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("fooBar", CamelCase("Foo Bar"))
	assert.Equal("fooBar", CamelCase("--foo-Bar--"))
	assert.Equal("fooBar", CamelCase("__foo-_Bar__"))
	assert.Equal("fooBar", CamelCase("__FOO BAR__"))
	assert.Equal("fooBar", CamelCase(" FOO BAR "))
	assert.Equal("fooBar", CamelCase("&FOO&baR "))
	assert.Equal("fooBar", CamelCase("&&foo&&bar__"))
}

func Example_camelCase() {
	fmt.Println(CamelCase("Foo Bar"))
	fmt.Println(CamelCase("--foo-Bar--"))
	fmt.Println(CamelCase("__foo-_Bar__"))
	fmt.Println(CamelCase("__FOO BAR__"))
	fmt.Println(CamelCase(" FOO BAR "))
	fmt.Println(CamelCase("&FOO&baR "))
	fmt.Println(CamelCase("&&foo&&bar__"))

	// Output:
	// fooBar
	// fooBar
	// fooBar
	// fooBar
	// fooBar
	// fooBar
	// fooBar
}

func TestString_SnakeCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("f", SnakeCase("F"))
	assert.Equal("foo", SnakeCase("Foo"))
	assert.Equal("foo_bar", SnakeCase("Foo Bar"))
	assert.Equal("foo_bar", SnakeCase("Foo_Bar"))
	assert.Equal("foo_bar", SnakeCase("fooBar"))
	assert.Equal("foo_bar", SnakeCase("foo_&Bar"))
	assert.Equal("foo_bar", SnakeCase("&Foo_&Bar"))
	assert.Equal("foo_bar", SnakeCase(" Foo_Bar "))
	assert.Equal("foo_bar", SnakeCase("foo__Bar"))
	assert.Equal("foo_bar_baz", SnakeCase("fooBarBaz"))
	assert.Equal("foo_bar_baz", SnakeCase("Foo BarBaz"))
	assert.Equal("foo_bar_baz", SnakeCase("Foo_Bar_Baz"))
	assert.Equal("foo_bar_baz_qux", SnakeCase("FooBarBazQux"))
	assert.Equal("foo_bar_baz_qux", SnakeCase("FooBarBaz_Qux"))
	assert.Equal("foo_bar_baz_qux", SnakeCase("FooBarBaz Qux"))
	assert.Equal("foo_bar_baz_qux", SnakeCase("Foo Bar_Baz&Qux"))
}

func Example_snakeCase() {
	fmt.Println(SnakeCase("fooBarBaz"))
	fmt.Println(SnakeCase("Foo BarBaz"))
	fmt.Println(SnakeCase("Foo_Bar_Baz"))

	// Output:
	// foo_bar_baz
	// foo_bar_baz
	// foo_bar_baz
}
func TestString_KebabCase(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("f", KebabCase("F"))
	assert.Equal("foo", KebabCase("Foo"))
	assert.Equal("foo-bar", KebabCase("Foo Bar"))
	assert.Equal("foo-bar", KebabCase("Foo_Bar"))
	assert.Equal("foo-bar", KebabCase("fooBar"))
	assert.Equal("foo-bar", KebabCase("foo_&Bar"))
	assert.Equal("foo-bar", KebabCase("&Foo_&Bar"))
	assert.Equal("foo-bar", KebabCase(" Foo_Bar "))
	assert.Equal("foo-bar", KebabCase("foo__Bar"))
	assert.Equal("foo-bar-baz", KebabCase("fooBarBaz"))
	assert.Equal("foo-bar-baz", KebabCase("Foo BarBaz"))
	assert.Equal("foo-bar-baz", KebabCase("Foo_Bar_Baz"))
	assert.Equal("foo-bar-baz-qux", KebabCase("FooBarBazQux"))
	assert.Equal("foo-bar-baz-qux", KebabCase("FooBarBaz_Qux"))
	assert.Equal("foo-bar-baz-qux", KebabCase("FooBarBaz Qux"))
	assert.Equal("foo-bar-baz-qux", KebabCase("Foo Bar_Baz&Qux"))
}

func Example_kebebCase() {
	fmt.Println(KebabCase("fooBarBaz"))
	fmt.Println(KebabCase("Foo BarBaz"))
	fmt.Println(KebabCase("Foo_Bar_Baz"))

	// Output:
	// foo-bar-baz
	// foo-bar-baz
	// foo-bar-baz
}

func TestString_SplitAtIndex(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"", "abcdef"}, SplitAtIndex("abcdef", -1))
	assert.Equal([]string{"a", "bcdef"}, SplitAtIndex("abcdef", 0))
	assert.Equal([]string{"ab", "cdef"}, SplitAtIndex("abcdef", 1))
	assert.Equal([]string{"abc", "def"}, SplitAtIndex("abcdef", 2))
	assert.Equal([]string{"abcdef", ""}, SplitAtIndex("abcdef", 5))
	assert.Equal([]string{"abcdef", ""}, SplitAtIndex("abcdef", 6))
}

func Example_splitAtIndex() {
	fmt.Println(SplitAtIndex("abcdef", -1))
	fmt.Println(SplitAtIndex("abcdef", 0))
	fmt.Println(SplitAtIndex("abcdef", 1))
	fmt.Println(SplitAtIndex("abcdef", 2))
	fmt.Println(SplitAtIndex("abcdef", 5))
	fmt.Println(SplitAtIndex("abcdef", 6))

	// Output:
	// [ abcdef]
	// [a bcdef]
	// [ab cdef]
	// [abc def]
	// [abcdef ]
	// [abcdef ]
}

func TestString_Pad(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", PadLeft("abc", 2, "."))
	assert.Equal("abc", PadLeft("abc", 3, "."))
	assert.Equal("_abc", PadLeft("abc", 4, "_"))
	assert.Equal(".....abc", PadLeft("abc", 8, "..."))
	assert.Equal("...abc", PadLeft("abc", 6, "........"))
	assert.Equal("_-_abc", PadLeft("abc", 6, "_-"))
	assert.Equal("_-|_abc", PadLeft("abc", 7, "_-|"))

	assert.Equal("abc", PadRight("abc", 2, "."))
	assert.Equal("abc", PadRight("abc", 3, "."))
	assert.Equal("abc_", PadRight("abc", 4, "_"))
	assert.Equal("abc.....", PadRight("abc", 8, "..."))
	assert.Equal("abc...", PadRight("abc", 6, "........"))
	assert.Equal("abc_-_", PadRight("abc", 6, "_-"))
	assert.Equal("abc_-|_", PadRight("abc", 7, "_-|"))

	assert.Equal("abc", Pad("abc", 2, "."))
	assert.Equal("abc", Pad("abc", 3, "."))
	assert.Equal("abc.", Pad("abc", 4, "."))
	assert.Equal(".abc.", Pad("abc", 5, "."))
	assert.Equal(".abc..", Pad("abc", 6, "."))
	assert.Equal("  abc  ", Pad("abc", 7, " "))
	assert.Equal("_-abc_-_", Pad("abc", 8, "_-"))
}

func Example_pad() {
	fmt.Println(Pad("abc", 2, "."))
	fmt.Println(Pad("abc", 3, "."))
	fmt.Println(Pad("abc", 4, "."))
	fmt.Println(Pad("abc", 5, "."))

	fmt.Println(PadRight("abc", 8, "..."))
	fmt.Println(PadRight("abc", 6, "........"))

	fmt.Println(PadLeft("abc", 8, "..."))
	fmt.Println(PadLeft("abc", 4, "_"))
	fmt.Println(PadLeft("abc", 6, "_-"))

	// Output:
	// abc
	// abc
	// abc.
	// .abc.
	// abc.....
	// abc...
	// .....abc
	// _abc
	// _-_abc
}

func TestString_ReverseStr(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("fedcba", ReverseStr("abcdef"))
	assert.Equal("FEDCBA", ReverseStr("ABCDEF"))
	assert.Equal("654321", ReverseStr("123456"))
}

func TestString_Wrap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", Wrap("abc", ""))
	assert.Equal("'abc'", Wrap("abc", "'"))
	assert.Equal("*abc*", Wrap("abc", "*"))
	assert.Equal(`\abc\`, Wrap("abc", `\`))
	assert.Equal(`|abc|`, Wrap("abc", `|`))
}

func TestString_Unwrap(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", Unwrap("'abc'", "'"))
	assert.Equal("abc", Unwrap("*abc*", "*"))
	assert.Equal("a*bc", Unwrap("*a*bc*", "*"))
	assert.Equal("abc", Unwrap("''abc''", "''"))
	assert.Equal("abc", Unwrap("\"abc\"", "\""))
	assert.Equal("'", Unwrap("'''", "'"))
}

func Example_unwrap() {
	fmt.Println(Unwrap("'abc'", "'"))
	fmt.Println(Unwrap("*abc*", "*"))
	fmt.Println(Unwrap("*a*bc*", "*"))
	fmt.Println(Unwrap("''abc''", "''"))
	fmt.Println(Unwrap("\"abc\"", "\""))

	// Output:
	// abc
	// abc
	// a*bc
	// abc
	// abc
}

func TestString_WrapAllRune(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("abc", WrapAllRune("abc", ""))
	assert.Equal("'a''b''c'", WrapAllRune("abc", "'"))
	assert.Equal("*a**b**c*", WrapAllRune("abc", "*"))
	assert.Equal("-a--b--c-", WrapAllRune("abc", "-"))
}

func Example_wrapAllRune() {
	fmt.Println(WrapAllRune("abc", ""))
	fmt.Println(WrapAllRune("abc", "'"))
	fmt.Println(WrapAllRune("abc", "*"))
	fmt.Println(WrapAllRune("abc", "-"))

	// Output:
	// abc
	// 'a''b''c'
	// *a**b**c*
	// -a--b--c-
}
