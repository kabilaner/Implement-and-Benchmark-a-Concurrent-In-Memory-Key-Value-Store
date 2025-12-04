# Concurrent In-Memory Key-Value Store (Go)

This project implements:
- A concurrent key-value store using lock striping
- A baseline version using a single global mutex
- A multithreaded benchmark suite comparing both implementations

## Run benchmark

cd cmd/benchmain
go run .

markdown
Copy code

Results will print throughput (ops/sec) for:
- 1 thread
- 2 threads
- 4 threads
- 8 threads
- 16 threads
- 32 threads

## Folder Structure

concurrent_kv_store/
cmd/
benchmain/
main.go
pkg/
kv/
concurrent.go
baseline.go
bench/
bench.go
go.mod

Copy code
7. RESULTS.txt
(leave empty; your benchmark will fill it manually)

 NEXT STEPS FOR YOU
1. Create the folders
cmd/benchmain, pkg/kv, pkg/bench

2. Create the files
Paste the code exactly into each file.

3. Run:
bash
Copy code
cd cmd/benchmain
go run .
 When you run it, you should see output like:
yaml
Copy code
=== Benchmark: Concurrent vs Baseline Map ===

Threads: 1
ConcurrentMap: 500000 ops/sec (10ms)
BaselineMap:   200000 ops/sec (25ms)

Threads: 2
ConcurrentMap: 900000 ops/sec
BaselineMap:   210000 ops/sec
...
