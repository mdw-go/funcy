package funcy

import "sort"

func Is[T any](v any) bool {
	_, ok := v.(T)
	return ok
}
func As[T any](v any) T {
	t, _ := v.(T)
	return t
}
func FilterAs[T any](collection []any) []T {
	return Map(As[T], Filter(Is[T], collection))
}
func Filter[T any](predicate func(t T) bool, values []T) (result []T) {
	for _, value := range values {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}
func Remove[T any](predicate func(t T) bool, values []T) []T {
	return Filter(Complement(predicate), values)
}
func Complement[T any](predicate func(t T) bool) func(t T) bool {
	return func(t T) bool { return !predicate(t) }
}
func Map[I, O any](transform func(i I) O, values []I) (result []O) {
	for _, value := range values {
		result = append(result, transform(value))
	}
	return result
}
func MapCat[I, O any](transform func(i I) []O, values []I) (result []O) {
	for _, value := range values {
		result = append(result, transform(value)...)
	}
	return result
}
func Reduce[T any](calc func(a, b T) T, start T, values []T) (result T) {
	result = start
	for _, next := range values {
		result = calc(result, next)
	}
	return result
}
func Sum[T Number](rolls []T) T {
	return Reduce(Add[T], 0, rolls)
}
func Add[T Number](a, b T) T {
	return a + b
}
func Product[T Number](rolls []T) T {
	return Reduce(Multiply[T], 1, rolls)
}
func Multiply[T Number](a, b T) T {
	return a * b
}
func Range[N Number](start, stop N) (result []N) {
	for start != stop {
		result = append(result, start)
		if start < stop {
			start++
		} else {
			start--
		}
	}
	return result
}
func Take[T any](n int, values []T) []T {
	if n >= len(values) {
		n = len(values)
	}
	return values[:n]
}
func Drop[T any](n int, values []T) []T {
	if n >= len(values) {
		n = len(values)
	}
	return values[n:]
}
func First[T any](values []T) T {
	return values[0]
}
func Last[T any](values []T) T {
	return values[len(values)-1]
}
func Rest[T any](values []T) []T {
	return Drop(1, values)
}
func AllBut[T any](n int, values []T) []T {
	return Drop(len(values)-n, values)
}
func TakeWhile[T any](predicate func(T) bool, values []T) (result []T) {
	for _, value := range values {
		if !predicate(value) {
			break
		}
		result = append(result, value)
	}
	return result
}
func DropWhile[T any](predicate func(T) bool, values []T) (result []T) {
	for _, value := range values {
		if predicate(value) {
			continue
		}
		result = append(result, value)
	}
	return result
}
func IndexBy[K comparable, V any](key func(V) K, list []V) map[K]V {
	result := make(map[K]V)
	for _, value := range list {
		result[key(value)] = value
	}
	return result
}
func SlicedIndexBy[K comparable, V any](key func(V) K, list []V) map[K][]V {
	result := make(map[K][]V)
	for _, value := range list {
		key := key(value)
		result[key] = append(result[key], value)
	}
	return result
}
func GroupBy[K comparable, V any](key func(V) K, list []V) (result [][]V) {
	// TODO: return maps.Values(SlicedIndexBy(key, list))
	// (but only once the maps package moves from golang.org/x/exp into the std lib)
	index := SlicedIndexBy(key, list)
	for _, value := range index {
		result = append(result, value)
	}
	return result
}
func Drain[T any](channel <-chan T) (slice []T) {
	for item := range channel {
		slice = append(slice, item)
	}
	return slice
}
func Load[T any](result chan<- T, stream []T) {
	defer close(result)
	for _, item := range stream {
		result <- item
	}
}
func SortAscending[C LessThan, V any](key func(V) C, original []V) (result []V) {
	collection := make([]V, len(original))
	copy(collection, original)
	sort.Slice(collection, func(i, j int) bool { return key(collection[i]) < key(collection[j]) })
	return collection
}
func SortDescending[C LessThan, V any](key func(V) C, original []V) (result []V) {
	collection := make([]V, len(original))
	copy(collection, original)
	sort.Slice(collection, func(i, j int) bool { return key(collection[i]) > key(collection[j]) })
	return collection
}
func Frequencies[T comparable](values []T) map[T]int {
	result := make(map[T]int)
	for _, v := range values {
		result[v]++
	}
	return result
}
func Flatten[T any](matrix [][]T) (result []T) {
	for _, row := range matrix {
		for _, col := range row {
			result = append(result, col)
		}
	}
	return result
}
func Zip[A, B any](a []A, b []B) (result []Pair[A, B]) {
	length := len(a)
	if len(b) < len(a) {
		length = len(b)
	}
	for x := 0; x < length; x++ {
		result = append(result, Pair[A, B]{A: a[x], B: b[x]})
	}
	return result
}
func Repeat[T any](n int, t T) (result []T) {
	for ; n > 0; n-- {
		result = append(result, t)
	}
	return result
}

type Pair[A, B any] struct {
	A A
	B B
}

type (
	Number interface {
		Integer | uintptr | float64 | float32
	}
	Integer interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
	}
	LessThan interface {
		Number | string
	}
)
