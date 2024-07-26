package funcy

import (
	"slices"
	"testing"

	"github.com/mdwhatcott/funcy/v2/internal/should"
)

func TestTake(t *testing.T) {
	should.So(t, slices.Collect(Take(4, Range(0, 10))), should.Equal, []int{0, 1, 2, 3})
	should.So(t, slices.Collect(Take(8, Range(1, 5))), should.Equal, []int{1, 2, 3, 4})
	should.So(t, slices.Collect(Take(8, Range(1, 5))), should.Equal, []int{1, 2, 3, 4})
	should.So(t, slices.Collect(Take(1, Range(0, 0))), should.Equal, []int(nil))
}
func TestDrop(t *testing.T) {
	should.So(t, slices.Collect(Drop(4, Range(0, 10))), should.Equal, []int{4, 5, 6, 7, 8, 9})
	should.So(t, slices.Collect(Drop(8, Range(1, 5))), should.Equal, []int(nil))
	should.So(t, slices.Collect(Drop(10, Range(0, 0))), should.Equal, []int(nil))
	should.So(t, First(Drop(1, Range(1, 10))), should.Equal, 2)
	should.So(t, func() { First(Take(0, Range(0, 10))) }, should.Panic)
}
