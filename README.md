# TinyML

LongBridge Plain Text format.

[![Build Status](https://travis-ci.org/long-bridge/tinyml.svg?branch=master)](https://travis-ci.org/long-bridge/tinyml)

## Usage

```go
package main

import (
	"fmt"

	"github.com/long-bridge/tinyml"
)

func main() {
        out, _ := tinyml.ToText("消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。")
	fmt.Println(out)
	// 消息称 阿里巴巴.US 将于 5 月，在港股上市。

	out, _ = tinyml.ToHTML("消息称[st]ST/US/BABA#阿里巴巴.US[/st]将于5月，在港股上市。\n\n几家领头羊都处于第一步或者第二步。")
	fmt.Println(out)
	// <p>消息称 <span class="security-tag" value="ST/US/BABA#阿里巴巴.US" data-id="ST/US/BABA">阿里巴巴.US</span> 将于 5 月，在港股上市。</p><p>几家领头羊都处于第一步或者第二步。</p>
}
```

Example:

https://play.golang.org/p/vNGesNpEQzx

## Benchmark

Run `go test -bench=.` to benchmark.

```bash
BenchmarkToHTML-12    	   19084	     62621 ns/op
BenchmarkToText-12    	   41409	     28675 ns/op
```
