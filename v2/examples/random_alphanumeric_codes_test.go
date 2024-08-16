package examples

import (
	"testing"

	. "github.com/mdwhatcott/funcy/v2"
)

// https://michaelwhatcott.com/generating-random-alphanumeric-codes-in-clojure/

func TestGenerate(t *testing.T) {
	DoAll(
		func(s string) { t.Log(s) },
		Take(5, Repeatedly(Generate)),
	)
}
func Generate() string { return string(Slice(Take(8, Repeatedly(randRune)))) }
func randRune() rune   { return RandNth(alphaNumeric) }

var alphaNumeric = Concat(
	Range('a', 'z'),
	Range('A', 'Z'),
	Range('0', '9'),
)
