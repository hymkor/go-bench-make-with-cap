Report for benchmark between make with capacity vs append n-times
=================================================================

1. `[]string{}` and `append` N-times
2. `make([]string,0,N)` and `append` N-times
3. `make([]string,N)` and no `append` (access with index)

This repository also serves as your own template for how to take benchmarks

Source
------

```go
package main

import (
    "fmt"
    "testing"
)

func BenchmarkAppendN(b *testing.B) {
    buffer := []string{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        buffer = append(buffer, fmt.Sprintf("%d", i))
    }
}

func BenchmarkMakeWithCapAndAppend(b *testing.B) {
    buffer := make([]string, 0, b.N)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        buffer = append(buffer, fmt.Sprintf("%d", i))
    }
}

func BenchmarkNoAppend(b *testing.B) {
    buffer := make([]string, b.N)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        buffer[i] = fmt.Sprintf("%d", i)
    }
}
```

Result
------

```
go fmt
go test -bench . -benchmem
goos: windows
goarch: amd64
pkg: github.com/hymkor/go-bench-make-with-cap
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
BenchmarkAppendN-4                   5515822           237.3 ns/op       111 B/op          1 allocs/op
BenchmarkMakeWithCapAndAppend-4      9754629           125.8 ns/op        15 B/op          1 allocs/op
BenchmarkNoAppend-4                  9762398           125.3 ns/op        15 B/op          1 allocs/op
PASS
ok      github.com/hymkor/go-bench-make-with-cap    5.728s
```
