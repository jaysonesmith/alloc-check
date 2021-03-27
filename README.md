# alloc-check

Just some examples of different approaches to things using different packages.

`go test -benchmem -run=Bench -bench=.`

```
goos: darwin
goarch: amd64
pkg: github.com/jaysonesmith/alloc-check
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkViperDoubleDip-12                    520132       2038 ns/op        448 B/op       30 allocs/op
BenchmarkOSDoubleDip-12                      2797027        433.3 ns/op      144 B/op        7 allocs/op
BenchmarkViperSingleDip-12                    885140       1391 ns/op        368 B/op       22 allocs/op
BenchmarkOSSingleDip-12                      2783553        429.6 ns/op      144 B/op        7 allocs/op
BenchmarkViperGetBool-12                     4431043        270.7 ns/op       80 B/op        5 allocs/op
BenchmarkOSGetBool-12                       36298964         33.62 ns/op       0 B/op        0 allocs/op
BenchmarkViperGetBoolNotSet-12               4159965        271.5 ns/op       64 B/op        4 allocs/op
BenchmarkOSGetBoolNotSet-12                 40689571         29.62 ns/op       0 B/op        0 allocs/op
BenchmarkViperGetBoolWithError-12            2272287        528.1 ns/op      160 B/op       10 allocs/op
BenchmarkOSGetBoolWithError-12              36374661         33.11 ns/op       0 B/op        0 allocs/op
BenchmarkViperGetBoolWithErrorNotSet-12      4133439        287.8 ns/op       80 B/op        5 allocs/op
BenchmarkOSGetBoolWithErrorNotSet-12        22429851         53.40 ns/op      16 B/op        1 allocs/op
BenchmarkOSGetBoolWithErrorNotBool-12       21021061         56.84 ns/op      16 B/op        1 allocs/op
PASS
coverage: 75.3% of statements
ok    github.com/jaysonesmith/alloc-check 20.177s
```
