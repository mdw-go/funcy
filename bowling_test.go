package funcy

import (
	"testing"

	"github.com/mdwhatcott/funcy/should"
)

func TestBowling(t *testing.T) {
	should.So(t, Score(Repeat(20, 0)), should.Equal, 0)
	should.So(t, Score(Repeat(20, 1)), should.Equal, 20)
	should.So(t, Score([]int{5, 5, 2, 1}), should.Equal, 15)
	should.So(t, Score([]int{10, 3, 2, 1}), should.Equal, 21)
	should.So(t, Score(Repeat(12, 10)), should.Equal, 300)
}

func Score(rolls []int) (result int) {
	return Sum(Flatten(Take(10, toFrames(rolls))))
}
func toFrames(rolls []int) (result [][]int) {
	for len(rolls) > 0 {
		if isStrike(rolls) {
			result = append(result, Take(3, rolls))
			rolls = Rest(rolls)
		} else if isSpare(rolls) {
			result = append(result, Take(3, rolls))
			rolls = Drop(2, rolls)
		} else {
			result = append(result, Take(2, rolls))
			rolls = Drop(2, rolls)
		}
	}
	return result
}

func isSpare(rolls []int) bool  { return Sum(Take(2, rolls)) == 10 }
func isStrike(rolls []int) bool { return First(rolls) == 10 }
