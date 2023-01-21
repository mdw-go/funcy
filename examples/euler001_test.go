package examples

import (
	"testing"

	. "github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/funcy/internal/should"
)

// https://projecteuler.net/problem=1
func TestProjectEuler001(t *testing.T) {
	should.So(t, Calculate001(10), should.Equal, 23)
	should.So(t, Calculate001(1000), should.Equal, 233_168)
}
func Calculate001(n int) int { return Sum(Filter(Divisible, Range(0, n))) }
func Divisible(n int) bool   { return n%3 == 0 || n%5 == 0 }
