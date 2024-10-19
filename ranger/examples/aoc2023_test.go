package examples

import (
	"iter"
	"strconv"
	"testing"
	"unicode"

	. "github.com/mdw-go/funcy/ranger"
	"github.com/mdw-go/funcy/ranger/internal/should"
)

func TestAdventOfCode2023Day1Part1(t *testing.T) {
	// https://adventofcode.com/2023/day/1
	should.So(t, Calibrate("1abc2"), should.Equal, 12)
	should.So(t, Calibrate("pqr3stu8vwx"), should.Equal, 38)
	should.So(t, Calibrate("a1b2c3d4e5f"), should.Equal, 15)
	should.So(t, Calibrate("treb7uchet"), should.Equal, 77)
	should.So(t, CalibrateAll(Variadic("1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet")),
		should.Equal, 142)
}
func CalibrateAll(lines iter.Seq[string]) int {
	return Sum(Map(Calibrate, lines))
}
func Calibrate(s string) int {
	return ParseInt(string(Bookends(Filter(unicode.IsDigit, Iterator([]rune(s))))))
}
func Bookends(i iter.Seq[rune]) (result []rune) {
	return append(result, First(i), Last(i))
}
func ParseInt(s string) int {
	parsed, _ := strconv.Atoi(s)
	return parsed
}
