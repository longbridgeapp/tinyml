# TinyML

LongBridge Plain Text format.

## Usage

```go
import (
  "github.com/long-bridge/tinyml"
)

func main() {
  fmt.Println(tinyml.ToText("消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。"))
  // 消息称 阿里巴巴.US 将于 5 月，在港股上市。
}
```

## Benchmark

Run `go test -bench=.` to benchmark.

```bash
BenchmarkToText-12    	   41008	     28788 ns/op	    3626 B/op	      74 allocs/op
PASS
ok  	github.com/long-bridge/tinyml	2.037s
```
