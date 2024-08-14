package examples

import (
	"testing"

	"github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/funcy/internal/should"
)

func TestBowling(t *testing.T) {
	should.So(t, Score(funcy.Repeat(20, 0)), should.Equal, 0)
	should.So(t, Score(funcy.Repeat(20, 1)), should.Equal, 20)
	should.So(t, Score([]int{5, 5, 2, 1}), should.Equal, 15)
	should.So(t, Score([]int{10, 3, 2, 1}), should.Equal, 21)
	should.So(t, Score(funcy.Repeat(12, 10)), should.Equal, 300)
}

func Score(rolls []int) (result int) {
	return funcy.Sum(funcy.Flatten(funcy.Take(10, toFrames(rolls, nil))))
}
func toFrames(rolls []int, result [][]int) [][]int {
	if len(rolls) == 0 {
		return result
	} else if isStrike(rolls) {
		return toFrames(funcy.Rest(rolls), append(result, funcy.Take(3, rolls)))
	} else if isSpare(rolls) {
		return toFrames(funcy.Drop(2, rolls), append(result, funcy.Take(3, rolls)))
	} else {
		return toFrames(funcy.Drop(2, rolls), append(result, funcy.Take(2, rolls)))
	}
}
func isSpare(rolls []int) bool  { return funcy.Sum(funcy.Take(2, rolls)) == 10 }
func isStrike(rolls []int) bool { return funcy.First(rolls) == 10 }
