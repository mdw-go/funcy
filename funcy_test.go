package funcy

import (
	"testing"

	"github.com/mdwhatcott/funcy/internal/should"
)

var (
	one2four = Range(1, 5)
	digits   = Range(0, 10)
	reversed = Range(9, -1)
)

func duplicate[T any](t T) []T { return []T{t, t} }
func isLessThan[T Number](n T) func(T) bool {
	return func(t T) bool { return t < n }
}

func TestData(t *testing.T) {
	should.So(t, one2four, should.Equal, []int{1, 2, 3, 4})
	should.So(t, digits, should.Equal, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	should.So(t, reversed, should.Equal, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0})
}
func TestMap(t *testing.T) {
	should.So(t, Map(String[int], one2four), should.Equal, []string{"1", "2", "3", "4"})
	should.So(t, Map(Square[int], one2four), should.Equal, []int{1, 4, 9, 16})
	should.So(t, Map2(Add[int], digits, reversed), should.Equal, Repeat(10, 9))
	should.So(t, Map2(Add[int], digits, one2four), should.Equal, []int{1, 3, 5, 7})
	should.So(t, MapCat(duplicate[int], one2four), should.Equal, []int{1, 1, 2, 2, 3, 3, 4, 4})
}
func TestFilters(t *testing.T) {
	should.So(t, Filter(IsEven[int], one2four), should.Equal, []int{2, 4})
	should.So(t, Remove(IsOdd[int], one2four), should.Equal, []int{2, 4})
	should.So(t, FilterAs[int]([]any{1, "two", 3, "four", 5}), should.Equal, []int{1, 3, 5})
	should.So(t, Filter(IsEven[int], digits), should.Equal, RangeStep(0, 10, 2))
}
func TestReduce(t *testing.T) {
	should.So(t, Reduce(Add[int], 0, one2four), should.Equal, 10)
	should.So(t, Sum(one2four), should.Equal, 10)
	should.So(t, Product(one2four), should.Equal, 24)
}
func TestMinMax(t *testing.T) {
	should.So(t, Max(digits), should.Equal, 9)
	should.So(t, Min(reversed), should.Equal, 0)
}
func TestCombinations(t *testing.T) {
	should.So(t, Sum(Filter(IsEven[int], Map(Square[int], digits))), should.Equal, 2*2+4*4+6*6+8*8)
}
func TestTakeDropEtc(t *testing.T) {
	should.So(t, Take(4, Drop(1, digits)), should.Equal, one2four)
	should.So(t, Take(20, digits), should.Equal, digits)
	should.So(t, Drop(20, digits), should.BeEmpty)
	should.So(t, First(digits), should.Equal, 0)
	should.So(t, Last(digits), should.Equal, 9)
	should.So(t, Rest(one2four), should.Equal, []int{2, 3, 4})
	should.So(t, TakeAllBut(2, digits), should.Equal, []int{0, 1, 2, 3, 4, 5, 6, 7})
	should.So(t, DropAllBut(2, digits), should.Equal, []int{8, 9})
	should.So(t, TakeWhile(isLessThan(5), digits), should.Equal, []int{0, 1, 2, 3, 4})
	should.So(t, DropWhile(isLessThan(5), digits), should.Equal, []int{5, 6, 7, 8, 9})
}
func TestIndexing(t *testing.T) {
	should.So(t, IndexBy(ByLength[string], []string{"a", "ab", "c", "abc"}),
		should.Equal, map[int]string{1: "c", 2: "ab", 3: "abc"})
	should.So(t, SlicedIndexBy(ByLength[string], []string{"a", "ab", "c", "abc"}),
		should.Equal, map[int][]string{1: {"a", "c"}, 2: {"ab"}, 3: {"abc"}})
}
func TestSorting(t *testing.T) {
	should.So(t, SortDescending(ByLength[[]string], GroupBy(ByLength[string], []string{"a", "b", "c", "ab", "bc", "abc"})),
		should.Equal, [][]string{{"a", "b", "c"}, {"ab", "bc"}, {"abc"}})
	should.So(t, SortAscending(func(i int) int { return i }, reversed), should.Equal, digits)
	should.So(t, SortDescending(func(i int) int { return i }, digits), should.Equal, reversed)
}
func TestZipping(t *testing.T) {
	should.So(t, Zip([]int{1, 2, 3, 4}, []rune{'a', 'b', 'c'}), should.Equal,
		[]Pair[int, rune]{{A: 1, B: 'a'}, {A: 2, B: 'b'}, {A: 3, B: 'c'}})
	should.So(t, ZipMap([]int{1, 2, 3, 4}, []rune{'a', 'b', 'c'}), should.Equal, map[int]rune{1: 'a', 2: 'b', 3: 'c'})
}
func TestFrequencies(t *testing.T) {
	should.So(t, Frequencies([]rune{'a', 'b', 'c', 'b', 'a', 'a'}), should.Equal, map[rune]int{'a': 3, 'b': 2, 'c': 1})
}
func TestFlattenPartition(t *testing.T) {
	segments := [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}}
	should.So(t, Flatten(segments), should.Equal, []int{0, 1, 2, 3, 4, 5, 6, 7, 8})
	should.So(t, Partition(3, 3, digits), should.Equal, segments)
	should.So(t, Partition(3, 1, one2four), should.Equal, [][]int{{1, 2, 3}, {2, 3, 4}})
	should.So(t, Partition(100, 1, one2four), should.BeNil)
	should.So(t, Partition(1, 100, one2four), should.Equal, [][]int{{1}})
}
func TestLoadDrain(t *testing.T) {
	channel := make(chan int)
	go Load(channel, digits)
	digits2 := Drain(channel)
	should.So(t, digits2, should.Equal, digits)
}
func TestAnyAllNone(t *testing.T) {
	should.So(t, Any([]bool{false, false, false}), should.BeFalse)
	should.So(t, Any([]bool{false, false, true}), should.BeTrue)

	should.So(t, All([]bool{true, true, true}), should.BeTrue)
	should.So(t, All([]bool{true, true, false}), should.BeFalse)

	should.So(t, None([]bool{false, false, false}), should.BeTrue)
	should.So(t, None([]bool{false, false, true}), should.BeFalse)
}
func TestMapKeysValues(t *testing.T) {
	should.So(t, SortAscending(ByNumericValue[int], MapKeys(map[int]string{1: "a", 2: "b", 3: "c"})), should.Equal, []int{1, 2, 3})
	should.So(t, SortAscending(ByLength[string], MapValues(map[int]string{1: "a", 2: "bb", 3: "ccc"})), should.Equal, []string{"a", "bb", "ccc"})
}
