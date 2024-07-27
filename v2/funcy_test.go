package funcy

import (
	"testing"

	"github.com/mdwhatcott/funcy/v2/internal/should"
)

var (
	_12     = []int{1, 2}
	_123    = []int{1, 2, 3}
	_0123   = []int{0, 1, 2, 3}
	_1234   = []int{1, 2, 3, 4}
	_456789 = []int{4, 5, 6, 7, 8, 9}
	_nil    = []int(nil)
)

func TestSeq(t *testing.T) {
	should.So(t, Slice(Seq(_123)), should.Equal, _123)
	should.So(t, Slice2(Seq2(_123)), should.Equal, _123)
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
