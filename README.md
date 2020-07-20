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
