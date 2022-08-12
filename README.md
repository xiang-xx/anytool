# anytool
A tool to get data from any type.

## install
```
go get github.com/xiang-xx/anytool
```

## use
```go
a := map[string]interface{}{
        "users": []map[string]interface{}{
            {
                "id": "12",
            },
        },
    }
path := "users/0/id"
id, err := anytool.Get(a, path)

id, err := anytool.GetString(a, path)
```

Use "/" split data path, so **do not support "/" in map key**.

**Struct data is not supported**.

## Benchmark

The larger the input data, the slower the code.

```
goos: linux
goarch: amd64
pkg: github.com/xiang-xx/anytool
cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
BenchmarkGet-8                   1000000               530.5 ns/op
BenchmarkGetSlow-8               1000000               631.3 ns/op
BenchmarkGetTwo-8                1000000               667.4 ns/op
BenchmarkGetBig-8                1000000              8125 ns/op
BenchmarkGetBigSlow-8            1000000             14168 ns/op
```
