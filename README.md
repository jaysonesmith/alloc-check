# alloc-check

Just some examples of different approaches to things using different packages.

`go test -benchmem -run=Bench -bench=.`

```
goos: darwin
goarch: amd64
pkg: github.com/jaysonesmith/alloc-check
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkViperDoubleDip-12          457293        2500 ns/op     576 B/op      30 allocs/op
BenchmarkViperSingleDip-12          747566        1669 ns/op     432 B/op      22 allocs/op
BenchmarkOSDoubleDip-12            2596216       474.9 ns/op     144 B/op       7 allocs/op
BenchmarkOSSingleDip-12            2296303       471.5 ns/op     144 B/op       7 allocs/op
BenchmarkViperGetBool-12           2244579       644.4 ns/op     160 B/op      10 allocs/op
BenchmarkViperGetBoolNotSet-12     4427404       264.4 ns/op      64 B/op       4 allocs/op
BenchmarkOSGetBool-12             37025508        32.44 ns/op      0 B/op       0 allocs/op
BenchmarkOSGetBoolNotSet-12       40941196        29.01 ns/op      0 B/op       0 allocs/op
PASS
ok  	github.com/jaysonesmith/alloc-check	12.708s
```
