package examples

import (
	"testing"

	"github.com/mdwhatcott/funcy/ranger"
	"github.com/mdwhatcott/funcy/ranger/is"
)

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	sum := 0
	for i := 0; i < b.N; i++ {
		for x := range ranger.Filter(is.Even[int], ranger.Range(0, 100_000)) {
			sum += x
		}
	}
	b.StopTimer()
	b.Log(sum)
	// Benchmark    	    2684	    421081 ns/op	       0 B/op	       0 allocs/op
}
