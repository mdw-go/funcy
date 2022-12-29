package funcy_test

import (
	"fmt"
	"testing"

	"funcy"
	"funcy/should"
)

var (
	one2four = funcy.Range(1, 5)
	digits   = funcy.Range(0, 10)
)

func String[T any](t T) string { return fmt.Sprint(t) }
func Duplicate[T any](t T) []T { return []T{t, t} }
func LessThan[T funcy.Number](n T) func(T) bool {
	return func(t T) bool { return t < n }
}

func Test(t *testing.T) {
	should.So(t, funcy.Map(String[int], one2four), should.Equal, []string{"1", "2", "3", "4"})
	should.So(t, funcy.Map(funcy.Square[int], one2four), should.Equal, []int{1, 4, 9, 16})
	should.So(t, funcy.Filter(funcy.IsEven[int], one2four), should.Equal, []int{2, 4})
	should.So(t, funcy.Remove(funcy.IsOdd[int], one2four), should.Equal, []int{2, 4})
	should.So(t, funcy.Reduce(funcy.Add[int], 0, one2four), should.Equal, 10)
	should.So(t, funcy.MapCat(Duplicate[int], one2four), should.Equal, []int{1, 1, 2, 2, 3, 3, 4, 4})
	should.So(t, funcy.TakeWhile(LessThan(5), digits), should.Equal, []int{0, 1, 2, 3, 4})
	should.So(t, funcy.DropWhile(LessThan(5), digits), should.Equal, []int{5, 6, 7, 8, 9})
	should.So(t,
		funcy.Reduce(funcy.Add[int], 0,
			funcy.Filter(funcy.IsEven[int],
				funcy.Map(funcy.Square[int], digits))), should.Equal, 2*2+4*4+6*6+8*8)
}
