package op

import (
	"testing"

	"github.com/mdwhatcott/funcy/ranger/internal/should"
)

func Test(t *testing.T) {
	should.So(t, Square(4), should.Equal, 16)
	should.So(t, Add(4, 3), should.Equal, 7)
	should.So(t, Sub(4, 3), should.Equal, 1)
	should.So(t, Mul(4, 3), should.Equal, 12)
	should.So(t, Div(8, 4), should.Equal, 2)
	should.So(t, Pow(2, 4), should.Equal, 16)
	should.So(t, Mod(7, 3), should.Equal, 1)
	should.So(t, Abs(2), should.Equal, 2)
	should.So(t, Abs(-2), should.Equal, 2)
}
