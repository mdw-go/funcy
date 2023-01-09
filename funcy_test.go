package funcy_test

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/funcy/should"
)

var (
	one2four = funcy.Range(1, 5)
	digits   = funcy.Range(0, 10)
	reversed = funcy.Range(9, -1)
)

func Add[T funcy.Number](a, b T) T     { return a + b }
func Square[T funcy.Number](t T) T     { return t * t }
func IsEven[T funcy.Integer](t T) bool { return t%2 == 0 }
func IsOdd[T funcy.Integer](t T) bool  { return t%2 == 1 }
func String[T any](t T) string         { return fmt.Sprint(t) }
func Duplicate[T any](t T) []T         { return []T{t, t} }
func LessThan[T funcy.Number](n T) func(T) bool {
	return func(t T) bool { return t < n }
}
func byLength(s string) int { return len(s) }

func Test(t *testing.T) {
	should.So(t, reversed, should.Equal, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	should.So(t, funcy.Map(String[int], one2four), should.Equal, []string{"1", "2", "3", "4"})
	should.So(t, funcy.Map(Square[int], one2four), should.Equal, []int{1, 4, 9, 16})
	should.So(t, funcy.Filter(IsEven[int], one2four), should.Equal, []int{2, 4})
	should.So(t, funcy.Remove(IsOdd[int], one2four), should.Equal, []int{2, 4})
	should.So(t, funcy.Reduce(Add[int], 0, one2four), should.Equal, 10)
	should.So(t, funcy.MapCat(Duplicate[int], one2four), should.Equal, []int{1, 1, 2, 2, 3, 3, 4, 4})
	should.So(t, funcy.Take(4, funcy.Drop(1, digits)), should.Equal, one2four)
	should.So(t, funcy.Rest(one2four), should.Equal, []int{2, 3, 4})
	should.So(t, funcy.AllBut(2, digits), should.Equal, []int{8, 9})
	should.So(t, funcy.TakeWhile(LessThan(5), digits), should.Equal, []int{0, 1, 2, 3, 4})
	should.So(t, funcy.DropWhile(LessThan(5), digits), should.Equal, []int{5, 6, 7, 8, 9})
	should.So(t, funcy.IndexBy(byLength, []string{"a", "ab", "c", "abc"}),
		should.Equal, map[int]string{1: "c", 2: "ab", 3: "abc"})
	should.So(t, funcy.SlicedIndexBy(byLength, []string{"a", "ab", "c", "abc"}),
		should.Equal, map[int][]string{1: {"a", "c"}, 2: {"ab"}, 3: {"abc"}})
	should.So(t, funcy.FilterAs[int]([]any{1, "two", 3, "four", 5}), should.Equal, []int{1, 3, 5})
	should.So(t, funcy.SortAscending(func(i int) int { return i }, reversed), should.Equal, digits)
	should.So(t, funcy.SortDescending(func(i int) int { return i }, digits), should.Equal, reversed)
	should.So(t, funcy.Zip([]int{1, 2, 3, 4}, []rune{'a', 'b', 'c'}), should.Equal,
		[]funcy.Pair[int, rune]{{A: 1, B: 'a'}, {A: 2, B: 'b'}, {A: 3, B: 'c'}})
	should.So(t,
		funcy.Reduce(Add[int], 0,
			funcy.Filter(IsEven[int],
				funcy.Map(Square[int], digits))), should.Equal, 2*2+4*4+6*6+8*8)

	channel := make(chan int)
	go funcy.Load(channel, digits)
	digits2 := funcy.Drain(channel)
	should.So(t, digits2, should.Equal, digits)
}
