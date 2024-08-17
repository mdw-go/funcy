package ranger

import (
	"strconv"
	"strings"
	"testing"

	"github.com/mdwhatcott/funcy/ranger/internal/should"
	"github.com/mdwhatcott/funcy/ranger/is"
)

var (
	_12        = []int{1, 2}
	_123       = []int{1, 2, 3}
	_0123      = []int{0, 1, 2, 3}
	_1234      = []int{1, 2, 3, 4}
	_12345     = []int{1, 2, 3, 4, 5}
	_0246      = []int{0, 2, 4, 6}
	_135       = []int{1, 3, 5}
	_1357      = []int{1, 3, 5, 7}
	_456789    = []int{4, 5, 6, 7, 8, 9}
	_123456789 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	_987654321 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	_nil       = []int(nil)
)

func TestRange(t *testing.T) {
	should.So(t, Slice(Range(1, 10)), should.Equal, _123456789)
	should.So(t, Slice(Range(9, 0)), should.Equal, _987654321)
	should.So(t, Slice(RangeStep(9, 0, -1)), should.Equal, _987654321)
}
func TestIterator(t *testing.T) {
	should.So(t, Slice(Iterator(_123)), should.Equal, _123)
	should.So(t, Slice(Take(2, Iterator(_123))), should.Equal, _12)
	should.So(t, Slice(Variadic(1, 2, 3)), should.Equal, _123)
}
func TestTake(t *testing.T) {
	should.So(t, Slice(Take(4, Range(0, 10))), should.Equal, _0123)
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, _1234)
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, _1234)
	should.So(t, Slice(Take(1, Range(0, 0))), should.Equal, _nil)
	should.So(t, Slice(Take(2, Take(3, Range(1, 10)))), should.Equal, _12)
}
func TestTakeWhile(t *testing.T) {
	should.So(t, Slice(Take(4, TakeWhile(is.Even[int], Iterator([]int{0, 2, 4, 6, 8, 1, 3, 5, 7})))), should.Equal, _0246)
	should.So(t, Slice(TakeWhile(is.Even[int], Iterator([]int{1, 3, 5, 7, 0, 2, 4, 6, 8}))), should.Equal, _nil)
}
func TestTakeLast(t *testing.T) {
	should.So(t, Slice(TakeLast(20, Range(0, 10))), should.Equal, Slice(Range(0, 10)))
	should.So(t, Slice(Take(4, TakeLast(5, Range(0, 10)))), should.Equal, []int{5, 6, 7, 8})
}
func TestDrop(t *testing.T) {
	should.So(t, Slice(Drop(4, Range(0, 10))), should.Equal, _456789)
	should.So(t, Slice(Drop(8, Range(1, 5))), should.Equal, _nil)
	should.So(t, Slice(Drop(10, Range(0, 0))), should.Equal, _nil)
	should.So(t, Slice(Drop(1, Range(0, 0))), should.Equal, _nil)
}
func TestDropWhile(t *testing.T) {
	should.So(t, Slice(Take(3, DropWhile(is.Even[int], Iterator([]int{0, 2, 4, 6, 8, 1, 3, 5, 7})))), should.Equal, _135)
	should.So(t, Slice(DropWhile(is.Even[int], Range(1, 10))), should.Equal, Slice(Range(1, 10)))
}
func TestDropLast(t *testing.T) {
	should.So(t, Slice(Take(5, DropLast(4, Range(1, 11)))), should.Equal, _12345)
}
func TestFirst(t *testing.T) {
	should.So(t, func() { First(Take(0, Range(0, 10))) }, should.Panic)
	should.So(t, First(Drop(1, Range(1, 10))), should.Equal, 2)
}
func TestNth(t *testing.T) {
	should.So(t, func() { Nth(-1, Iterator(_1234)) }, should.Panic)
	should.So(t, Nth(2, Iterator(_1234)), should.Equal, 3)
}
func TestRandNth(t *testing.T) {
	should.So(t, RandNth(Variadic(42)), should.Equal, 42)
	var values []int
	for range 100 {
		nth := RandNth(Iterator(_123456789))
		values = append(values, nth)
		should.So(t, nth, should.BeGreaterThan, 0)
		should.So(t, nth, should.BeLessThan, 10)
	}
	freq := Frequencies(Iterator(values))
	for _, n := range _123456789 {
		should.So(t, freq[n], should.BeGreaterThan, 0)
	}
}
func TestLast(t *testing.T) {
	should.So(t, func() { Last(Take(0, Range(0, 10))) }, should.Panic)
	should.So(t, Last(Take(3, Range(1, 10))), should.Equal, 3)
}
func TestRest(t *testing.T) {
	should.So(t, Slice(Take(3, Rest(Range(1, 10)))), should.Equal, Slice(Range(2, 5)))
	should.So(t, Slice(Take(3, Rest(Range(0, 0)))), should.Equal, _nil)
}
func TestFilter(t *testing.T) {
	should.So(t, Slice(Take(4, Filter(is.Even[int], Range(0, 10)))), should.Equal, _0246)
	should.So(t, Slice(Take(4, Remove(is.Even[int], Range(0, 10)))), should.Equal, _1357)
}
func TestMap(t *testing.T) {
	square := func(n int) int64 { return int64(n * n) }
	upper := func(r rune) string { return strings.ToUpper(string(r)) }
	should.So(t, Slice(Take(5, Map(square, Range(2, 10)))), should.Equal, []int64{4, 9, 16, 25, 36})
	should.So(t, Slice(Map(upper, Iterator([]rune("asdf")))), should.Equal, []string{"A", "S", "D", "F"})
}
func TestMap2(t *testing.T) {
	add := func(a int, b int) int { return a + b }
	should.So(t, Slice(Take(5, Map2(add, Range(0, 10), RepeatN(10, 1)))), should.Equal, _12345)
	should.So(t, Slice(Map2(add, Range(0, 1), RepeatN(10, 1))), should.Equal, []int{1})
	should.So(t, Slice(Map2(add, Range(0, 10), RepeatN(2, 1))), should.Equal, []int{1, 2})
	should.So(t, Slice(Map2(add, Range(0, 0), RepeatN(1, 1))), should.Equal, []int(nil))
	should.So(t, Slice(Map2(add, Range(0, 1), RepeatN(0, 1))), should.Equal, []int(nil))
}
func TestReduce(t *testing.T) {
	add := func(a, b int) int { return a + b }
	should.So(t, Reduce(add, 0, Range(1, 6)), should.Equal, 15)
}
func TestRepeat(t *testing.T) {
	should.So(t,
		Slice(Take(10, RepeatN(20, 1))), should.Equal,
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	)
}
func TestConcat(t *testing.T) {
	should.So(t,
		Slice(Take(12, Concat(Range(0, 5), Range(5, 10), Range(10, 15)))), should.Equal,
		Slice(Range(0, 12)),
	)
}
func TestFlatten(t *testing.T) {
	should.So(t, Slice(Take(12, Flatten(Nest([][]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
	})))), should.Equal, Slice(Range(0, 12)))
}
func TestPartition(t *testing.T) {
	should.So(t, Slice(Map(Sum[int], Partition(3, 3, Range(1, 10)))), should.Equal, []int{6, 15, 24})
	should.So(t, Slice(Map(Sum[int], Take(2, Partition(3, 3, Range(1, 10))))), should.Equal, []int{6, 15})
	should.So(t, Slice(Map(Sum[int], Partition(3, 1, Range(1, 5)))), should.Equal, []int{6, 9})
}
func TestSum(t *testing.T) {
	should.So(t, Sum(Range(1, 6)), should.Equal, 1+2+3+4+5)
}
func TestProduct(t *testing.T) {
	should.So(t, Product(Range(1, 6)), should.Equal, 1*2*3*4*5)
}
func TestCount(t *testing.T) {
	should.So(t, Count(Range(0, 20)), should.Equal, 20)
}
func TestCycle(t *testing.T) {
	should.So(t, Slice(Take(9, Cycle(Range(0, 2)))), should.Equal, []int{0, 1, 0, 1, 0, 1, 0, 1, 0})
}
func TestInterleave(t *testing.T) {
	should.So(t, Slice(Take(5, Interleave(Range(0, 10), Range(10, 20)))), should.Equal, []int{0, 10, 1, 11, 2})
	should.So(t, Slice(Take(6, Interleave(Range(0, 10), Range(10, 20)))), should.Equal, []int{0, 10, 1, 11, 2, 12})
	should.So(t, Slice(Interleave(Range(0, 0), Range(0, 10))), should.Equal, _nil)
	should.So(t, Slice(Interleave(Range(0, 10), Range(0, 0))), should.Equal, _nil)
}
func TestInterpose(t *testing.T) {
	should.So(t, Slice(Interpose(-1, Range(0, 5))), should.Equal, []int{
		0, -1,
		1, -1,
		2, -1,
		3, -1,
		4,
	})
	should.So(t, Slice(Take(4, Interpose(-1, Range(0, 5)))), should.Equal, []int{
		0, -1,
		1, -1,
	})
}
func TestRepeatedly(t *testing.T) {
	should.So(t, Slice(Take(5, Repeatedly(func() int { return 3 }))), should.Equal, []int{3, 3, 3, 3, 3})
}
func TestReductions(t *testing.T) {
	add := func(a int, b int) int { return a + b }
	should.So(t, Slice(Take(5, Reductions(add, 0, Range(1, 10)))), should.Equal, []int{1, 3, 6, 10, 15})
	should.So(t, Slice(Reductions(add, 0, Range(1, 6))), should.Equal, []int{1, 3, 6, 10, 15})
	should.So(t, Last(Reductions(add, 0, Range(1, 6))), should.Equal, Reduce(add, 0, Range(1, 6)))
}
func TestIterate(t *testing.T) {
	inc := func(a int) int { return a + 1 }
	should.So(t, Slice(Take(5, Iterate(inc, 0))), should.Equal, Slice(Range(1, 6)))
}
func TestFrequencies(t *testing.T) {
	should.So(t, Frequencies(Variadic(1, 1, 2, 2, 2, 3, 4, 4)), should.Equal, map[int]int{
		1: 2,
		2: 3,
		3: 1,
		4: 2,
	})
}
func TestDoAll(t *testing.T) {
	var all []int
	store := func(a int) {
		all = append(all, a)
	}
	DoAll(store, Range(1, 10))
	should.So(t, all, should.Equal, _123456789)
}
func TestIndexBy(t *testing.T) {
	should.So(t, IndexBy(strconv.Itoa, Range(0, 5)), should.Equal,
		map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4},
	)
}
func TestGroupBy(t *testing.T) {
	should.So(t, GroupBy(strconv.Itoa, Concat(Range(0, 5), Range(1, 4))), should.Equal,
		map[string][]int{
			"0": {0},
			"1": {1, 1},
			"2": {2, 2},
			"3": {3, 3},
			"4": {4},
		},
	)
}
func TestZipMap(t *testing.T) {
	should.So(t, ZipMap(Range(0, 5), Range(10, 15)), should.Equal, map[int]int{
		0: 10,
		1: 11,
		2: 12,
		3: 13,
		4: 14,
	})
	should.So(t, ZipMap(Range(0, 6), Range(10, 15)), should.Equal, map[int]int{
		0: 10,
		1: 11,
		2: 12,
		3: 13,
		4: 14,
	})
	should.So(t, ZipMap(Range(0, 5), Range(10, 16)), should.Equal, map[int]int{
		0: 10,
		1: 11,
		2: 12,
		3: 13,
		4: 14,
	})
}
func TestMin(t *testing.T) {
	should.So(t, Min(Range(4, 20)), should.Equal, 4)
	should.So(t, func() { Min(Range(0, 0)) }, should.Panic)
	should.So(t, Min(Variadic(1, 6, -2, 3, 42)), should.Equal, -2)
}
func TestMax(t *testing.T) {
	should.So(t, Max(Range(4, 20)), should.Equal, 19)
	should.So(t, func() { Max(Range(0, 0)) }, should.Panic)
	should.So(t, Max(Variadic(1, 6, -2, 3, 42, 7)), should.Equal, 42)
}
