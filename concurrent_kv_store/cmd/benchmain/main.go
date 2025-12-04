package main

import (
	"fmt"
	"concurrent_kv_store/pkg/bench"
)

func main() {
	threadCounts := []int{1, 2, 4, 8, 16, 32}
	operations := 50000

	fmt.Println("=== Benchmark: Concurrent vs Baseline Map ===")

	for _, t := range threadCounts {
		fmt.Printf("\nThreads: %d\n", t)

		conRes := bench.RunBenchmark("concurrent", t, operations)
		fmt.Printf("ConcurrentMap: %.2f ops/sec  (%v)\n", conRes.Throughput, conRes.Duration)

		baseRes := bench.RunBenchmark("baseline", t, operations)
		fmt.Printf("BaselineMap:   %.2f ops/sec  (%v)\n", baseRes.Throughput, baseRes.Duration)
	}
}
