package examples

import (
	"iter"
	"testing"

	. "github.com/mdwhatcott/funcy/ranger"
	"github.com/mdwhatcott/funcy/ranger/internal/should"
)

func TestBowling(t *testing.T) {
	should.So(t, Score(RepeatN(20, 0)), should.Equal, 0)
	should.So(t, Score(RepeatN(20, 1)), should.Equal, 20)
	should.So(t, Score(finishGutters(5, 5, 2, 1)), should.Equal, 15)
	should.So(t, Score(finishGutters(MaxPins, 3, 2, 1)), should.Equal, 21)
	should.So(t, Score(RepeatN(12, MaxFrames)), should.Equal, 300)
}

func finishGutters(rolls ...int) iter.Seq[int] {
	return Concat(Iterator(rolls), RepeatN(20, 0))
}
func Score(rolls iter.Seq[int]) int {
	return Sum(Flatten(AllFrames(rolls)))
}
func AllFrames(rolls iter.Seq[int]) iter.Seq[iter.Seq[int]] {
	return func(yield func(iter.Seq[int]) bool) {
		for x := 0; x < MaxFrames; x++ {
			frame, throws := SingleFrame(rolls)
			if !yield(frame) {
				return
			}
			rolls = Drop(throws, rolls)
		}
	}
}
func SingleFrame(rolls iter.Seq[int]) (frame iter.Seq[int], rollsInFrame int) {
	switch {
	case isStrike(rolls):
		return Take(3, rolls), 1
	case isSpare(rolls):
		return Take(3, rolls), 2
	default:
		return Take(2, rolls), 2
	}
}

func isSpare(rolls iter.Seq[int]) bool  { return Sum(Take(2, rolls)) == MaxPins }
func isStrike(rolls iter.Seq[int]) bool { return First(rolls) == MaxPins }

const (
	MaxPins   = 10
	MaxFrames = 10
)
