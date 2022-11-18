# torx

Torx is a versatile, comprehensive, reusable and efficient utility functions and data structures library taking advantage of the Go generics. It was inspired by other well established and consecrated frameworks like [underscore.js](https://underscorejs.org/), [lodash](https://lodash.com/) and some concepts familiar to the functional programming paradigms. 

Its main purpose is to help the developers in their day-to-day jobs by exposing some generic helper functions to facilitate the work with slices, maps and strings, but also implementing some of the most used data structures.

## Installation

```bash
$ go install github.com/esimov/torx@latest
```

## Usage
```go
package main

import "github.com/esimov/torx"

func main() {
  // main program
}
```

## Specifications
- General utility functions
  - [Abs](<#func-abs>)
  - [Clamp](<#func-clamp>)
  - [Compare](<#func-compare>)
  - [Equal](<#func-equal>)
  - [InRange](<#func-inrange>)
  - [Invert](<#func-invert>)
  - [Less](<#func-less>)
  - [Max](<#func-max>)
  - [Min](<#func-min>)

- Strings utility functions
  - [CamelCase](<#func-camelcase>)
  - [Capitalize](<#func-capitalize>)
  - [KebabCase](<#func-kebabcase>)
  - [N](<#func-n>)
  - [Null](<#func-null>)
  - [NumToString](<#func-numtostring>)
  - [Pad](<#func-pad>)
  - [PadLeft](<#func-padleft>)
  - [PadRight](<#func-padright>)
  - [ReverseStr](<#func-reversestr>)
  - [SnakeCase](<#func-snakecase>)
  - [SplitAtIndex](<#func-splitatindex>)
  - [Substr](<#func-substr>)
  - [ToLower](<#func-tolower>)
  - [ToUpper](<#func-toupper>)
  - [Unwrap](<#func-unwrap>)
  - [Wrap](<#func-wrap>)
  - [WrapAllRune](<#func-wrapallrune>)

- Slice utility functions
  - [Chunk](<#func-chunk>)
  - [SumBy](#sumby)
  - [Contains](<#func-contains>)
  - [Difference](<#func-difference>)
  - [DifferenceBy](<#func-differenceby>)
  - [Drop](<#func-drop>)
  - [DropWhile](<#func-dropwhile>)
  - [DropRightWhile](<#func-droprightwhile>)
  - [Duplicate](<#func-duplicate>)
  - [DuplicateWithIndex](<#func-duplicatewithindex>)
  - [Every](<#func-every>)
  - [Filter](<#func-filter>)
  - [FindAll](<#func-findall>)
  - [FindIndex](<#func-findindex>)
  - [FindLastIndex](<#func-findlastindex>)
  - [FindMax](<#func-findmax>)
  - [FindMaxBy](<#func-findmaxby>)
  - [FindMaxByKey](<#func-findmaxbykey>)
  - [FindMin](<#func-findmin>)
  - [FindMinBy](<#func-findminby>)
  - [FindMinByKey](<#func-findminbykey>)
  - [Flatten](<#func-flatten>)
  - [ForEach](<#func-foreach>)
  - [ForEachRight](<#func-foreachright>)
  - [GroupBy](<#func-groupby>)
  - [IndexOf](<#func-indexof>)
  - [Intersection](<#func-intersection>)
  - [IntersectionBy](<#func-intersectionby>)
  - [LastIndexOf](<#func-lastindexof>)
  - [Mean](<#func-mean>)
  - [Merge](<#func-merge>)
  - [Nth](<#func-nth>)
  - [Partition](<#func-partition>)
  - [PartitionMap](<#func-partitionmap>)
  - [Pluck](<#func-pluck>)
  - [Range](<#func-range>)
  - [RangeRight](<#func-rangeright>)
  - [Reduce](<#func-reduce>)
  - [Reject](<#func-reject>)
  - [Reverse](<#func-reverse>)
  - [Shuffle](<#func-shuffle>)
  - [SliceToMap](<#func-slicetomap>)
  - [Some](<#func-some>)
  - [Sum](<#func-sum>)
  - [SumBy](<#func-sumby>)
  - [ToSlice](<#func-toslice>)
  - [Union](<#func-union>)
  - [Unique](<#func-unique>)
  - [UniqueBy](<#func-uniqueby>)
  - [Unzip](<#func-unzip>)
  - [Without](<#func-without>)
  - [Zip](<#func-zip>)

- Map utility functions
  - [Filter2DMapCollection](<#func-filter2dmapcollection>)
  - [FilterMap](<#func-filtermap>)
  - [FilterMapCollection](<#func-filtermapcollection>)
  - [Find](<#func-find>)
  - [FindByKey](<#func-findbykey>)
  - [FindKey](<#func-findkey>)
  - [Keys](<#func-keys>)
  - [Map](<#func-map>)
  - [MapCollection](<#func-mapcollection>)
  - [MapContains](<#func-mapcontains>)
  - [MapEvery) bool](<#func-mapevery>)
  - [MapKeys](<#func-mapkeys>)
  - [MapSome](<#func-mapsome>)
  - [MapUnique](<#func-mapunique>)
  - [MapValues](<#func-mapvalues>)
  - [Omit](<#func-omit>)
  - [OmitBy](<#func-omitby>)
  - [Pick](<#func-pick>)
  - [PickBy](<#func-pickby>)
  - [Values](<#func-values>)

- Concurrency and time related utility functions
  - [After](<#func-after>)
  - [Before](<#func-before>)
  - [Delay](<#func-delay>)
  - [Flip](<#func-flip>)
  - [Memoize](<#func-memoizert-v-memoize>)
  - [NewDebounce](<#func-newdebounce>)
  - [Once](<#func-once>)
  - [Retry](<#func-rtypet-retry>)
  - [RetryWithDelay](<#func-rtypet-retrywithdelay>)

- Generic Data Structures
  - [`bst`](https://github.com/esimov/torx/tree/master/bstree): Binary Search Tree data structure implementation, where each node has at most two child nodes and the key of its internal node is greater than all the keys in the respective node's left subtree and less than the ones in the right subtree
  - [`btree`](https://github.com/esimov/torx/tree/master/btree): B-tree data structure implementation which is a self-balancing tree data structure maintaining its values in sorted order
  - [`cache`](https://github.com/esimov/torx/tree/master/cache): a basic in-memory key-value storage system
  - [`heap`](https://github.com/esimov/torx/tree/master/heap): Binary Heap data structure implementation where each node of the subtree is greather or equal then the parent node
  - [`list`](https://github.com/esimov/torx/tree/master/list): implementing a singly and doubly linked list data structure
  - [`queue`](https://github.com/esimov/torx/tree/master/queue): implementing a FIFO (First-In-First-Out) data structure in two forms: using as storage system a resizing array and a doubly linked list

## func Abs

```go
func Abs[T Number](x T) T
```

Abs returns the absolut value of x.

## func After

```go
func After[V constraints.Signed](n *V, fn func())
```

After creates a function wrapper that does nothing at first. From the nth call onwards, it starts actually calling the callback function. Useful for grouping responses, where you want to be sure that all the calls have finished, before proceeding.

## func Before

```go
func Before[S ~string, T any, V constraints.Signed](n *V, c *cache.Cache[S, T], fn func() T) T
```

Before creates a function wrapper that memoizes its return value. From the nth call onwards, the memoized result of the last invocation is returned immediately instead of invoking function again. So the wrapper will invoke function at most n\-1 times.

## func CamelCase

```go
func CamelCase[T ~string](str T) T
```

CamelCase converts a string to camelCase \(https://en.wikipedia.org/wiki/CamelCase\).

## func Capitalize

```go
func Capitalize[T ~string](str T) T
```

Capitalize converts the first letter of the string to uppercase and the remaining letters to lowercase.

## func Chunk

```go
func Chunk[T comparable](slice []T, size int) [][]T
```

Chunk split the slice into groups of slices each having the length of size. In case the source slice cannot be distributed equally, the last slice will contain fewer elements.

## func Clamp

```go
func Clamp[T Number](num, lo, up T) T
```

Clamp restricts a number between two other numbers.

## func Compare

```go
func Compare[T comparable](a, b T, comp CompFn[T]) int
```

Compare compares two values using as comparator the the callback function argument.

## func Contains

```go
func Contains[T comparable](slice []T, value T) bool
```

Contains returns true if the value is present in the collection.

## func Delay

```go
func Delay(delay time.Duration, fn func()) *time.Timer
```

Delay invokes the function with a predefined delay.

## func Difference

```go
func Difference[T comparable](s1, s2 []T) []T
```

Difference is similar to Without, but returns the values from the first slice that are not present in the second slice.

## func DifferenceBy

```go
func DifferenceBy[T comparable](s1, s2 []T, fn func(T) T) []T
```

DifferenceBy is like Difference, except that invokes a callback function on each element of the slice, applying the criteria by which the difference is computed.

## func Drop

```go
func Drop[T any](slice []T, n int) []T
```

Drop creates a new slice with n elements dropped from the beginning. If n \< 0 the elements will be dropped from the back of the collection.

## func DropRightWhile

```go
func DropRightWhile[T any](slice []T, fn func(T) bool) []T
```

DropRightWhile creates a new slice excluding the elements dropped from the end. Elements are dropped by applying the conditional invoked in the callback function.

## func DropWhile

```go
func DropWhile[T any](slice []T, fn func(T) bool) []T
```

DropWhile creates a new slice excluding the elements dropped from the beginning. Elements are dropped by applying the conditional invoked in the callback function.

## func Duplicate

```go
func Duplicate[T comparable](slice []T) []T
```

Duplicate returns the duplicated values of a collection.

## func DuplicateWithIndex

```go
func DuplicateWithIndex[T comparable](slice []T) map[T]int
```

DuplicateWithIndex puts the duplicated values of a collection into a map as a key value pair, where the key is the collection element and the value is it's position.

## func Equal

```go
func Equal[T comparable](a, b T) bool
```

Equal checks if two values are equal.

## func Every

```go
func Every[T any](slice []T, fn func(T) bool) bool
```

Every returns true if all of the elements of a slice satisfies the criteria of the callback function.

## func Filter

```go
func Filter[T any](slice []T, fn func(T) bool) []T
```

Filter returns all the elements from the collection which satisfies the conditional logic of the callback function.

## func Filter2DMapCollection

```go
func Filter2DMapCollection[K comparable, V any](collection []map[K]map[K]V, fn func(map[K]V) bool) []map[K]map[K]V
```

Filter2DMapCollection filter out a two dimmensional collection of map items by applying the conditional logic of the callback function.

## func FilterMap

```go
func FilterMap[K comparable, V any](m map[K]V, fn func(V) bool) map[K]V
```

FilterMap iterates over the elements of a collection and returns a new collection representing all the items which satisfies the criteria formulated in the callback function.

## func FilterMapCollection

```go
func FilterMapCollection[K comparable, V any](collection []map[K]V, fn func(V) bool) []map[K]V
```

FilterMapCollection filter out a one dimmensional collection of map items by applying the conditional logic of the callback function.

## func Find

```go
func Find[K constraints.Ordered, V any](m map[K]V, fn func(V) bool) map[K]V
```

Find iterates over the elements of a map and returns the first item for which the callback function returns true.

## func FindAll

```go
func FindAll[T any](s []T, fn func(T) bool) map[int]T
```

FindAll is like FindIndex, but inserts into a map all the values which stisfies the conditional logic of the callback function. The map key represents the position of the found value and the value is the item itself.

## func FindByKey

```go
func FindByKey[K comparable, V any](m map[K]V, fn func(K) bool) map[K]V
```

FindByKey is like Find, but returns the first item for which the callback function returns true.

## func FindIndex

```go
func FindIndex[T any](s []T, fn func(T) bool) int
```

FindIndex returns the index of the first found element.

## func FindKey

```go
func FindKey[K comparable, V any](m map[K]V, fn func(V) bool) K
```

FindKey is like Find, but returns the first item key position for which the callback function returns true.

## func FindLastIndex

```go
func FindLastIndex[T any](s []T, fn func(T) bool) int
```

FindLastIndex is like FindIndex, only that returns the index of last found element.

## func FindMax

```go
func FindMax[T constraints.Ordered](s []T) T
```

FindMax finds the maximum value of a slice.

## func FindMaxBy

```go
func FindMaxBy[T constraints.Ordered](s []T, fn func(val T) T) T
```

FindMaxBy is like FindMax except that it accept a callback function and the conditional logic is applied over the resulted value. If there are more than one identical values resulted from the callback function the first one is used.

## func FindMaxByKey

```go
func FindMaxByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) (T, error)
```

FindMaxByKey finds the maximum value from a map by using some existing key as a parameter.

## func FindMin

```go
func FindMin[T constraints.Ordered](s []T) T
```

FindMin finds the minumum value of a slice.

## func FindMinBy

```go
func FindMinBy[T constraints.Ordered](s []T, fn func(val T) T) T
```

FindMinBy is like FindMin except that it accept a callback function and the conditional logic is applied over the resulted value. If there are more than one identical values resulted from the callback function the first one is used.

## func FindMinByKey

```go
func FindMinByKey[K comparable, T constraints.Ordered](mapSlice []map[K]T, key K) (T, error)
```

FindMinByKey finds the minimum value from a map by using some existing key as a parameter.

## func Flatten

```go
func Flatten[T any](slice any) ([]T, error)
```

Flatten flattens the slice all the way to the deepest nesting level.

## func Flip

```go
func Flip[T any](fn func(args ...T) []T) func(args ...T) []T
```

Flip creates a function that invokes fn with arguments reversed.

## func ForEach

```go
func ForEach[T any](slice []T, fn func(T))
```

ForEach iterates over the elements of a collection and invokes the callback fn function on each element.

## func ForEachRight

```go
func ForEachRight[T any](slice []T, fn func(T))
```

ForEachRight is the same as ForEach, but starts the iteration from the last element.

## func GroupBy

```go
func GroupBy[T1, T2 comparable](slice []T1, fn func(T1) T2) map[T2][]T1
```

GroupBy splits a collection into a key\-value set, grouped by the result of running each value through the callback function fn. The return value is a map where the key is the conditional logic of the callback function and the values are the callback function returned values.

## func InRange

```go
func InRange[T Number](num, lo, up T) bool
```

InRange checks if a number is inside a range.

## func IndexOf

```go
func IndexOf[T comparable](s []T, val T) int
```

IndexOf returns the index at which value can be found in the slice, or \-1 if value is not present in the slice.

## func Intersection

```go
func Intersection[T comparable](params ...[]T) []T
```

Intersection computes the list of values that are the intersection of all the slices. Each value in the result should be present in each of the provided slices.

## func IntersectionBy

```go
func IntersectionBy[T comparable](fn func(T) T, params ...[]T) []T
```

IntersectionBy is like Intersection, except that it accepts and callback function which is invoked on each element of the collection.

## func Invert

```go
func Invert[K, V comparable](m map[K]V) map[V]K
```

Invert returns a copy of the map where the keys become the values and the values the keys. For this to work, all of your map's values should be unique.

## func KebabCase

```go
func KebabCase[T ~string](str T) T
```

KebabCase converts a string to kebab\-case \(https://en.wikipedia.org/wiki/Letter_case#Kebab_case\).

## func Keys

```go
func Keys[K comparable, V any](m map[K]V) []K
```

Keys retrieve all the existing keys of a map.

## func LastIndexOf

```go
func LastIndexOf[T comparable](s []T, val T) int
```

LastIndexOf returns the index of the last occurrence of a value.

## func Less

```go
func Less[T constraints.Ordered](a, b T) bool
```

Less checks if the first value is less than the second.

## func Map

```go
func Map[T1, T2 any](slice []T1, fn func(T1) T2) []T2
```

Map produces a new slice of values by mapping each value in the list through a transformation function.

## func MapCollection

```go
func MapCollection[K comparable, V any](m map[K]V, fn func(V) V) []V
```

MapCollection is like the Map method applied on slices, but this time applied on maps. It runs each element over an iteratee function and saves the resulted values into a new map.

## func MapContains

```go
func MapContains[K, V comparable](m map[K]V, value V) bool
```

MapContains returns true if the value is present in the list otherwise false.

## func MapEvery

```go
func MapEvery[K comparable, V any](m map[K]V, fn func(V) bool) bool
```

MapEvery returns true if all of the elements of a map satisfies the criteria of the callback function.

## func MapKeys

```go
func MapKeys[K comparable, V any, R comparable](m map[K]V, fn func(K, V) R) map[R]V
```

MapKeys is the opposite of MapValues. It creates a new map with the same number of elements as the original one, but this time the callback function \(fn\) is invoked over the map keys.

## func MapSome

```go
func MapSome[K comparable, V any](m map[K]V, fn func(V) bool) bool
```

MapSome returns true if some of the elements of a map satisfies the criteria of the callback function.

## func MapUnique

```go
func MapUnique[K, V comparable](m map[K]V) map[K]V
```

MapUnique removes the duplicate values from a map.

## func MapValues

```go
func MapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R
```

MapValues creates a new map with the same number of elements as the original one, but running each map value through a callback function \(fn\).

## func Max

```go
func Max[T constraints.Ordered](values ...T) T
```

Max returns the biggest value from the provided parameters.

## func Mean

```go
func Mean[T Number](slice []T) T
```

Mean computes the mean value of the slice elements.

## func Merge

```go
func Merge[T any](s []T, params ...[]T) []T
```

Merge merges the first slice with the other slices defined as variadic parameter.

## func Min

```go
func Min[T constraints.Ordered](values ...T) T
```

Min returns the lowest value from the provided parameters.

## func N

```go
func N[T Number](s string) (T, error)
```

N converts a string to a generic number.

## func NewDebounce

```go
func NewDebounce(wait time.Duration) (func(f func()), func())
```

NewDebounce creates a new debounced version of the invoked function which will postpone the execution until the time duration has elapsed since the last invocation passed in as a function argument.

It returns a callback function which will be invoked after the predefined delay and also a cancel function which should be invoked to cancel a scheduled debounce.

## func Nth

```go
func Nth[T any](slice []T, nth int) (T, error)
```

Nth returns the nth element of the collection. In case of negative value the nth element is returned from the end of the collection. In case nth is out of bounds an error is returned.

## func Null

```go
func Null[T any]() T
```

## func NumToString

```go
func NumToString[T Number](n T) string
```

NumToString converts a number to a string. In case of a number of type float \(float32|float64\) this will be rounded to 2 decimal places.

## func Omit

```go
func Omit[K comparable, V any](collection map[K]V, keys ...K) map[K]V
```

Omit is the opposite of Pick, it extracts all the map elements which keys are not omitted.

## func OmitBy

```go
func OmitBy[K comparable, V any](collection map[K]V, fn func(key K, val V) bool) map[K]V
```

OmitBy is the opposite of Omit, it removes all the map elements for which the callback function returns true.

## func Once

```go
func Once[S ~string, T any](c *cache.Cache[S, T], fn func() T) T
```

Once is like Before, but it's invoked only once. Repeated calls to the modified function will have no effect, returning the value from the cache.

## func Pad

```go
func Pad[T ~string](str T, size int, token string) T
```

Pads string on the left and right sides if it's shorter than length. Padding characters are truncated if they can't be evenly divided by length.

## func PadLeft

```go
func PadLeft[T ~string](str T, size int, token string) T
```

PadLeft pads string on the left side if it's shorter than length. Padding characters are truncated if they exceed length.

## func PadRight

```go
func PadRight[T ~string](str T, size int, token string) T
```

PadRight pads string on the right side if it's shorter than length. Padding characters are truncated if they exceed length.

## func Partition

```go
func Partition[T comparable](slice []T, fn func(T) bool) [2][]T
```

Partition splits the collection elements into two, the ones which satisfies the condition expressed in the callback function \(fn\) and those which does not satisfies the condition.

## func PartitionMap

```go
func PartitionMap[K comparable, V any](mapSlice []map[K]V, fn func(map[K]V) bool) [2][]map[K]V
```

PartitionMap split the collection into two arrays, the one whose elements satisfies the condition expressed in the callback function \(fn\) and one whose elements don't satisfies the condition.

## func Pick

```go
func Pick[K comparable, V any](collection map[K]V, keys ...K) (map[K]V, error)
```

Pick extracts the elements from the map which have the key defined in the allowed keys.

## func PickBy

```go
func PickBy[K comparable, V any](collection map[K]V, fn func(key K, val V) bool) map[K]V
```

PickBy extracts all the map elements for which the callback function returns truthy.

## func Pluck

```go
func Pluck[K comparable, V any](mapSlice []map[K]V, key K) []V
```

Pluck extracts all the values of a map by the key definition.

## func Range

```go
func Range[T Number](args ...T) ([]T, error)
```

Range creates a slice of numbers \(integers\) progressing from start up to, but not including end. This method can accept 1, 2 or 3 arguments. Depending on the number of provided parameters, \`start\`, \`step\` and \`end\` has the following meaning: \[start=0\]: The start of the range. If ommited it defaults to 0. \[step=1\]: The value to increment or decrement by. end: The end of the range. In case you'd like negative values, use a negative step.

## func RangeRight

```go
func RangeRight[T Number](params ...T) ([]T, error)
```

RangeRight is like Range, only that it populates the slice in descending order.

## func Reduce

```go
func Reduce[T1, T2 any](slice []T1, fn func(T1, T2) T2, initVal T2) T2
```

Reduce reduces the collection to a value which is the accumulated result of running each element in the collection through the callback function yielding a single value.

## func Reject

```go
func Reject[T any](slice []T, fn func(val T) bool) []T
```

Reject is the opposite of Filter. It returns the values from the collection without the elements for which the callback function returns true.

## func Reverse

```go
func Reverse[T any](sl []T) []T
```

Reverse reverses the order of elements, so that the first element becomes the last, the second element becomes the second to last, and so on.

## func ReverseStr

```go
func ReverseStr[T ~string](str T) T
```

ReverseStr returns a new string with the characters in reverse order.

## func Shuffle

```go
func Shuffle[T any](src []T) []T
```

Shuffle implements the Fisher\-Yates shuffle algorithm applied to a slice.

## func SliceToMap

```go
func SliceToMap[K comparable, T any](s1 []K, s2 []T) map[K]T
```

SliceToMap converts a slice to a map. It panic in case the parameter slices lenght are not identical. The map keys will be the items from the first slice and the values the items from the second slice.

## func SnakeCase

```go
func SnakeCase[T ~string](str T) T
```

SnakeCase converts a string to snake\_case \(https://en.wikipedia.org/wiki/Snake_case\).

## func Some

```go
func Some[T any](slice []T, fn func(T) bool) bool
```

Some returns true if some of the elements of a slice satisfies the criteria of the callback function.

## func SplitAtIndex

```go
func SplitAtIndex[T ~string](str T, index int) []T
```

SplitAtIndex split the string at the specified index and returns a slice with the resulted two substrings.

## func Substr

```go
func Substr[T ~string](str T, offset, length int) T
```

If length is negative, then that many characters will be omitted from the end of string starting from the offset position.

## func Sum

```go
func Sum[T Number](slice []T) T
```

Sum returns the sum of the slice items. These have to satisfy the type constraints declared as Number.

## func SumBy

```go
func SumBy[T1 any, T2 Number](slice []T1, fn func(T1) T2) T2
```

SumBy is like Sum except the it accept a callback function which is invoked for each element in the slice to generate the value to be summed.

## func ToLower

```go
func ToLower[T ~string](str T) T
```

ToLower converts a string to Lowercase.

## func ToSlice

```go
func ToSlice[T any](args ...T) []T
```

ToSlice returns the function arguments as a slice.

## func ToUpper

```go
func ToUpper[T ~string](str T) T
```

ToUpper converts a string to Uppercase.

## func Union

```go
func Union[T comparable](slice any) ([]T, error)
```

Union computes the union of the passed\-in slice and returns in order the list of unique items that are present in one or more of the slices.

## func Unique

```go
func Unique[T comparable](slice []T) []T
```

Unique returns the collection unique values.

## func UniqueBy

```go
func UniqueBy[T comparable](slice []T, fn func(T) T) []T
```

UniqueBy is like Unique except that it accept a callback function which is invoked on each element of the slice applying the criteria by which the uniqueness is computed.

## func Unwrap

```go
func Unwrap[T ~string](str T, token string) T
```

Unwrap a string with the specified token.

## func Unzip

```go
func Unzip[T any](slices ...[]T) [][]T
```

Zip iteratively merges together the values of the slice parameters with the values at the corresponding position.

## func Values

```go
func Values[K comparable, V any](m map[K]V) []V
```

Values retrieve all the existing values of a map.

## func Without

```go
func Without[T1 comparable, T2 any](slice []T1, values ...T1) []T1
```

Without returns a copy of the slice with all the values defined in the variadic parameter removed.

## func Wrap

```go
func Wrap[T ~string](str T, token string) T
```

Wrap a string with the specified token.

## func WrapAllRune

```go
func WrapAllRune[T ~string](str T, token string) T
```

## func Zip

```go
func Zip[T any](slices ...[]T) [][]T
```

Zip iteratively merges together the values of the slice parameters with the values at the corresponding position.

## type Bound

```go
type Bound[T constraints.Signed] struct {
    Min, Max T
}
```

### func \(Bound\[T\]\) Enclose

```go
func (b Bound[T]) Enclose(nth T) bool
```

Enclose checks if an element is inside the bounds.

## type CompFn

CompFn is a generic function type for comparing two values.

```go
type CompFn[T any] func(a, b T) bool
```

## type Memoizer

Memoizer is a two component struct type used to memoize the results of a function execution. It holds an exported Cache storage and a singleflight.Group which is used to guarantee that only one function execution is in flight for a given key.

```go
type Memoizer[T ~string, V any] struct {
    Cache *cache.Cache[T, V]
    // contains filtered or unexported fields
}
```

### func NewMemoizer

```go
func NewMemoizer[T ~string, V any](expiration, cleanup time.Duration) *Memoizer[T, V]
```

NewMemoizer instantiates a new Memoizer.

### func \(Memoizer\[T, V\]\) Memoize

```go
func (m Memoizer[T, V]) Memoize(key T, fn func() (*cache.Item[V], error)) (*cache.Item[V], error)
```

Memoize returns the item under a specific key instantly in case the key exists, otherwise returns the results of the given function, making sure that only one execution is in\-flight for a given key at a time. This method is useful for caching the result of a time consuming operation when is more important to return a slightly outdated result, than to wait for an operation to complete before serving it.

## type Number

Number is a custom type set of constraints extending the Float and Integer type set from the experimental constraints package.

```go
type Number interface {
    // contains filtered or unexported methods
}
```

## type RType

RType is a generic struct type used as method receiver on retry operations.

```go
type RType[T any] struct {
    Input T
}
```

### func \(RType\[T\]\) Retry

```go
func (v RType[T]) Retry(n int, fn func(T) error) (int, error)
```

Retry tries to invoke the callback function n times. It runs until the number of attempts is reached or the returned value of the callback function is nil.

### func \(RType\[T\]\) RetryWithDelay

```go
func (v RType[T]) RetryWithDelay(n int, delay time.Duration, fn func(time.Duration, T) error) (time.Duration, int, error)
```

RetryWithDelay tries to invoke the callback function n times, but with a delay between each calls. It runs until the number of attempts is reached or the error return value of the callback function is nil.



