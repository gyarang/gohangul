# gohangul
gohangul 은 한글로 이루어진 문장의 자음과 모음을 분리하는 Go 라이브러리입니다. 자모 분리 또는 초성 검색에 사용할 수 있습니다. Hangul.js 의 Go 구현을 목표로 합니다.

[![Go Reference](https://pkg.go.dev/badge/github.com/gyarang/gohangul.svg)](https://pkg.go.dev/github.com/gyarang/gohangul)

## Installation
```bash
go get github.com/gyarang/gohangul
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/gyarang/gohangul"
)

func main() {
	dism := gohangul.Disassemble("안녕") // []rune{'ㅇ', 'ㅏ', 'ㄴ', 'ㄴ', 'ㅕ', 'ㅇ'}
	fmt.Println(gohangul.Assemble(dism)) // 안녕
}
```

## Documentation
[GoDoc](https://pkg.go.dev/github.com/gyarang/gohangul)

## 참조
[Hangul.js](https://github.com/e-/Hangul.js)
