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
	next, stop := iter.Pull[V](s)
	defer stop()
	v, ok := next()
	if !ok {
		panic("runtime error: index out of range [0] with length 0")
	}
	return v
}
func Last[V any](s iter.Seq[V]) V {
	next, stop := iter.Pull[V](s)
	defer stop()
	var prev V
	for x := 0; ; x++ {
		this, ok := next()
		if !ok && x == 0 {
			panic("runtime error: index out of range [0] with length 0")
		} else if !ok {
			return prev
		}
		prev = this
	}

}
func Take[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull[V](s)
		defer stop()
		for x := 0; x < n; x++ {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}
func Drop[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		next, stop := iter.Pull[V](s)
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
