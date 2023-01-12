package funcy

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/funcy/should"
)

var (
	one2four = Range(1, 5)
	digits   = Range(0, 10)
	reversed = Range(9, -1)
)

func Square[T Number](t T) T                { return t * t }
func IsEven[T Integer](t T) bool            { return t%2 == 0 }
func IsOdd[T Integer](t T) bool             { return t%2 == 1 }
func String[T any](t T) string              { return fmt.Sprint(t) }
func Duplicate[T any](t T) []T              { return []T{t, t} }
func byLength[S string | []string](s S) int { return len(s) }
func isLessThan[T Number](n T) func(T) bool {
	return func(t T) bool { return t < n }
}

func Test(t *testing.T) {
	should.So(t, reversed, should.Equal, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
	should.So(t, Map(String[int], one2four), should.Equal, []string{"1", "2", "3", "4"})
	should.So(t, Map(Square[int], one2four), should.Equal, []int{1, 4, 9, 16})
	should.So(t, Filter(IsEven[int], one2four), should.Equal, []int{2, 4})
	should.So(t, Remove(IsOdd[int], one2four), should.Equal, []int{2, 4})
	should.So(t, Reduce(Add[int], 0, one2four), should.Equal, 10)
	should.So(t, Sum(one2four), should.Equal, 10)
	should.So(t, Product(one2four), should.Equal, 24)
	should.So(t, MapCat(Duplicate[int], one2four), should.Equal, []int{1, 1, 2, 2, 3, 3, 4, 4})
	should.So(t, Take(4, Drop(1, digits)), should.Equal, one2four)
	should.So(t, Take(20, digits), should.Equal, digits)
	should.So(t, Drop(20, digits), should.BeEmpty)
	should.So(t, First(digits), should.Equal, 0)
	should.So(t, Last(digits), should.Equal, 9)
	should.So(t, Rest(one2four), should.Equal, []int{2, 3, 4})
	should.So(t, AllBut(2, digits), should.Equal, []int{8, 9})
	should.So(t, TakeWhile(isLessThan(5), digits), should.Equal, []int{0, 1, 2, 3, 4})
	should.So(t, DropWhile(isLessThan(5), digits), should.Equal, []int{5, 6, 7, 8, 9})
	should.So(t, IndexBy(byLength[string], []string{"a", "ab", "c", "abc"}),
		should.Equal, map[int]string{1: "c", 2: "ab", 3: "abc"})
	should.So(t, SlicedIndexBy(byLength[string], []string{"a", "ab", "c", "abc"}),
		should.Equal, map[int][]string{1: {"a", "c"}, 2: {"ab"}, 3: {"abc"}})
	should.So(t, SortDescending(byLength[[]string], GroupBy(byLength[string], []string{"a", "b", "c", "ab", "bc", "abc"})),
		should.Equal, [][]string{{"a", "b", "c"}, {"ab", "bc"}, {"abc"}})
	should.So(t, FilterAs[int]([]any{1, "two", 3, "four", 5}), should.Equal, []int{1, 3, 5})
	should.So(t, SortAscending(func(i int) int { return i }, reversed), should.Equal, digits)
	should.So(t, SortDescending(func(i int) int { return i }, digits), should.Equal, reversed)
	should.So(t, Zip([]int{1, 2, 3, 4}, []rune{'a', 'b', 'c'}), should.Equal,
		[]Pair[int, rune]{{A: 1, B: 'a'}, {A: 2, B: 'b'}, {A: 3, B: 'c'}})
	should.So(t, Frequencies([]rune{'a', 'b', 'c', 'b', 'a', 'a'}), should.Equal, map[rune]int{
		'a': 3,
		'b': 2,
		'c': 1,
	})
	should.So(t, Flatten([][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8, 9}}), should.Equal, digits)
	should.So(t,
		Reduce(Add[int], 0,
			Filter(IsEven[int],
				Map(Square[int], digits))), should.Equal, 2*2+4*4+6*6+8*8)

	channel := make(chan int)
	go Load(channel, digits)
	digits2 := Drain(channel)
	should.So(t, digits2, should.Equal, digits)
}
