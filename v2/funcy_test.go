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
	_0246   = []int{0, 2, 4, 6}
	_1357   = []int{1, 3, 5, 7}
	_456789 = []int{4, 5, 6, 7, 8, 9}
	_nil    = []int(nil)
)

func TestSeq(t *testing.T) {
	should.So(t, Slice(Seq(_123)), should.Equal, _123)
	should.So(t, Slice(Take(2, Seq(_123))), should.Equal, _12)
	should.So(t, Slice(Take(2, Seq2Seq(Seq2(_123)))), should.Equal, _12)
}
func TestTake(t *testing.T) {
	should.So(t, Slice(Take(4, Range(0, 10))), should.Equal, _0123)
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, _1234)
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, _1234)
	should.So(t, Slice(Take(1, Range(0, 0))), should.Equal, _nil)
}
func TestDrop(t *testing.T) {
	should.So(t, Slice(Drop(4, Range(0, 10))), should.Equal, _456789)
	should.So(t, Slice(Drop(8, Range(1, 5))), should.Equal, _nil)
	should.So(t, Slice(Drop(10, Range(0, 0))), should.Equal, _nil)
	should.So(t, Slice(Drop(1, Range(0, 0))), should.Equal, _nil)
}
func TestFirst(t *testing.T) {
	should.So(t, func() { First(Take(0, Range(0, 10))) }, should.Panic)
	should.So(t, First(Drop(1, Range(1, 10))), should.Equal, 2)
}
func TestLast(t *testing.T) {
	should.So(t, func() { Last(Take(0, Range(0, 10))) }, should.Panic)
	should.So(t, Last(Take(3, Range(1, 10))), should.Equal, 3)
}
func TestFilter(t *testing.T) {
	should.So(t, Slice(Take(4, Filter(is.Even[int], Range(0, 10)))), should.Equal, _0246)
	should.So(t, Slice(Take(4, Remove(is.Even[int], Range(0, 10)))), should.Equal, _1357)
}
func TestMap(t *testing.T) {
	square := func(n int) int64 { return int64(n * n) }
	upper := func(r rune) string { return strings.ToUpper(string(r)) }
	should.So(t, Slice(Take(5, Map(square, Range(2, 10)))), should.Equal, []int64{4, 9, 16, 25, 36})
	should.So(t, Slice(Map(upper, Seq([]rune("asdf")))), should.Equal, []string{"A", "S", "D", "F"})
}
func TestReduce(t *testing.T) {
	add := func(a, b int) int { return a + b }
	should.So(t, Reduce(add, 0, Range(1, 6)), should.Equal, 15)
}
