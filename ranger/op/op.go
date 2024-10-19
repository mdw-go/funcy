package op

import (
	"math"

	"github.com/mdw-go/funcy/ranger/is"
)

func Square[N is.Integer](n N) N { return Mul(n, n) }

func Add[N is.Number](a, b N) N  { return a + b }
func Sub[N is.Number](a, b N) N  { return a - b }
func Div[N is.Number](a, b N) N  { return a / b }
func Mul[N is.Number](a, b N) N  { return a * b }
func Pow[N is.Number](a, b N) N  { return N(math.Pow(float64(a), float64(b))) }
func Mod[N is.Integer](a, b N) N { return a % b }
func Abs[N is.Number](n N) N {
	if n < 0 {
		return -n
	}
	return n
}
