package is_test

import (
	"iter"
	"strings"
	"testing"

	"github.com/mdwhatcott/funcy/ranger/internal/should"
	"github.com/mdwhatcott/funcy/ranger/is"
)

func seq[V any](s ...V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}
func TestEmpty(t *testing.T) {
	should.So(t, is.Empty(seq[int]()), should.BeTrue)
	should.So(t, is.Empty(seq[int](1)), should.BeFalse)
	should.So(t, is.Empty(seq[int](1, 2, 3)), should.BeFalse)
}
func TestZero(t *testing.T) {
	should.So(t, is.Zero(0), should.BeTrue)
	should.So(t, is.Zero(1), should.BeFalse)
	should.So(t, is.Zero(0.0), should.BeTrue)
	should.So(t, is.Zero(1.0), should.BeFalse)
}
func TestPositive(t *testing.T) {
	should.So(t, is.Positive(-1.0), should.BeFalse)
	should.So(t, is.Positive(-1), should.BeFalse)
	should.So(t, is.Positive(0), should.BeFalse)
	should.So(t, is.Positive(1), should.BeTrue)
	should.So(t, is.Positive(1.0), should.BeTrue)
}
func TestNegative(t *testing.T) {
	should.So(t, is.Negative(-1.0), should.BeTrue)
	should.So(t, is.Negative(-1), should.BeTrue)
	should.So(t, is.Negative(0), should.BeFalse)
	should.So(t, is.Negative(1), should.BeFalse)
	should.So(t, is.Negative(1.0), should.BeFalse)
}
func TestEven(t *testing.T) {
	should.So(t, is.Even(-5), should.BeFalse)
	should.So(t, is.Even(-4), should.BeTrue)
	should.So(t, is.Even(-3), should.BeFalse)
	should.So(t, is.Even(-2), should.BeTrue)
	should.So(t, is.Even(-1), should.BeFalse)
	should.So(t, is.Even(0), should.BeTrue)
	should.So(t, is.Even(1), should.BeFalse)
	should.So(t, is.Even(2), should.BeTrue)
	should.So(t, is.Even(3), should.BeFalse)
	should.So(t, is.Even(4), should.BeTrue)
	should.So(t, is.Even(5), should.BeFalse)
}
func TestOdd(t *testing.T) {
	should.So(t, is.Odd(-5), should.BeTrue)
	should.So(t, is.Odd(-4), should.BeFalse)
	should.So(t, is.Odd(-3), should.BeTrue)
	should.So(t, is.Odd(-2), should.BeFalse)
	should.So(t, is.Odd(-1), should.BeTrue)
	should.So(t, is.Odd(0), should.BeFalse)
	should.So(t, is.Odd(1), should.BeTrue)
	should.So(t, is.Odd(2), should.BeFalse)
	should.So(t, is.Odd(3), should.BeTrue)
	should.So(t, is.Odd(4), should.BeFalse)
	should.So(t, is.Odd(5), should.BeTrue)
}
func TestNil(t *testing.T) {
	var a any
	should.So(t, is.Nil(a), should.BeTrue)
	should.So(t, is.Nil([]int(nil)), should.BeTrue)
	var b *strings.Builder
	should.So(t, is.Nil(b), should.BeTrue)
	should.So(t, is.Nil(4), should.BeFalse)
	should.So(t, is.Nil(strings.NewReader("HI")), should.BeFalse)
}
func TestGreaterThan(t *testing.T) {
	should.So(t, is.GreaterThan(1.0, 2.0), should.BeFalse)
	should.So(t, is.GreaterThan(1.0, 1.0), should.BeFalse)
	should.So(t, is.GreaterThan(2.0, 1.0), should.BeTrue)
}
func TestLessThan(t *testing.T) {
	should.So(t, is.LessThan(1.0, 2.0), should.BeTrue)
	should.So(t, is.LessThan(1.0, 1.0), should.BeFalse)
	should.So(t, is.LessThan(2.0, 1.0), should.BeFalse)
}
func TestEqualTo(t *testing.T) {
	should.So(t, is.EqualTo(1.0, 2.0), should.BeFalse)
	should.So(t, is.EqualTo(1.0, 1.0), should.BeTrue)
	should.So(t, is.EqualTo(2.0, 1.0), should.BeFalse)
}
