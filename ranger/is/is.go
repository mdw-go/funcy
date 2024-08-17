package is

import (
	"iter"
	"reflect"
)

func abs[T Integer](t T) T {
	if t < 0 {
		return -t
	}
	return t
}
func Zero[T Number](t T) bool     { return t == 0 }
func Positive[T Number](t T) bool { return t > 0 }
func Negative[T Number](t T) bool { return t < 0 }
func Even[T Integer](t T) bool    { return t%2 == 0 }
func Odd[T Integer](t T) bool     { return abs(t)%2 == 1 }
func Nil[T any](t T) bool {
	if any(t) == nil {
		return true
	}
	switch v := reflect.ValueOf(t); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}
func Empty[V any](s iter.Seq[V]) bool {
	for _ = range s {
		return false
	}
	return true
}
