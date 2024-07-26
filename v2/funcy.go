package funcy

import "iter"

func Range(start, stop int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for x := start; x < stop; x++ {
			if !yield(x) {
				return
			}
		}
	}
}
func First[T any](s iter.Seq[T]) T {
	next, stop := iter.Pull[T](s)
	defer stop()
	v, ok := next()
	if !ok {
		panic("runtime error: index out of range [0] with length 0")
	}
	return v
}
func Take[T any](n int, s iter.Seq[T]) func(func(T) bool) {
	return func(yield func(T) bool) {
		next, stop := iter.Pull[T](s)
		defer stop()
		for x := 0; x < n; x++ {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}
func Drop[T any](n int, s iter.Seq[T]) func(func(T) bool) {
	return func(yield func(T) bool) {
		next, stop := iter.Pull[T](s)
		defer stop()
		for x := 0; ; x++ {
			v, ok := next()
			if !ok {
				return
			}
			if x < n {
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}
