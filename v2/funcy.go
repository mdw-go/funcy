package funcy

import "iter"

func Seq[S ~[]T, T any](s S) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
func Seq2[S ~[]V, V any](s S) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for k, v := range s {
			if !yield(k, v) {
				return
			}
		}
	}
}
func Slice[T any](seq iter.Seq[T]) (result []T) {
	for v := range seq {
		result = append(result, v)
	}
	return result
}
func Slice2[K, V any](seq iter.Seq2[K, V]) (result []V) {
	for _, v := range seq {
		result = append(result, v)
	}
	return result
}
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
