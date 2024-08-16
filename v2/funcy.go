package funcy

import (
	"fmt"
	"iter"
	"math/rand/v2"
	"slices"

	"github.com/mdwhatcott/funcy/v2/internal/ring"
	"github.com/mdwhatcott/funcy/v2/is"
)

/*
TODO:
- https://clojuredocs.org/clojure.core/sort-by
- https://clojuredocs.org/clojure.core/zipmap
*/

func Variadic[V any](vs ...V) iter.Seq[V] {
	return Iterator(vs)
}
func Iterator[S ~[]V, V any](s S) iter.Seq[V] {
	return slices.Values(s)
}
func Slice[V any](seq iter.Seq[V]) (result []V) {
	return slices.Collect(seq)
}
func Range[N is.Number](start, stop N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for x := start; x < stop; x++ {
			if !yield(x) {
				return
			}
		}
	}
}
func First[V any](s iter.Seq[V]) V {
	return Nth(0, s)
}
func Nth[V any](n int, s iter.Seq[V]) V {
	if n < 0 {
		panic(fmt.Sprintf("runtime error: index out of range [%d]", n))
	}
	c := 0
	for v := range s {
		if c == n {
			return v
		}
		c++
	}
	panic(fmt.Sprintf("runtime error: index out of range [%d] with length %d", n, c))
}
func RandNth[V any](s iter.Seq[V]) V {
	return Nth(rand.N(Count(s)), s)
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
func DoAll[V any](f func(V), seq iter.Seq[V]) {
	for s := range seq {
		f(s)
	}
}
func Reduce[V any](calc func(a, b V) V, start V, seq iter.Seq[V]) (result V) {
	return Last(Reductions(calc, start, seq))
}
func Reductions[V any](calc func(a, b V) V, start V, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		result := start
		for next := range seq {
			result = calc(result, next)
			if !yield(result) {
				return
			}
		}
	}
}
func Repeat[V any](v V) iter.Seq[V] {
	return Repeatedly(func() V { return v })
}
func RepeatN[V any](n int, v V) iter.Seq[V] {
	return Take(n, Repeatedly(func() V { return v }))
}
func Repeatedly[V any](v func() V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for yield(v()) {
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
			_ = yield(Iterator(row))
		}
	}
}
func Partition[V any](chunkLength, stride int, seq iter.Seq[V]) iter.Seq[iter.Seq[V]] {
	return func(yield func(iter.Seq[V]) bool) {
		for {
			chunk := Take(chunkLength, seq)
			if Count(chunk) < chunkLength {
				return
			}
			if !yield(chunk) {
				return
			}
			seq = Drop(stride, seq)
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
func Interpose[V any](sep V, seq iter.Seq[V]) iter.Seq[V] {
	return Drop(1, Interleave(Repeat(sep), seq))
}
func Iterate[V any](f func(V) V, v V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			v = f(v)
			if !yield(v) {
				return
			}
		}
	}
}
func Frequencies[V comparable](seq iter.Seq[V]) map[V]int {
	result := make(map[V]int)
	for s := range seq {
		result[s]++
	}
	return result
}
func IndexBy[K comparable, V any](f func(V) K, seq iter.Seq[V]) map[K]V {
	result := make(map[K]V)
	for v := range seq {
		result[f(v)] = v
	}
	return result
}
func GroupBy[K comparable, V any](f func(V) K, seq iter.Seq[V]) map[K][]V {
	result := make(map[K][]V)
	for v := range seq {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}
