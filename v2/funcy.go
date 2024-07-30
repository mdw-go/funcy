package funcy

import "iter"

func Seq[S ~[]V, V any](s S) iter.Seq[V] {
	return func(yield func(V) bool) {
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
func Seq2Seq[K, V any](seq2 iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq2 {
			if !yield(v) {
				return
			}
		}
	}
}
func Slice[V any](seq iter.Seq[V]) (result []V) {
	for v := range seq {
		result = append(result, v)
	}
	return result
}
func Range(start, stop int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for x := start; x < stop; x++ {
			if !yield(x) {
				return
			}
		}
	}
}
func First[V any](s iter.Seq[V]) V {
	for v := range s {
		return v
	}
	panic("runtime error: index out of range [0] with length 0")
}
func Last[V any](s iter.Seq[V]) (result V) {
	count := 0
	for result = range s {
		count++
	}
	if count == 0 {
		panic("runtime error: index out of range [0] with length 0")
	}
	return result
}
func Take[V any](count int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		n := 0
		for v := range s {
			if n < count {
				if !yield(v) {
					return
				}
			}
			n++
		}
	}
}
func Drop[V any](count int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		n := 0
		for v := range s {
			if n >= count {
				if !yield(v) {
					return
				}
			}
			n++
		}
	}
}
func Filter[V any](predicate func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for s := range seq {
			if predicate(s) {
				if !yield(s) {
					return
				}
			}
		}
	}
}
func Remove[V any](predicate func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return Filter(Complement(predicate), seq)
}
func Complement[V any](predicate func(t V) bool) func(t V) bool {
	return func(t V) bool { return !predicate(t) }
}
func Map[I, O any](f func(I) O, seq iter.Seq[I]) iter.Seq[O] {
	return func(yield func(O) bool) {
		for s := range seq {
			if !yield(f(s)) {
				return
			}
		}
	}
}
func Reduce[V any](calc func(a, b V) V, start V, seq iter.Seq[V]) (result V) {
	result = start
	for next := range seq {
		result = calc(result, next)
	}
	return result
}
