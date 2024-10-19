package examples

import (
	"testing"

	. "github.com/mdw-go/funcy/ranger"
	"github.com/mdw-go/funcy/ranger/is"
	"github.com/mdw-go/funcy/ranger/op"
)

// A somewhat interesting example, based on this Clojure threading macro example:
// https://clojuredocs.org/clojure.core/-%3E%3E#example-542692c8c026201cdc326a52
// (->> (range) (map #(* % %)) (filter even?) (take 10) (reduce +))  ; output: 1140
// The Clojure solution is 64 bytes of code, but only 77 in Go!
func TestThreadingMacroExample(t *testing.T) {
	result := Reduce(op.Add, 0, Take(10, Filter(is.Even, Map(op.Square, RangeOpen(0, 1)))))
	if result != 1140 {
		t.Fatalf("got %d, want %d", result, 1140)
	}
}
