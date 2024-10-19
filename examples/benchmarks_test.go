package examples

import (
	"testing"

	"github.com/mdw-go/funcy"
)

func Benchmark(b *testing.B) {
	b.ReportAllocs()
	sum := 0
	for i := 0; i < b.N; i++ {
		for _, x := range funcy.Filter(funcy.IsEven[int], funcy.Range(0, 100_000)) {
			sum += x
		}
	}
	b.StopTimer()
	b.Log(sum)
	// BenchmarkBowlingv1-12    	     687	   1459875 ns/op	 6056488 B/op	      53 allocs/op
}
