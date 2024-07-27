package is

import "reflect"

func Zero[T Number](t T) bool     { return t == 0 }
func Positive[T Number](t T) bool { return t > 0 }
func Negative[T Number](t T) bool { return t < 0 }
func Even[T Integer](t T) bool    { return t%2 == 0 }
func Odd[T Integer](t T) bool     { return t%2 == 1 }
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
