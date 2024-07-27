package funcy

import (
	"testing"

	"github.com/mdwhatcott/funcy/v2/internal/should"
)

func TestSeq(t *testing.T) {
	should.So(t, Slice(Seq([]int{1, 2, 3})), should.Equal, []int{1, 2, 3})
	should.So(t, Slice2(Seq2([]int{1, 2, 3})), should.Equal, []int{1, 2, 3})
}
func TestTake(t *testing.T) {
	should.So(t, Slice(Take(4, Range(0, 10))), should.Equal, []int{0, 1, 2, 3})
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, []int{1, 2, 3, 4})
	should.So(t, Slice(Take(8, Range(1, 5))), should.Equal, []int{1, 2, 3, 4})
	should.So(t, Slice(Take(1, Range(0, 0))), should.Equal, []int(nil))
}
func TestDrop(t *testing.T) {
	should.So(t, Slice(Drop(4, Range(0, 10))), should.Equal, []int{4, 5, 6, 7, 8, 9})
	should.So(t, Slice(Drop(8, Range(1, 5))), should.Equal, []int(nil))
	should.So(t, Slice(Drop(10, Range(0, 0))), should.Equal, []int(nil))
	should.So(t, Slice(Drop(1, Range(0, 0))), should.Equal, []int(nil))
}
func TestFirst(t *testing.T) {
	should.So(t, func() { First(Take(0, Range(0, 10))) }, should.Panic)
	should.So(t, First(Drop(1, Range(1, 10))), should.Equal, 2)
}
