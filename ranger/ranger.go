package ranger

import (
	"fmt"
	"iter"
	"math/rand/v2"
	"slices"

	"github.com/mdw-go/funcy/ranger/internal/ring"
	"github.com/mdw-go/funcy/ranger/is"
	"github.com/mdw-go/funcy/ranger/op"
)

func Complement[V any](predicate func(t V) bool) func(t V) bool {
	return func(t V) bool { return !predicate(t) }
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
func DoAll[V any](f func(V), seq iter.Seq[V]) {
	for s := range seq {
		f(s)
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
func DropLast[V any](n int, s iter.Seq[V]) iter.Seq[V] {
	return Map2(func(a, _ V) V { return a }, s, Drop(n, s))
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
func First[V any](s iter.Seq[V]) V {
	return Nth(0, s)
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
func Frequencies[V comparable](seq iter.Seq[V]) map[V]int {
	result := make(map[V]int)
	for s := range seq {
		result[s]++
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
func IndexBy[K comparable, V any](f func(V) K, seq iter.Seq[V]) map[K]V {
	result := make(map[K]V)
	for v := range seq {
		result[f(v)] = v
	}
	return result
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
func Iterator[S ~[]V, V any](s S) iter.Seq[V] {
	return slices.Values(s)
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
func MapPairs[M ~map[K]V, K comparable, V any](m M) iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for k, v := range m {
			if !yield(Pair[K, V]{A: k, B: v}) {
				return
			}
		}
	}
}
func Max[V is.Comparable](s iter.Seq[V]) (result V) {
	result = First(s)
	for s := range Rest(s) {
		if s > result {
			result = s
		}
	}
	return result
}
func Min[V is.Comparable](s iter.Seq[V]) (result V) {
	result = First(s)
	for s := range Rest(s) {
		if s < result {
			result = s
		}
	}
	return result
}
func Nest[V any](matrix [][]V) iter.Seq[iter.Seq[V]] {
	return func(yield func(iter.Seq[V]) bool) {
		for _, row := range matrix {
			_ = yield(Iterator(row))
		}
	}
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
func PairsMap[K comparable, V any](pairs iter.Seq[Pair[K, V]]) map[K]V {
	result := make(map[K]V)
	for pair := range pairs {
		result[pair.A] = pair.B
	}
	return result
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
func Product[N is.Number](seq iter.Seq[N]) N {
	return Reduce(op.Mul[N], N(1), seq)
}
func RandNth[V any](s iter.Seq[V]) V {
	return Nth(rand.N(Count(s)), s)
}
func Range[N is.Number](start, stop N) iter.Seq[N] {
	var step N = 1
	if stop < start {
		step = -step
	}
	return RangeStep(start, stop, step)
}
func RangeOpen[N is.Number](start, step N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for x := start; ; x += step {
			if !yield(x) {
				break
			}
		}
	}
}
func RangeStep[N is.Number](start, stop, step N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for x := start; x != stop; x += step {
			if !yield(x) {
				return
			}
		}
	}
}
func Reduce[V any](calc func(a, b V) V, start V, seq iter.Seq[V]) (result V) {
	result = start
	for s := range seq {
		result = calc(result, s)
	}
	return result
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
func Remove[V any](predicate func(V) bool, seq iter.Seq[V]) iter.Seq[V] {
	return Filter(Complement(predicate), seq)
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
func Rest[V any](s iter.Seq[V]) iter.Seq[V] {
	return Drop(1, s)
}
func Slice[V any](seq iter.Seq[V]) (result []V) {
	return slices.Collect(seq)
}
func Sum[N is.Number](seq iter.Seq[N]) N {
	return Reduce(op.Add[N], N(0), seq)
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
func Variadic[V any](vs ...V) iter.Seq[V] {
	return Iterator(vs)
}
func ZipMap[K comparable, V any](k iter.Seq[K], v iter.Seq[V]) map[K]V {
	nextA, stopA := iter.Pull(k)
	defer stopA()
	nextB, stopB := iter.Pull(v)
	defer stopB()
	result := make(map[K]V)
	for {
		aa, okA := nextA()
		if !okA {
			break
		}
		bb, okB := nextB()
		if !okB {
			break
		}
		result[aa] = bb
	}
	return result
}
func ZipPairs[A, B any](a iter.Seq[A], b iter.Seq[B]) iter.Seq[Pair[A, B]] {
	return func(yield func(Pair[A, B]) bool) {
		nextA, stopA := iter.Pull(a)
		defer stopA()
		nextB, stopB := iter.Pull(b)
		defer stopB()
		for {
			aa, okA := nextA()
			if !okA {
				break
			}
			bb, okB := nextB()
			if !okB {
				break
			}
			if !yield(Pair[A, B]{A: aa, B: bb}) {
				return
			}
		}
	}
}

type Pair[A, B any] struct {
	A A
	B B
}
