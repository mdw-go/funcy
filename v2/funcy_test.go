package funcy

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/funcy/v2/internal/should"
	"github.com/mdwhatcott/funcy/v2/is"
)

var (
	_12     = []int{1, 2}
	_123    = []int{1, 2, 3}
	_0123   = []int{0, 1, 2, 3}
	_1234   = []int{1, 2, 3, 4}
	_12345  = []int{1, 2, 3, 4, 5}
	_0246   = []int{0, 2, 4, 6}
	_135    = []int{1, 3, 5}
	_1357   = []int{1, 3, 5, 7}
	_456789 = []int{4, 5, 6, 7, 8, 9}
	_nil    = []int(nil)
)

func TestIterate(t *testing.T) {
	should.So(t, Slice(Iterate(_123)), should.Equal, _123)
	should.So(t, Slice(Take(2, Iterate(_123))), should.Equal, _12)
	should.So(t, Slice(Variadic(1, 2, 3)), should.Equal, _123)
}
func TestTake(t *testing.T) {
	should.So(t, Slice(Take(4, Range(0, 10))), should.Equal, _0123)
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, _1234)
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, _1234)
	should.So(t, Slice(Take(1, Range(0, 0))), should.Equal, _nil)
}
func TestTakeWhile(t *testing.T) {
	should.So(t, Slice(Take(4, TakeWhile(is.Even[int], Iterate([]int{0, 2, 4, 6, 8, 1, 3, 5, 7})))), should.Equal, _0246)
	should.So(t, Slice(TakeWhile(is.Even[int], Iterate([]int{1, 3, 5, 7, 0, 2, 4, 6, 8}))), should.Equal, _nil)
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
	should.So(t, Slice(Take(3, DropWhile(is.Even[int], Iterate([]int{0, 2, 4, 6, 8, 1, 3, 5, 7})))), should.Equal, _135)
	should.So(t, Slice(DropWhile(is.Even[int], Range(1, 10))), should.Equal, Slice(Range(1, 10)))
}
func TestDropLast(t *testing.T) {
	should.So(t, Slice(Take(5, DropLast(4, Range(1, 11)))), should.Equal, _12345)
}
func TestFirst(t *testing.T) {
	should.So(t, func() { First(Take(0, Range(0, 10))) }, should.Panic)
	should.So(t, First(Drop(1, Range(1, 10))), should.Equal, 2)
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
	should.So(t, Slice(Map(upper, Iterate([]rune("asdf")))), should.Equal, []string{"A", "S", "D", "F"})
}
func TestMap2(t *testing.T) {
	add := func(a int, b int) int { return a + b }
	should.So(t, Slice(Take(5, Map2(add, Range(0, 10), Repeat(10, 1)))), should.Equal, _12345)
	should.So(t, Slice(Map2(add, Range(0, 1), Repeat(10, 1))), should.Equal, []int{1})
	should.So(t, Slice(Map2(add, Range(0, 10), Repeat(2, 1))), should.Equal, []int{1, 2})
	should.So(t, Slice(Map2(add, Range(0, 0), Repeat(1, 1))), should.Equal, []int(nil))
	should.So(t, Slice(Map2(add, Range(0, 1), Repeat(0, 1))), should.Equal, []int(nil))
}
func TestReduce(t *testing.T) {
	add := func(a, b int) int { return a + b }
	should.So(t, Reduce(add, 0, Range(1, 6)), should.Equal, 15)
}
func TestRepeat(t *testing.T) {
	should.So(t,
		Slice(Take(10, Repeat(20, 1))), should.Equal,
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
func TestSum(t *testing.T) {
	should.So(t, Sum(Range(1, 6)), should.Equal, 15)
}
