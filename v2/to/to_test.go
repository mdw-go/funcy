package to_test

import (
	"testing"
	"time"

	"github.com/mdwhatcott/funcy/v2/internal/should"
	"github.com/mdwhatcott/funcy/v2/to"
)

func TestString(t *testing.T) {
	should.So(t, to.String(time.Second), should.Equal, "1s")
	should.So(t, to.String("1"), should.Equal, "1")
	should.So(t, to.String(1), should.Equal, "1")
	should.So(t, to.String(uint8(1)), should.Equal, "1")
	should.So(t, to.String(uint16(1)), should.Equal, "1")
	should.So(t, to.String(uint32(1)), should.Equal, "1")
	should.So(t, to.String(uint64(1)), should.Equal, "1")
	should.So(t, to.String(int8(1)), should.Equal, "1")
	should.So(t, to.String(int16(1)), should.Equal, "1")
	should.So(t, to.String(int32(1)), should.Equal, "1")
	should.So(t, to.String(int64(1)), should.Equal, "1")
	should.So(t, to.String(uintptr(1)), should.Equal, "1")
	should.So(t, to.String(float32(1)), should.Equal, "1")
	should.So(t, to.String(float64(1)), should.Equal, "1")
	should.So(t, to.String(uint(1)), should.Equal, "1")
	should.So(t, to.String(true), should.Equal, "true")
	should.So(t, to.String([]int(nil)), should.Equal, "[]")
}
