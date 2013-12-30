
Cirru Grammar
------

[Cirru Parser][parser] implemented in Go.
[parser]: https://github.com/Cirru/cirru-parser

Visit http://cirru.org for more.

### Usage

[![GoDoc](https://godoc.org/github.com/Cirru/cirru-grammar?status.png)](https://godoc.org/github.com/Cirru/cirru-grammar)

Package could be run like this when a file named `demo.cr` is given:

```go
package main

import (
  "github.com/Cirru/cirru-grammar"
  "io/ioutil"
)

func main() {
  filename := "demo.cr"
  codeByte, _ := ioutil.ReadFile(filename)
  code := string(codeByte)
  ast := cirru.Parse(code, filename)
  navigate(ast)
}

func navigate(tree cirru.List) {
  for _, item := range tree {
    if token, ok := item.(cirru.Token); ok {
      println(token.Text)
    }
    if list, ok := item.(cirru.List); ok {
      navigate(list)
    }
  }
}
```

**This package is still in developing, read code before using**

### License

MIT