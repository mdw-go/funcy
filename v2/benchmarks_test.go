package funcy

import (
	"testing"

	"github.com/mdwhatcott/funcy/v2/is"
)

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	sum := 0
	for i := 0; i < b.N; i++ {
		for x := range Filter(is.Even[int], Range(0, 100_000)) {
			sum += x
		}
	}
	b.StopTimer()
	b.Log(sum)
	// BenchmarkBowlingv2-12    	    2684	    421081 ns/op	       0 B/op	       0 allocs/op
}
