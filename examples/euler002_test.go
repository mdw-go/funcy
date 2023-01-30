package examples

import (
	"testing"

	. "github.com/mdwhatcott/funcy"
	"github.com/mdwhatcott/funcy/internal/should"
)

// https://projecteuler.net/problem=2
func TestProjectEuler002(t *testing.T) {
	should.So(t, Calculate002(90), should.Equal, 44)
	should.So(t, Calculate002(4_000_000), should.Equal, 4_613_732)
}
func Calculate002(ceiling int) int {
	return Sum(Filter(IsEven[int], Fib(ceiling, []int{0, 1})))
}
func Fib(ceiling int, results []int) []int {
	next := Sum(DropAllBut(2, results))
	if next >= ceiling {
		return results
	}
	return Fib(ceiling, append(results, next))
}
