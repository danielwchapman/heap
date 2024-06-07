# heap

A generic heap implementation. It's advanage over Go's container/heap package is less code is required be written to use it. While this code is not overly complex, it is not out-of-the-box and certainly more error prone. See [IntHeap and Priority Queue examples from official documentation](https://pkg.go.dev/container/heap#pkg-overview).

It rivals the performance of `container/heap`. After several trivial benchmark runs on an M1 MacBook Air, it constantly runs slightly slower for `Push` operations and slightly faster for `Pop` operatations. If performance is a major requirement, it's worth building out the benchmark tests more.

Sample Benchmark Output:
```
BenchmarkIntPush-8                      42626164                27.47 ns/op
BenchmarkIntPop-8                        8839105               202.1 ns/op
BenchmarkContainerIntHeapPush-8         100000000               20.10 ns/op
BenchmarkContainerIntHeapPop-8           6532185               248.3 ns/op
BenchmarkStructPush-8                   12415504                90.49 ns/op
BenchmarkStructPop-8                     2975512               535.5 ns/op
BenchmarkContainerStructHeapPush-8      11499754               106.3 ns/op
BenchmarkContainerStructHeapPop-8        2774511               602.6 ns/op
```

### Testing

```
go test ./...
```

```
go test -bench .
```
