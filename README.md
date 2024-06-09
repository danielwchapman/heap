# heap

A generic heap implementation for Go.

It's advanage over Go's container/heap package is less code is required be written to use it. While this code is not overly complex, it is not out-of-the-box and certainly more error prone. See the IntHeap and Priority Queue samples from [official documentation](https://pkg.go.dev/container/heap#pkg-overview) to get a feel for overhead code required by the standard library implementation.

It rivals the performance of `container/heap`. After several trivial benchmarks, it constantly runs slower for `Push` operations and consistently faster for `Pop` operatations. If performance is an important requirement, it's worth building out the benchmark tests more.

Sample Benchmark Output from an M1 MacBook Air:
```
go test -bench .
goos: darwin
goarch: arm64
pkg: github.com/danielwchapman/heap
BenchmarkIntPush-8                      42510336                27.25 ns/op
BenchmarkContainerIntPush-8             100000000               18.59 ns/op
BenchmarkIntPop-8                        8255620               198.7 ns/op
BenchmarkContainerIntPop-8               6659751               233.2 ns/op
BenchmarkStructPush-8                   14317743                87.51 ns/op
BenchmarkContainerStructPush-8          12182853               106.0 ns/op
BenchmarkStructPop-8                     3216393               508.5 ns/op
BenchmarkContainerStructPop-8            2928676               587.8 ns/op
PASS
ok      github.com/danielwchapman/heap  15.639s
```
Benchmark tests with `container` in the name refer to benchmarking the `container/heap` package for comparision.

### Installation
```
go get github.com/danielwchapman/heap           
```

### Testing
```
go test ./...
```

```
go test -bench .
```
