package funcy

import (
	"fmt"
	"reflect"
	"sort"
)

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
func MapVoid[T any](f func(T), values []T) {
	for _, value := range values {
		f(value)
	}
}
func Map[I, O any](transform func(i I) O, values []I) (result []O) {
	for _, value := range values {
		result = append(result, transform(value))
	}
	return result
}
func Map2[Ia, Ib, O any](transform func(Ia, Ib) O, a []Ia, b []Ib) (result []O) {
	length := len(a)
	if len(b) < len(a) {
		length = len(b)
	}
	for x := 0; x < length; x++ {
		result = append(result, transform(a[x], b[x]))
	}
	return result
}
func MapCat[I, O any](transform func(i I) []O, values []I) (result []O) {
	for _, value := range values {
		result = append(result, transform(value)...)
	}
	return result
}
func MapAsAny[T any](items []T) (result []any) {
	for _, item := range items {
		result = append(result, item)
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
func Min[T LessThan](all []T) T {
	min := all[0]
	for _, a := range all[1:] {
		if a < min {
			min = a
		}
	}
	return min
}
func Max[T LessThan](all []T) T {
	max := all[0]
	for _, a := range all[1:] {
		if a > max {
			max = a
		}
	}
	return max
}
func Abs[T Number](t T) T {
	if t < 0 {
		return -t
	}
	return t
}
func Square[T Number](t T) T         { return t * t }
func IsEven[T Integer](t T) bool     { return t%2 == 0 }
func IsOdd[T Integer](t T) bool      { return t%2 == 1 }
func String[T any](t T) string       { return fmt.Sprint(t) }
func ByLength[T any](t T) int        { return reflect.ValueOf(t).Len() }
func ByNumericValue[T Number](t T) T { return t }
func Range[N Number](start, stop N) (result []N) {
	return RangeStep(start, stop, 1)
}
func RangeStep[N Number](start, stop, step N) (result []N) {
	var (
		compare func(start, stop N) bool
		next    func(start, step N) N
	)
	if start < stop {
		compare = func(start, stop N) bool { return start < stop }
		next = func(start, step N) N { return start + step }
	} else {
		compare = func(start, stop N) bool { return start > stop }
		next = func(start, step N) N { return start - step }
	}

	for compare(start, stop) {
		result = append(result, start)
		start = next(start, step)
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
func FirstOrDefault[T any](values []T) (default_ T) {
	if len(values) == 0 {
		return default_
	}
	return First(values)
}
func FirstNonDefault[T comparable](values ...T) (zero T) {
	for _, value := range values {
		if value != zero {
			return value
		}
	}
	return zero
}
func Last[T any](values []T) T {
	return values[len(values)-1]
}
func LastOrDefault[T any](values []T) (default_ T) {
	if len(values) == 0 {
		return default_
	}
	return Last(values)
}
func Rest[T any](values []T) []T {
	return Drop(1, values)
}
func DropAllBut[T any](n int, values []T) []T {
	return Drop(len(values)-n, values)
}
func TakeAllBut[T any](n int, values []T) []T {
	return Take(len(values)-n, values)
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
func KeyValueIndexBy[K comparable, V1, V2 any](kv func(V1) (K, V2), list []V1) map[K]V2 {
	result := make(map[K]V2)
	for _, item := range list {
		key, value := kv(item)
		result[key] = value
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
func SlicedKeyValueIndexBy[K comparable, V1, V2 any](kv func(V1) (K, V2), list []V1) map[K][]V2 {
	result := make(map[K][]V2)
	for _, item := range list {
		key, value := kv(item)
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
func Partition[T any](length, increment int, values []T) (result [][]T) {
	for x := 0; x+length <= len(values); x += increment {
		result = append(result, values[x:x+length])
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
func ZipMap[A comparable, B any](a []A, b []B) (result map[A]B) {
	length := len(a)
	if len(b) < len(a) {
		length = len(b)
	}
	result = make(map[A]B, length)
	for x := 0; x < length; x++ {
		result[a[x]] = b[x]
	}
	return result
}
func Repeat[T any](n int, t T) (result []T) {
	for ; n > 0; n-- {
		result = append(result, t)
	}
	return result
}
func All(inputs []bool) bool  { return len(TakeWhile(isTrue, inputs)) == len(inputs) }
func Any(inputs []bool) bool  { return len(DropWhile(Complement(isTrue), inputs)) > 0 }
func None(inputs []bool) bool { return len(TakeWhile(Complement(isTrue), inputs)) == len(inputs) }
func isTrue(b bool) bool      { return b }
func MapKeys[K comparable, V any](m map[K]V) (results []K) {
	results = make([]K, 0, len(m))
	for key := range m {
		results = append(results, key)
	}
	return results
}
func MapValues[K comparable, V any](m map[K]V) (results []V) {
	results = make([]V, 0, len(m))
	for _, value := range m {
		results = append(results, value)
	}
	return results
}
func MapPairs[K comparable, V any](m map[K]V) (results []Pair[K, V]) {
	results = make([]Pair[K, V], 0, len(m))
	for key, value := range m {
		results = append(results, Pair[K, V]{A: key, B: value})
	}
	return results
}
func PairsMap[K comparable, V any](pairs []Pair[K, V]) map[K]V {
	result := make(map[K]V, len(pairs))
	for _, p := range pairs {
		result[p.A] = p.B
	}
	return result
}
func MapLookup[K comparable, V any](m map[K]V) func(K) V { return func(k K) V { return m[k] } }
func SliceLookup[V any](s []V) func(int) V               { return func(i int) V { return s[i] } }

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
