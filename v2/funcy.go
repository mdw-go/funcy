package funcy

import (
	"iter"
	"slices"

	"github.com/mdwhatcott/funcy/v2/internal/ring"
	"github.com/mdwhatcott/funcy/v2/is"
)

/*
TODO:
- https://clojuredocs.org/clojure.core/conj
- https://clojuredocs.org/clojure.core/doall
- https://clojuredocs.org/clojure.core/frequencies
- https://clojuredocs.org/clojure.core/group-by
- https://clojuredocs.org/clojure.core/interpose
- https://clojuredocs.org/clojure.core/iterate
- https://clojuredocs.org/clojure.core/merge
- https://clojuredocs.org/clojure.core/partition
- https://clojuredocs.org/clojure.core/reductions
- https://clojuredocs.org/clojure.core/repeatedly
- https://clojuredocs.org/clojure.core/sort-by
- https://clojuredocs.org/clojure.core/zipmap
*/

func Variadic[V any](vs ...V) iter.Seq[V] {
	return Iterate(vs)
}
func Iterate[S ~[]V, V any](s S) iter.Seq[V] {
	return slices.Values(s)
}
func Slice[V any](seq iter.Seq[V]) (result []V) {
	return slices.Collect(seq)
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
	for v := range s {
		result = v
		count++
	}
	if count > 0 {
		return result
	}
	panic("runtime error: index out of range [0] with length 0")
}
func Take[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		count := 0
		for v := range s {
			if count >= n {
				return
			}
			if !yield(v) {
				return
			}
			count++
		}
	}
}
func TakeWhile[V any](pred func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range s {
			if !pred(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}
func TakeLast[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		r := ring.New[V](n)
		count := 0
		for v := range s {
			count++
			r.Value = v
			r = r.Next()
		}
		for x := 0; x < count; x++ {
			r = r.Next()
		}
		for j := 0; j < count; j++ {
			if !yield(r.Value) {
				return
			}
			r = r.Next()
		}
	}
}
func Drop[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		count := 0
		for v := range s {
			if count < n {
				count++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}
func DropWhile[V any](pred func(V) bool, s iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		dropping := true
		for v := range s {
			if dropping && pred(v) {
				continue
			} else if dropping {
				dropping = false
			}
			if !yield(v) {
				return
			}
		}
	}
}
func DropLast[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return Map2(func(a, _ V) V { return a }, s, Drop(n, s))
}
func Rest[V any](s iter.Seq[V]) iter.Seq[V] {
	return Drop(1, s)
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
func Map2[A, B, O any](f func(A, B) O, a iter.Seq[A], b iter.Seq[B]) iter.Seq[O] {
	return func(yield func(O) bool) {
		nextA, stopA := iter.Pull(a)
		defer stopA()
		nextB, stopB := iter.Pull(b)
		defer stopB()
		for {
			aa, okA := nextA()
			if !okA {
				return
			}
			bb, okB := nextB()
			if !okB {
				return
			}
			if !yield(f(aa, bb)) {
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
func Repeat[V any](n int, v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for range n {
			if !yield(v) {
				return
			}
		}
	}
}
func Concat[V any](all ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range all {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}
func Sum[N is.Number](seq iter.Seq[N]) (zero N) {
	add := func(a, b N) N { return a + b }
	return Reduce(add, zero, seq)
}
func Nest[V any](matrix [][]V) iter.Seq[iter.Seq[V]] {
	return func(yield func(iter.Seq[V]) bool) {
		for _, row := range matrix {
			_ = yield(Iterate(row))
		}
	}
}
func Flatten[V any](matrix iter.Seq[iter.Seq[V]]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for row := range matrix {
			for cell := range row {
				if !yield(cell) {
					return
				}
			}
		}
	}
}
func Count[V any](seq iter.Seq[V]) (result int) {
	for _ = range seq {
		result++
	}
	return result
}
func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			for s := range seq {
				if !yield(s) {
					return
				}
			}
		}
	}
}
func Interleave[V any](a, b iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		nextA, stopA := iter.Pull(a)
		defer stopA()
		nextB, stopB := iter.Pull(b)
		defer stopB()
		for {
			aa, okA := nextA()
			if !okA {
				return
			}
			bb, okB := nextB()
			if !okB {
				return
			}
			if !yield(aa) {
				return
			}
			if !yield(bb) {
				return
			}
		}
	}
}
