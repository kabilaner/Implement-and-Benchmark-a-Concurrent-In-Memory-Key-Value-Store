package bench

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"concurrent_kv_store/pkg/kv"
)

type BenchResult struct {
	Ops        int
	Duration   time.Duration
	Throughput float64
}

// RunBenchmark runs a mixed read/write/delete workload.
func RunBenchmark(mapType string, threads int, operations int) BenchResult {
	// store must implement Get, Put, Delete
	var store interface {
		Put(string, string)
		Get(string) (string, bool)
		Delete(string)
	}

	if mapType == "concurrent" {
		store = kv.NewConcurrentMap()
	} else {
		store = kv.NewBaselineMap()
	}

	rand.Seed(time.Now().UnixNano())

	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(threads)

	for i := 0; i < threads; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < operations; j++ {
				op := rand.Intn(100)
				key := fmt.Sprintf("key_%d", rand.Intn(50000))

				if op < 70 {
					// 70% GET → ignore returned value
					_, _ = store.Get(key)

				} else if op < 90 {
					// 20% PUT
					store.Put(key, "value")

				} else {
					// 10% DELETE
					store.Delete(key)
				}
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	thr := float64(threads*operations) / duration.Seconds()

	return BenchResult{
		Ops:        threads * operations,
		Duration:   duration,
		Throughput: thr,
	}
}
