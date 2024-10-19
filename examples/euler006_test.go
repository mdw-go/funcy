package examples

import (
	"testing"

	. "github.com/mdw-go/funcy"
	"github.com/mdw-go/funcy/internal/should"
)

// https://projecteuler.net/problem=6
func TestProjectEuler006(t *testing.T) {
	should.So(t, Calculate006(10), should.Equal, 2640)
	should.So(t, Calculate006(100), should.Equal, 25164150)
}
func Calculate006(n int) int { return SquareOfSums(n) - SumOfSquares(n) }
func SquareOfSums(n int) int { return Square(Sum(Range(1, n+1))) }
func SumOfSquares(n int) int { return Sum(Map(Square[int], Range(1, n+1))) }
